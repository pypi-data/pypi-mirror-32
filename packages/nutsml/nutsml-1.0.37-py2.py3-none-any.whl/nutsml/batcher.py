"""
.. module:: batcher
   :synopsis: Collecting samples in mini-batches for GPU-based training.
"""
import warnings
import numpy as np
import nutsml.imageutil as ni

from nutsflow.base import Nut
from nutsflow.iterfunction import take, PrefetchIterator


def build_number_batch(numbers, dtype):
    """
    Return numpy array with given dtype for given numbers.

    >>> numbers = (1, 2, 3, 1)
    >>> build_number_batch(numbers, 'uint8')
    array([1, 2, 3, 1], dtype=uint8)

    :param iterable number numbers: Numbers to create batch from
    :param numpy data type dtype: Data type of batch, e.g. 'uint8'
    :return: Numpy array for numbers
    :rtype: numpy.array
    """
    return np.array(numbers, dtype=dtype)


def build_one_hot_batch(class_ids, dtype, num_classes):
    """
    Return one hot vectors for class ids.

    >>> class_ids = [0, 1, 2, 1]
    >>> build_one_hot_batch(class_ids, 'uint8', 3)
    array([[1, 0, 0],
           [0, 1, 0],
           [0, 0, 1],
           [0, 1, 0]], dtype=uint8)

    :param iterable class_ids: Class indices in {0, ..., num_classes-1}
    :param numpy data type dtype: Data type of batch, e.g. 'uint8'
    :param num_classes: Number of classes
    :return: One hot vectors for class ids.
    :rtype: numpy.array
    """
    class_ids = np.array(class_ids, dtype=np.uint16)
    return np.eye(num_classes, dtype=dtype)[class_ids]


def build_vector_batch(vectors, dtype):
    """
    Return batch of vectors.

    >>> from nutsml.datautil import shapestr
    >>> vectors = [np.array([1,2,3]), np.array([2, 3, 4])]
    >>> batch = build_vector_batch(vectors, 'uint8')
    >>> shapestr(batch)
    '2x3'

    >>> batch
    array([[1, 2, 3],
           [2, 3, 4]], dtype=uint8)

    :param iterable vectors: Numpy row vectors
    :param numpy data type dtype: Data type of batch, e.g. 'uint8'
    :return: vstack of vectors
    :rtype: numpy.array
    """
    if not len(vectors):
        raise ValueError('No vectors to build batch!')
    return np.vstack(vectors).astype(dtype)


def build_tensor_batch(tensors, dtype, axes=None):
    """
    Return batch of tensors.

    >>> from nutsml.datautil import shapestr
    >>> tensors = [np.zeros((2, 3)), np.ones((2, 3))]
    >>> batch = build_tensor_batch(tensors, 'uint8')
    >>> shapestr(batch)
    '2x2x3'

    >>> print(batch)
    [[[0 0 0]
      [0 0 0]]
    <BLANKLINE>
     [[1 1 1]
      [1 1 1]]]


    >>> batch = build_tensor_batch(tensors, 'uint8', axes = (1, 0))
    >>> shapestr(batch)
    '2x3x2'

    >>> print(batch)
    [[[0 0]
      [0 0]
      [0 0]]
    <BLANKLINE>
     [[1 1]
      [1 1]
      [1 1]]]

    :param iterable tensors: Numpy tensors
    :param numpy data type dtype: Data type of batch, e.g. 'uint8'
    :param tuple|None axes: axes order, e.g. to move a channel axis to the
      last position. (see numpy transpose for details)
    :return: stack of tensors, with batch axis first.
    :rtype: numpy.array
    """
    if not len(tensors):
        raise ValueError('No tensors to build batch!')
    if axes:
        tensors = [np.transpose(t, axes) for t in tensors]
    return np.stack(tensors).astype(dtype)


def build_image_batch(images, dtype, channelfirst=False):
    """
    Return batch of images.

    If images have no channel a channel axis is added. For channelfirst=True
    it will be added/moved to front otherwise the channel comes last.
    All images in batch will have a channel axis. Batch is of shape
    (n, c, h, w) or (n, h, w, c) depending on channelfirst, where n is
    the number of images in the batch.

    >>> from nutsml.datautil import shapestr
    >>> images = [np.zeros((2, 3)), np.ones((2, 3))]
    >>> batch = build_image_batch(images, 'uint8', True)
    >>> shapestr(batch)
    '2x1x2x3'

    >>> batch
    array([[[[0, 0, 0],
             [0, 0, 0]]],
    <BLANKLINE>
    <BLANKLINE>
           [[[1, 1, 1],
             [1, 1, 1]]]], dtype=uint8)

    :param numpy array images: Images to batch. Must be of shape (w,h,c)
           or (w,h). Gray-scale with channel is fine (w,h,1) and also
           alpha channel is fine (w,h,4).
    :param numpy data type dtype: Data type of batch, e.g. 'uint8'
    :param bool channelfirst: If True, channel is added/moved to front.
    :return: Image batch with shape (n, c, h, w) or (n, h, w, c).
    :rtype: np.array
    """

    def _targetshape(image):
        shape = image.shape
        return (shape[0], shape[1], 1) if image.ndim == 2 else shape

    n = len(images)
    if not n:
        raise ValueError('No images to build batch!')
    h, w, c = _targetshape(images[0])  # shape of first(=all) images
    if c > w or c > h:
        raise ValueError('Channel not at last axis: ' + str((h, w, c)))
    batch = np.empty((n, c, h, w) if channelfirst else (n, h, w, c))
    for i, image in enumerate(images):
        image = ni.add_channel(image, channelfirst)
        if image.shape != batch.shape[1:]:
            raise ValueError('Images vary in shape: ' + str(image.shape))
        batch[i, :, :, :] = image
    return batch.astype(dtype)


class BuildBatch(Nut):
    """
    Build batches for GPU-based neural network training.
    """

    def __init__(self, batchsize, prefetch=1, fmt=None):
        """
        iterable >> BuildBatch(batchsize, prefetch=1)

        Take samples in iterable, extract specified columns, convert
        column data to numpy arrays of various types, aggregate converted
        samples into a batch.

        >>> from nutsflow import Collect
        >>> numbers = [4.1, 3.2, 1.1]
        >>> images = [np.zeros((5, 3)), np.ones((5, 3)) , np.ones((5, 3))]
        >>> class_ids = [1, 2, 1]
        >>> samples = zip(numbers, images, class_ids)

        >>> build_batch = (BuildBatch(2, prefetch=0)
        ...                .input(0, 'number', float)
        ...                .input(1, 'image', np.uint8, True)
        ...                .output(2, 'one_hot', np.uint8, 3))
        >>> batches = samples >> build_batch >> Collect()

        Sample columns can be ignored or reused. Assuming an autoencoder, one
        might which to use the sample image as input and output:

        >>> build_batch = (BuildBatch(2, prefetch=0)
        ...                .input(1, 'image', np.uint8, True)
        ...                .output(1, 'image', np.uint8, True))
        >>> batches = samples >> build_batch >> Collect()

        A batch typically is of the format [[inputs], [outputs]], e.g. in the
        firs case above [[number, image], [one_hot]], where each of the columns
        is a Numpy array. If a different structure is required a format function
        can be specified. The following example shows how to build a batch as
        as tuple of columns instead of a list.

        >>> samples = zip(images, images, class_ids)
        >>> build_batch = (BuildBatch(2, prefetch=0,
        ...                fmt=lambda b: (b[0], b[1], b[2]) )
        ...                .input(0, 'image', np.uint8, True)
        ...                .input(1, 'image', np.uint8, True)
        ...                .input(2, 'one_hot', np.uint8, 3))
        >>> batches = samples >> build_batch >> Collect()

        :param int batchsize: Size of batch = number of rows in batch.
        :param int prefetch: Number of batches to prefetch. This speeds up
           GPU based training, since one batch is built on CPU while the
           another is procesed on the GPU.
        :param function|None fmt: Function to format output.
        """
        self.batchsize = batchsize
        self.fmt = fmt
        self.prefetch = prefetch
        self.colspecs = []
        self.builder = {'image': build_image_batch,
                        'number': build_number_batch,
                        'vector': build_vector_batch,
                        'tensor': build_tensor_batch,
                        'one_hot': build_one_hot_batch}

    def by(self, col, name, *args, **kwargs):
        """
        Specify and add batch columns to create batch.

        DEPRECATED: Use input() and output() instead.

        :param int col: column of the sample to extract and to create a
          batch column from.
        :param string name: Name of the column function to apply to create
            a batch column, e.g. 'image'
            See the following functions for more details:
            'image': nutsflow.batcher.build_image_batch
            'number': nutsflow.batcher.build_number_batch
            'vector': nutsflow.batcher.build_vector_batch
            'tensor': nutsflow.batcher.build_tensor_batch
            'one_hot': nutsflow.batcher.build_one_hot_batch
        :param args args: Arguments for column function, e.g. dtype
        :param kwargs kwargs: Keyword arguments for column function
        :return: instance of BuildBatch
        :rtype: BuildBatch
        """
        warnings.warn("'by()' is deprecated. Use input() or output() instead.",
                      DeprecationWarning)
        self.colspecs.append((col, name, True, args, kwargs))
        return self

    def input(self, col, name, *args, **kwargs):
        """
        Specify and add input columns for batch to create

        :param int col: column of the sample to extract and to create a
          batch input column from.
        :param string name: Name of the column function to apply to create
            a batch column, e.g. 'image'
            See the following functions for more details:
            'image': nutsflow.batcher.build_image_batch
            'number': nutsflow.batcher.build_number_batch
            'vector': nutsflow.batcher.build_vector_batch
            'tensor': nutsflow.batcher.build_tensor_batch
            'one_hot': nutsflow.batcher.build_one_hot_batch
        :param args args: Arguments for column function, e.g. dtype
        :param kwargs kwargs: Keyword arguments for column function
        :return: instance of BuildBatch
        :rtype: BuildBatch
        """
        self.colspecs.append((col, name, True, args, kwargs))
        return self

    def output(self, col, name, *args, **kwargs):
        """
        Specify and add output columns for batch to create

        :param int col: column of the sample to extract and to create a
          batch output column from.
        :param string name: Name of the column function to apply to create
            a batch column, e.g. 'image'
            See the following functions for more details:
            'image': nutsflow.batcher.build_image_batch
            'number': nutsflow.batcher.build_number_batch
            'vector': nutsflow.batcher.build_vector_batch
            'tensor': nutsflow.batcher.build_tensor_batch
            'one_hot': nutsflow.batcher.build_one_hot_batch
        :param args args: Arguments for column function, e.g. dtype
        :param kwargs kwargs: Keyword arguments for column function
        :return: instance of BuildBatch
        :rtype: BuildBatch
        """
        self.colspecs.append((col, name, False, args, kwargs))
        return self

    def _batch_generator(self, iterable):
        """Return generator over batches for given iterable of samples"""
        while 1:
            batchsamples = list(take(iterable, self.batchsize))
            if not batchsamples:
                break
            cols = list(zip(*batchsamples))  # flip rows to cols
            batch = [[], []]  # in, out columns of batch
            for colspec in self.colspecs:
                col, func, isinput, args, kwargs = colspec
                if not func in self.builder:
                    raise ValueError('Invalid builder: ' + func)
                coldata = self.builder[func](cols[col], *args, **kwargs)
                batch[0 if isinput else 1].append(coldata)
            if not batch[1]:  # no output (prediction phase)
                batch = batch[0]  # flatten and take only inputs
            if self.fmt:  # format function for batch given
                batch = self.fmt(batch)
            yield batch

    def __rrshift__(self, iterable):
        """
        Convert samples in iterable into mini-batches.

        Structure of output depends on fmt function used. If None
        output is a list of np.arrays

        :param iterable iterable: Iterable over samples.
        :return: Mini-batches
        :rtype: list of np.array if fmt=None
        """
        prefetch = self.prefetch
        batch_gen = self._batch_generator(iter(iterable))
        return PrefetchIterator(batch_gen, prefetch) if prefetch else batch_gen
