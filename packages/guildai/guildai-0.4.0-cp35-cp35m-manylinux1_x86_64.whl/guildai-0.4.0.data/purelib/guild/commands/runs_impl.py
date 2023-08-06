# Copyright 2017-2018 TensorHub, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from __future__ import absolute_import
from __future__ import division

import logging
import os
import re
import shutil
import signal

import six
import yaml

import guild.opref
import guild.remote
import guild.run

from guild import cli
from guild import click_util
from guild import cmd_impl_support
from guild import config
from guild import remote_run_support
from guild import util
from guild import var

log = logging.getLogger("guild")

RUN_DETAIL = [
    "id",
    "operation",
    "status",
    "started",
    "stopped",
    "run_dir",
    "command",
    "exit_status",
    "pid",
]

ALL_RUNS_ARG = [":"]
LATEST_RUN_ARG = ["0"]

CORE_RUN_ATTRS = [
    "cmd",
    "deps",
    "env",
    "exit_status",
    "exit_status.remote",
    "flags",
    "label",
    "opref",
    "started",
    "stopped",
]

RUNS_PER_GROUP = 20

def runs_for_args(args, ctx=None, force_deleted=False):
    filtered = filtered_runs(args, force_deleted)
    return select_runs(filtered, args.runs, ctx)

def filtered_runs(args, force_deleted=False):
    return var.runs(
        _runs_root_for_args(args, force_deleted),
        sort=["-started"],
        filter=_runs_filter(args),
        run_init=init_opref_attr)

def _runs_root_for_args(args, force_deleted):
    archive = getattr(args, "archive", None)
    if archive:
        if not os.path.exists(archive):
            cli.error("archive directory '%s' does not exist" % archive)
        return archive
    deleted = force_deleted or getattr(args, "deleted", False)
    return var.runs_dir(deleted=deleted)

def _runs_filter(args):
    filters = []
    _apply_status_filter(args, filters)
    _apply_ops_filter(args, filters)
    _apply_labels_filter(args, filters)
    return var.run_filter("all", filters)

def _apply_status_filter(args, filters):
    status_filters = [
        var.run_filter("attr", "status", status)
        for status in ["running", "completed", "error", "terminated"]
        if getattr(args, status, False)
    ]
    if status_filters:
        filters.append(var.run_filter("any", status_filters))

def _apply_ops_filter(args, filters):
    if args.ops:
        filters.append(_op_run_filter(args.ops))

def _op_run_filter(op_refs):
    def f(run):
        op_desc = format_op_desc(run.opref, nowarn=True)
        return any((ref in op_desc for ref in op_refs))
    return f

def _apply_labels_filter(args, filters):
    if args.labels and args.unlabeled:
        cli.error("--label and --unlabeled cannot both be used")
    if args.labels:
        filters.append(_label_filter(args.labels))
    elif args.unlabeled:
        filters.append(_unlabeled_filter())

def _label_filter(labels):
    def f(run):
        return any((l in run.get("label", "") for l in labels))
    return f

def _unlabeled_filter():
    def f(run):
        return not run.get("label", "").strip()
    return f

def select_runs(runs, select_specs, ctx=None):
    if not select_specs:
        return runs
    selected = []
    for spec in select_specs:
        try:
            slice_start, slice_end = _parse_slice(spec)
        except ValueError:
            selected.append(_find_run_by_id(spec, runs, ctx))
        else:
            if _in_range(slice_start, slice_end, runs):
                selected.extend(runs[slice_start:slice_end])
            else:
                selected.append(_find_run_by_id(spec, runs, ctx))
    return selected

def _parse_slice(spec):
    try:
        index = int(spec)
    except ValueError:
        m = re.match("(\\d+)?:(\\d+)?", spec)
        if m:
            try:
                return (
                    _slice_part(m.group(1)),
                    _slice_part(m.group(2), incr=True)
                )
            except ValueError:
                pass
        raise ValueError(spec)
    else:
        return index, index + 1

def _slice_part(s, incr=False):
    if s is None:
        return None
    elif incr:
        return int(s) + 1
    else:
        return int(s)

def _find_run_by_id(id_part, runs, ctx):
    matches = [run for run in runs if run.id.startswith(id_part)]
    return cmd_impl_support.one_run(matches, id_part, ctx)

def _in_range(slice_start, slice_end, l):
    return (
        (slice_start is None or slice_start >= 0) and
        (slice_end is None or slice_end <= len(l))
    )

def list_runs(args):
    if args.archive and args.deleted:
        cli.error("--archive and --deleted may not both be used")
    runs = filtered_runs(args)
    formatted = []
    for i, run in enumerate(runs):
        try:
            formatted_run = format_run(run, i)
        except Exception:
            log.exception("formatting run in %s", run.path)
        else:
            formatted.append(formatted_run)
    formatted = _limit_runs(formatted, args)
    cols = ["index", "operation", "started", "status", "label"]
    detail = RUN_DETAIL if args.verbose else None
    cli.table(formatted, cols=cols, detail=detail)

def _limit_runs(runs, args):
    if args.all:
        return runs
    limited = runs[:(args.more + 1) * RUNS_PER_GROUP]
    if len(limited) < len(runs):
        cli.note(
            "Showing the first %i runs (%i total) - use --all "
            "to show all or -m to show more"
            % (len(limited), len(runs)))
    return limited

def _no_selected_runs_error(help_msg=None):
    help_msg = (
        help_msg or
        "No matching runs\n"
        "Try 'guild runs list' to list available runs."
    )
    cli.out(help_msg, err=True)
    cli.error()

def init_opref_attr(run):
    try:
        opref = guild.opref.OpRef.from_run(run)
    except guild.opref.OpRefError as e:
        log.warning("unable to read opref for run %s: %s", run.id, e)
        return None
    else:
        run.opref = opref
        return run

def format_run(run, index=None):
    return {
        "id": run.id,
        "index": _format_run_index(run, index),
        "short_index": _format_run_index(run),
        "model": run.opref.model_name,
        "op_name": run.opref.op_name,
        "operation": format_op_desc(run.opref),
        "pkg": run.opref.pkg_name,
        "status": run.status,
        "label": run.get("label") or "",
        "pid": run.pid or "",
        "started": util.format_timestamp(run.get("started")),
        "stopped": util.format_timestamp(run.get("stopped")),
        "run_dir": run.path,
        "command": _format_command(run.get("cmd", "")),
        "exit_status": _exit_status(run)
    }

def _format_run_index(run, index=None):
    if index is not None:
        return "[%i:%s]" % (index, run.short_id)
    else:
        return "[%s]" % run.short_id

def format_op_desc(opref, nowarn=False):
    if opref.pkg_type == "guildfile":
        return _format_guildfile_op(opref)
    elif opref.pkg_type == "package":
        return _format_package_op(opref)
    else:
        if not nowarn:
            log.warning(
                "cannot format op desc, unexpected pkg type: %s (%s)",
                opref.pkg_type, opref.pkg_name)
        return "?"

def _format_guildfile_op(opref):
    relpath = os.path.relpath(opref.pkg_name, config.cwd())
    if relpath[0] != '.':
        relpath = os.path.join('.', relpath)
    return "%s/%s:%s" % (relpath, opref.model_name, opref.op_name)

def _format_package_op(opref):
    return "%s/%s:%s" % (opref.pkg_name, opref.model_name, opref.op_name)

def _format_command(cmd):
    if not cmd:
        return ""
    return " ".join([_maybe_quote_arg(arg) for arg in cmd])

def _maybe_quote_arg(arg):
    return '"%s"' % arg if " " in arg else arg

def _exit_status(run):
    return run.get("exit_status.remote", "") or run.get("exit_status", "")

def _format_attr_val(val):
    if isinstance(val, list):
        return _format_attr_list(val)
    elif isinstance(val, dict):
        return _format_attr_dict(val)
    else:
        return str(val)

def _format_attr_list(l):
    return "\n%s" % "\n".join([
        "  %s" % item for item in l
    ])

def _format_attr_dict(d):
    return "\n%s" % "\n".join([
        "  %s: %s" % (key, d[key] or "") for key in sorted(d)
    ])

def _runs_op(args, ctx, force_deleted, preview_msg, confirm_prompt,
             no_runs_help, op_callback, default_runs_arg=None,
             confirm_default=False):
    default_runs_arg = default_runs_arg or ALL_RUNS_ARG
    runs_arg = args.runs or default_runs_arg
    filtered = filtered_runs(args, force_deleted)
    selected = select_runs(filtered, runs_arg, ctx)
    if not selected:
        _no_selected_runs_error(no_runs_help)
    preview = [format_run(run) for run in selected]
    if not args.yes:
        cli.out(preview_msg)
        cols = ["short_index", "operation", "started", "status", "label"]
        cli.table(preview, cols=cols, indent=2)
    if args.yes or cli.confirm(confirm_prompt, confirm_default):
        op_callback(selected)

def delete_runs(args, ctx):
    if args.permanent:
        preview = (
            "WARNING: You are about to permanently delete "
            "the following runs:")
        confirm = "Permanently delete these runs?"
    else:
        preview = "You are about to delete the following runs:"
        confirm = "Delete these runs?"
    no_runs_help = "Nothing to delete."
    def delete(selected):
        running = [run for run in selected if run.status == "running"]
        if running and not args.yes:
            cli.out(
                "WARNING: one or more runs are still running "
                "and will be stopped before being deleted.")
            if not cli.confirm("Really delete these runs?"):
                return
        for run in running:
            _stop_run(run, no_wait=True)
        var.delete_runs(selected, args.permanent)
        if args.permanent:
            cli.out("Permanently deleted %i run(s)" % len(selected))
        else:
            cli.out("Deleted %i run(s)" % len(selected))
    _runs_op(
        args, ctx, False, preview, confirm, no_runs_help,
        delete, confirm_default=not args.permanent)

def purge_runs(args, ctx):
    preview = (
        "WARNING: You are about to permanently delete "
        "the following runs:")
    confirm = "Permanently delete these runs?"
    no_runs_help = "Nothing to purge."
    def purge(selected):
        var.purge_runs(selected)
        cli.out("Permanently deleted %i run(s)" % len(selected))
    _runs_op(args, ctx, True, preview, confirm, no_runs_help, purge)

def restore_runs(args, ctx):
    preview = "You are about to restore the following runs:"
    confirm = "Restore these runs?"
    no_runs_help = "Nothing to restore."
    def restore(selected):
        var.restore_runs(selected)
        cli.out("Restored %i run(s)" % len(selected))
    _runs_op(
        args, ctx, True, preview, confirm, no_runs_help,
        restore, confirm_default=True)

def run_info(args, ctx):
    run = one_run(args, ctx)
    if args.page_output:
        _page_run_output(run)
    else:
        _print_run_info(run, args)

def one_run(args, ctx):
    filtered = filtered_runs(args)
    if not filtered:
        cli.out("No matching runs", err=True)
        cli.error()
    runspec = args.run or "0"
    selected = select_runs(filtered, [runspec], ctx)
    return cmd_impl_support.one_run(selected, runspec, ctx)

def _page_run_output(run):
    reader = util.RunOutputReader(run.path)
    lines = []
    try:
        lines = reader.read()
    except IOError as e:
        cli.error("error reading output for run %s: %s" % (run.id, e))
    lines = [_format_output_line(stream, line) for _time, stream, line in lines]
    cli.page("\n".join(lines))

def _format_output_line(stream, line):
    if stream == 1:
        return cli.style(line, fg="red")
    return line

def _print_run_info(run, args):
    formatted = format_run(run)
    out = cli.out
    for name in RUN_DETAIL:
        out("%s: %s" % (name, formatted[name]))
    for name in other_attr_names(run):
        out("%s: %s" % (name, _format_attr(run.get(name))))
    if args.env:
        out("environment:", nl=False)
        out(_format_attr_val(run.get("env", "")))
    if args.flags:
        out("flags:", nl=False)
        out(_format_attr_val(run.get("flags", "")))
    if args.deps:
        out("dependencies:")
        deps = run.get("deps", {})
        for name in sorted(deps):
            out("  %s:" % name)
            for path in deps[name]:
                out("    %s" % path)
    if args.output:
        out("output:")
        for line in _iter_output(run):
            out("  %s" % line, nl=False)
    if args.files or args.all_files:
        out("files:")
        for path in sorted(run.iter_files(args.all_files, args.follow_links)):
            if not args.full_path:
                path = os.path.relpath(path, run.path)
            out("  %s" % path)

def other_attr_names(run):
    return [
        name for name in sorted(run.attr_names())
        if name[0] != "_" and name not in CORE_RUN_ATTRS]

def _format_attr(val):
    if val is None:
        return ""
    elif isinstance(val, (int, float, six.string_types)):
        return str(val)
    else:
        return _format_yaml(val)

def _iter_output(run):
    with open(run.guild_path("output"), "r") as f:
        for line in f:
            yield line

def _format_yaml(val):
    formatted = yaml.dump(val)
    lines = formatted.split("\n")
    padded = ["  " + line for line in lines]
    return "\n" + "\n".join(padded).rstrip()

def label(args, ctx):
    if args.clear:
        _clear_labels(args, ctx)
    elif args.label:
        _set_labels(args, ctx)
    else:
        cli.error("specify a label")

def _clear_labels(args, ctx):
    preview = "You are about to clear the labels for the following runs:"
    confirm = "Continue?"
    no_runs = "No runs to modify."
    def clear_labels(selected):
        for run in selected:
            run.del_attr("label")
        cli.out("Cleared label for %i run(s)" % len(selected))
    _runs_op(
        args, ctx, False, preview, confirm, no_runs,
        clear_labels, LATEST_RUN_ARG, True)

def _set_labels(args, ctx):
    preview = (
        "You are about to label the following runs with '%s':"
        % args.label)
    confirm = "Continue?"
    no_runs = "No runs to modify."
    def set_labels(selected):
        for run in selected:
            run.write_attr("label", args.label)
        cli.out("Labeled %i run(s)" % len(selected))
    _runs_op(
        args, ctx, False, preview, confirm, no_runs,
        set_labels, LATEST_RUN_ARG, True)

def stop_run(args, ctx):
    preview = "WARNING: You are about to stop the following runs:"
    confirm = "Stop these runs?"
    no_runs_help = "Nothing to stop."
    args.running = True
    def stop(selected):
        for run in selected:
            _stop_run(run, args.no_wait)
    _runs_op(args, ctx, False, preview, confirm, no_runs_help, stop)

def _stop_run(run, no_wait):
    remote_lock = remote_run_support.lock_for_run(run)
    if remote_lock:
        _try_stop_remote_run(run, remote_lock, no_wait)
    else:
        _try_stop_local_run(run)

def _try_stop_remote_run(run, remote_lock, no_wait):
    try:
        plugin = guild.plugin.for_name(remote_lock.plugin_name)
    except LookupError:
        log.warning(
            "error syncing run '%s': plugin '%s' not available",
            run.id, remote_lock.plugin_name)
    else:
        cli.out("Stopping %s (remote)" % run.id)
        plugin.stop_run(run, dict(no_wait=no_wait))

def _try_stop_local_run(run):
    pid = run.pid
    if pid and util.pid_exists(pid):
        cli.out("Stopping %s (pid %i)" % (run.id, run.pid))
        os.kill(pid, signal.SIGTERM)

def export(args, ctx):
    if not os.path.isdir(args.location):
        cli.error("directory '{}' does not exist".format(args.location))
    preview = (
        "You are about to %s the following runs to '%s':" %
        (args.move and "move" or "copy", args.location))
    confirm = "Continue?"
    no_runs = "No runs to export."
    def export(selected):
        if args.copy_resources and not args.yes:
            cli.out(
                "WARNING: you have specified --copy-resources, which will "
                "copy resources used by each run!"
                "")
            if not cli.confirm("Really copy resources exported runs?"):
                return
        exported = 0
        for run in selected:
            dest = os.path.join(args.location, run.id)
            if os.path.exists(dest):
                log.warning("%s exists, skipping", dest)
                continue
            if args.move:
                cli.out("Moving {}".format(run.id))
                if args.copy_resources:
                    shutil.copytree(run.path, dest)
                    shutil.rmtree(run.path)
                else:
                    shutil.move(run.path, dest)
            else:
                cli.out("Copying {}".format(run.id))
                symlinks = not args.copy_resources
                shutil.copytree(run.path, dest, symlinks)
            exported += 1
        cli.out("Exported %i run(s)" % exported)
    _runs_op(
        args, ctx, False, preview, confirm, no_runs,
        export, ALL_RUNS_ARG, True)

def import_(args, ctx):
    if not os.path.isdir(args.archive):
        cli.error("directory '{}' does not exist".format(args.archive))
    preview = (
        "You are about to import (%s) the following runs from '%s':" %
        (args.move and "move" or "copy", args.archive))
    confirm = "Continue?"
    no_runs = "No runs to import."
    def export(selected):
        if args.copy_resources and not args.yes:
            cli.out(
                "WARNING: you have specified --copy-resources, which will "
                "copy resources used by each run!"
                "")
            if not cli.confirm("Really copy resources exported runs?"):
                return
        imported = 0
        for run in selected:
            dest = os.path.join(var.runs_dir(), run.id)
            if os.path.exists(dest):
                log.warning("%s exists, skipping", run.id)
                continue
            if args.move:
                cli.out("Moving {}".format(run.id))
                if args.copy_resources:
                    shutil.copytree(run.path, dest)
                    shutil.rmtree(run.path)
                else:
                    shutil.move(run.path, dest)
            else:
                cli.out("Copying {}".format(run.id))
                symlinks = not args.copy_resources
                shutil.copytree(run.path, dest, symlinks)
            imported += 1
        cli.out("Imported %i run(s)" % imported)
    _runs_op(
        args, ctx, False, preview, confirm, no_runs,
        export, ALL_RUNS_ARG, True)

def push(args, ctx):
    remote = _remote_for_name(args.remote)
    preview = (
        "You are about to copy the following runs to '%s':" %
        remote.push_dest())
    confirm = "Continue?"
    no_runs = "No runs to copy."
    def push(runs):
        remote.push(runs, args.verbose)
    _runs_op(
        args, ctx, False, preview, confirm, no_runs,
        push, ALL_RUNS_ARG, True)

def _remote_for_name(name):
    try:
        return guild.remote.for_name(name)
    except guild.remote.NoSuchRemote:
        cli.error(
            "remote '%s' is not defined\n"
            "Remotes are defined in ~/.guild/config.yml. "
            "Try 'guild remotes --help' for more information."
            % name)
    except guild.remote.UnsupportedRemoteType as e:
        cli.error(
            "remote '%s' in ~/.guild/config.yml has unsupported "
            "type: %s" % (name, e.args[0]))
    except guild.remote.MissingRequiredConfig as e:
        cli.error(
            "remote '%s' in ~/.guild/config.yml is missing required "
            "config: %s" % (name, e.args[0]))

def pull(args, ctx):
    if not args.runs and not args.all:
        cli.error(
            "specify one or more runs or use --all\n"
            "Try '%s' for more information."
            % click_util.cmd_help(ctx))
    if args.all and args.runs:
        cli.error(
            "RUN cannot be used with --all\n"
            "Try '%s' for more information."
            % click_util.cmd_help(ctx))
    remote = _remote_for_name(args.remote)
    if args.all:
        run_ids = None
        preview = (
            "You are about to copy all runs from '%s'"
            % remote.pull_src())
    else:
        run_ids = _validate_remote_run_ids(args.runs)
        formatted_run_ids = "\n".join(["  %s" % run_id for run_id in run_ids])
        preview = (
            "You are about to copy the following runs from '%s':\n"
            "%s" % (remote.pull_src(), formatted_run_ids))
    confirm = "Continue?"
    if not args.yes:
        cli.out(preview)
    if args.yes or cli.confirm(confirm, True):
        if args.all:
            remote.pull_all(True)
        else:
            assert run_ids
            remote.pull(run_ids, args.verbose)

def _validate_remote_run_ids(run_ids):
    for run_id in run_ids:
        if len(run_id) != 32:
            cli.error(
                "invalid remote RUN '%s'\n"
                "Remote run IDs must be 32 characters long."
                % run_id)
    return run_ids
