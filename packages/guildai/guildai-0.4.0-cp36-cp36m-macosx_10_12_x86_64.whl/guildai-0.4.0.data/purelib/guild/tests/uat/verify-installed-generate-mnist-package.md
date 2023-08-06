# Verify installed generated `mnist` package

    >>> run("guild packages info mnist")
    name: mnist
    version: ...
    summary: CNN and softmax regression classifiers for MNIST digits
    home-page: https://github.com/guildai/index/tree/master/mnist
    author: Guild AI
    author-email: packages@guild.ai
    license: Apache 2.0
    location: /.../lib/python.../site-packages
    requires: []
    <exit 0>

    >>> run("guild models mnist", ignore="FutureWarning")
    mnist/mnist-cnn      CNN classifier for MNIST
    mnist/mnist-samples  Sample MNIST images
    mnist/mnist-softmax  Softmax regression classifier for MNIST
    <exit 0>

    >>> run("guild operations mnist", ignore="FutureWarning")
    mnist/mnist-cnn:evaluate      Evaluate a trained CNN
    mnist/mnist-cnn:train         Train the CNN
    mnist/mnist-samples:prepare   Generate a set of sample MNIST images
    mnist/mnist-softmax:evaluate  Evaluate a trained softmax regression
    mnist/mnist-softmax:train     Train the softmax regression
    <exit 0>

    >>> run("guild resources mnist", ignore="FutureWarning")
    mnist/dataset  Yann Lecun's MNIST dataset in compressed IDX format
    <exit 0>
