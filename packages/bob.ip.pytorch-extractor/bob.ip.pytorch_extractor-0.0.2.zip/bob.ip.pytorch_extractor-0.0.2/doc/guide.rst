==========
User guide
==========

This package has been done with the intent of using it in conjunction with
both ``bob.learn.pytorch`` and ``bob.{bio, pad}.face`` to run face verification
or presentation attack detection experiments.

In particular, the goal here is to extract features from a face image using CNN models 
trained with PyTorch

For this purpose, you can specify your feature extractor in configuration
file to be used together with either with the ``verifiy.py`` script from ``bob.bio.base``,
of with the ``spoof.py`` script from ``bob.pad.base``.


A concrete example
------------------

Imagine that you trained the CASIA-Net model (using ``bob.learn.pytorch``), and you would 
like to use the penultimate layer as a feature to encode identity. 

Your extractor should be defined this way in the configuration file:

.. code:: python

  extractor = from bob.ip.pytorch_extractor import CasiaNetExtractor

  _model = 'path/to/your/model.pth'
  _num_classes = 10575
  extractor = CasiaNetExtractor(_model, _num_classes)

Note that the number of classes is irrelevant here, but is required to build the 
network (before loading it).

You can easily implement your own extractor based on your own network too. Just have
a look at the code in ``bob/ip/pytorch_extractor``.

