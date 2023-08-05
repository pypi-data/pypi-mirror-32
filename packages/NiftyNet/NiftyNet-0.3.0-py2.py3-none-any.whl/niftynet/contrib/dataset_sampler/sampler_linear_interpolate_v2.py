# -*- coding: utf-8 -*-
"""
Generating samples by linearly combining two input images.
"""
from __future__ import absolute_import, print_function, division

import numpy as np
import tensorflow as tf

from niftynet.contrib.dataset_sampler.image_window_dataset import \
    ImageWindowDataset
from niftynet.engine.image_window import N_SPATIAL, LOCATION_FORMAT


class LinearInterpolateSampler(ImageWindowDataset):
    """
    This class reads two feature vectors from files (often generated
    by running feature extractors on images in advance)
    and returns n linear combinations of the vectors.
    The coefficients are generated by::

        np.linspace(0, 1, n_interpolations)
    """

    def __init__(self,
                 reader,
                 window_sizes,
                 batch_size=10,
                 n_interpolations=10,
                 queue_length=10,
                 name='linear_interpolation_sampler'):
        ImageWindowDataset.__init__(
            self,
            reader,
            window_sizes=window_sizes,
            batch_size=batch_size,
            queue_length=queue_length,
            shuffle=False,
            epoch=1,
            name=name)
        self.n_interpolations = n_interpolations
        # only try to use the first spatial shape available
        image_spatial_shape = list(self.reader.shapes.values())[0][:3]
        self.window.set_spatial_shape(image_spatial_shape)
        tf.logging.info(
            "initialised linear interpolation sampler %s ", self.window.shapes)
        assert not self.window.has_dynamic_shapes, \
            "dynamic shapes not supported, please specify " \
            "spatial_window_size = (1, 1, 1)"

    def layer_op(self, *_unused_args, **_unused_kwargs):
        """
        This function first reads two vectors, and interpolates them
        with self.n_interpolations mixing coefficients.

        Location coordinates are set to ``np.ones`` for all the vectors.
        """
        while True:
            image_id_x, data_x, _ = self.reader(idx=None, shuffle=False)
            image_id_y, data_y, _ = self.reader(idx=None, shuffle=True)
            if not data_x or not data_y:
                break
            if image_id_x == image_id_y:
                continue
            embedding_x = data_x[self.window.names[0]]
            embedding_y = data_y[self.window.names[0]]

            steps = np.linspace(0, 1, self.n_interpolations)
            for (_, mixture) in enumerate(steps):
                output_vector = \
                    embedding_x * mixture + embedding_y * (1 - mixture)
                coordinates = np.ones((N_SPATIAL * 2 + 1), dtype=np.int32)
                coordinates[0:2] = [image_id_x, image_id_y]
                output_dict = {}
                for name in self.window.names:
                    coordinates_key = LOCATION_FORMAT.format(name)
                    image_data_key = name
                    output_dict[coordinates_key] = coordinates
                    output_dict[image_data_key] = output_vector
                yield output_dict
