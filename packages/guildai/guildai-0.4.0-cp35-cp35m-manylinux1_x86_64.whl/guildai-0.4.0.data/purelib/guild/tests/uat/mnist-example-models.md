# MNIST example models

In this test we'll list the available models for the MNIST example. As
example directory contains a file named `guild.yml` ("guildfile")
Guild includes models from that directory, as well as globally
installed models.

We can reference a guildfile in one of two ways:

- Change the current directory to the guildfile directory
- Using the `-C` option of the `guild` command to reference the
  guildfile drectory

Here are the models associated with the MNIST example (using the `-C`
option):

    >>> run("guild -C examples/mnist2 models", ignore="FutureWarning")
    ./examples/mnist2/mnist-expert  MNIST model from TensorFlow expert tutorial
    ./examples/mnist2/mnist-intro   MNIST model from TensorFlow intro tutorial
    hello/...
    keras.mnist/...
    <exit 0>

Here's the same list after we've changed to that directory:

    >>> cd("examples/mnist2")
    >>> run("guild models", ignore="FutureWarning")
    ./mnist-expert  MNIST model from TensorFlow expert tutorial
    ./mnist-intro   MNIST model from TensorFlow intro tutorial
    hello/...
    keras.mnist/...
    <exit 0>

We can limit the results to model defined in the current directory by
specifying it as a filter:

    >>> run("guild models .", ignore="FutureWarning")
    ./mnist-expert       MNIST model from TensorFlow expert tutorial
    ./mnist-intro        MNIST model from TensorFlow intro tutorial
    <exit 0>
