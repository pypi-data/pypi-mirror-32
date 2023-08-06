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

from guild import cli
from guild import cmd_impl_support
from guild import model
from guild import util

def main(args):
    cmd_impl_support.init_model_path()
    gf_dirs, filters = cmd_impl_support.guildfile_dirs(args.filters)
    formatted = [_format_model(m) for m in iter_models(gf_dirs)]
    filtered = [m for m in formatted if _filter_model(m, filters)]
    cli.table(
        sorted(filtered, key=lambda m: m["fullname"]),
        cols=["fullname", "description"],
        detail=(["source", "operations", "details"] if args.verbose else [])
    )

def iter_models(gf_dirs):
    for m in model.iter_models():
        if not gf_dirs:
            yield m
        if any((m.modeldef.guildfile.dir == dir for dir in gf_dirs)):
            yield m

def _format_model(model):
    modeldef = model.modeldef
    description, details = util.split_description(modeldef.description)
    return {
        "fullname": model.fullname,
        "name": modeldef.name,
        "source": modeldef.guildfile.src,
        "description": util.strip_trailing_period(description),
        "details": details,
        "operations": ", ".join([op.name for op in modeldef.operations])
    }

def _filter_model(model, filters):
    filter_vals = [
        model["fullname"],
        model["description"],
    ]
    return util.match_filters(filters, filter_vals)
