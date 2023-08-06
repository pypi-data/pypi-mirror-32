# Transcendental (π)

Artificial Intelligence image recognition.

Learn from input data and generate model, predict from previous generated model

## Disclaimer

This is testing project purpose so it's currently in development and features
are minimal don't juge me :)

## Features

- Recognize cats
- Recognize dogs

## Install

```shell
$ pip install transcendental
```

## Usage

### trans-predict

This command allow you to recognize an image passed in argument.

Predict take exactly one argument who is the image to predict
```shell
$ trans-predict --help 
usage: trans-predict [-h] image

Recognize cat and dogs

positional arguments:
  image       an image to analyze

  optional arguments:
    -h, --help  show this help message and exit

$ trans-predict ./image.png
```

### trans-train

This command allow you to train your IA.

This command generate models automaticaly used by predict at launch

Currently in development!

### Requires

- python 2.7+
- keras
- tensorflow
- Pillow
- h5py

## Resources
- https://becominghuman.ai/building-an-image-classifier-using-deep-learning-in-python-totally-from-a-beginners-perspective-be8dbaf22dd8
- https://drive.google.com/drive/folders/1XaFM8BJFligrqeQdE-_5Id0V_SubJAZe
- https://machinelearningmastery.com/save-load-keras-deep-learning-models/
- https://www.quora.com/How-do-artificial-neural-networks-work
- http://torch.ch/blog/2015/11/13/gan.html
