# Train multiple matches

Guild lets the user specify a partial model name for operations, but
if there are multiple matches for the specified term, Guild exits with
an error message.

    >>> run("guild train -y mnist epochs=1", ignore="FutureWarning")
    guild: there are multiple models that match 'mnist'
    Try specifying one of the following:
      mnist/mnist-cnn
      mnist/mnist-samples
      mnist/mnist-softmax
    <exit 1>
