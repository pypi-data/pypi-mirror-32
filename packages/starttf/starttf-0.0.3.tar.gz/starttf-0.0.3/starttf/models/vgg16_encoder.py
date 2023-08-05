import tensorflow as tf

from starttf.layers.caffe_tensorflow import Network

"""
This model is implemented by Michael Fuerst and supports caffe converted weights.

Convert weights using: https://github.com/ethereon/caffe-tensorflow
Or download weights: https://www.cs.toronto.edu/~frossard/vgg16/vgg16_weights.npz
"""


class Vgg16Encoder(Network):
    def setup(self):
        (self.feed('data')
             .conv(3, 3, 64, 1, 1, name='conv1_1')
             .conv(3, 3, 64, 1, 1, name='conv1_2')
             .max_pool(2, 2, 2, 2, name='pool1')
             .conv(3, 3, 128, 1, 1, name='conv2_1')
             .conv(3, 3, 128, 1, 1, name='conv2_2')
             .max_pool(2, 2, 2, 2, name='pool2')
             .conv(3, 3, 256, 1, 1, name='conv3_1')
             .conv(3, 3, 256, 1, 1, name='conv3_2')
             .conv(3, 3, 256, 1, 1, name='conv3_3')
             .max_pool(2, 2, 2, 2, name='pool3')
             .conv(3, 3, 512, 1, 1, name='conv4_1')
             .conv(3, 3, 512, 1, 1, name='conv4_2')
             .conv(3, 3, 512, 1, 1, name='conv4_3')
             .max_pool(2, 2, 2, 2, name='pool4')
             .conv(3, 3, 512, 1, 1, name='conv5_1')
             .conv(3, 3, 512, 1, 1, name='conv5_2')
             .conv(3, 3, 512, 1, 1, name='conv5_3')
             .max_pool(2, 2, 2, 2, name='pool5'))


def create_model(input_tensor, mode, hyper_params):
    """
    A vgg16 encoder network taken which can load caffe converted weights.

    Either convert them yourself or download them from a third party.

    Convert weights using: https://github.com/ethereon/caffe-tensorflow
    Or download weights: https://www.cs.toronto.edu/~frossard/vgg16/vgg16_weights.npz

    :param input_tensor: The input tensor dict containing a "image" rgb tensor.
    :param mode: Execution mode as a tf.estimator.ModeKeys
    :param hyper_params: The hyper param file. "vgg16" : {"encoder_only": Boolean}
    :return: A dictionary containing all output tensors.
    """
    with tf.variable_scope("vgg16") as scope:
        image = tf.cast(input_tensor["image"], dtype=tf.float32, name="input/cast")
        mean = tf.constant([123.68, 116.779, 103.939], dtype=tf.float32, shape=[1, 1, 1, 3], name='img_mean')
        normalized_image = image - mean

        network = Vgg16Encoder({"data": normalized_image}, weight_file=hyper_params.vgg16.weight_file, ignore_missing=True)
        model = dict(network.layers)
        model["image"] = image
        model["normalized_image"] = normalized_image
    return model
