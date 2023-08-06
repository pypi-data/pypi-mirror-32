"""Utility functions."""
from operator import itemgetter

from matplotlib import pyplot as plt

from prysm import mathops as m


def is_odd(int):
    """Determine if an interger is odd using binary operations.

    Parameters
    ----------
    int : `int`
        an integer

    Returns
    -------
    `bool`
        true if odd, False if even

    """
    return int & 0x1


def is_power_of_2(value):
    """Check if a value is a power of 2 using binary operations.

    Parameters
    ----------
    value : `number`
        value to check

    Returns
    -------
    `bool`
        true if the value is a power of two, False if the value is no

    Notes
    -----
    c++ inspired implementation, see SO:
    https://stackoverflow.com/questions/29480680/finding-if-a-number-is-a-power-of-2-using-recursion

    """
    if value is 1:
        return False
    else:
        return bool(value and not value & (value - 1))


def pupil_sample_to_psf_sample(pupil_sample, num_samples, wavelength, efl):
    """Convert pupil sample spacing to PSF sample spacing.

    Parameters
    ----------
    pupil_sample : `float`
        sample spacing in the pupil plane
    num_samples : `int`
        number of samples present in both planes (must be equal)
    wavelength : `float`
        wavelength of light, in microns
    efl : `float`
        effective focal length of the optical system in mm

    Returns
    -------
    `float`
        the sample spacing in the PSF plane

    """
    return (wavelength * efl * 1e3) / (pupil_sample * num_samples)


def psf_sample_to_pupil_sample(psf_sample, num_samples, wavelength, efl):
    """Convert PSF sample spacing to pupil sample spacing.

    Parameters
    ----------
    psf_sample : `float`
        sample spacing in the PSF plane
    num_samples : `int`
        number of samples present in both planes (must be equal)
    wavelength : `float`
        wavelength of light, in microns
    efl : `float`
        effective focal length of the optical system in mm

    Returns
    -------
    `float`
        the sample spacing in the pupil plane

    """
    return (wavelength * efl * 1e3) / (psf_sample * num_samples)


def correct_gamma(img, encoding=2.2):
    """Apply an inverse gamma curve to image data that linearizes the given encoding.

    Parameters
    ----------
    img : `numpy.ndarray`
        array of image data, floats avoid quantization error
    encoding : `float`
        gamma to encode that data to (1.0 is linear)

    Returns
    -------
    `numpy.ndarray`
        Array of corrected data

    """
    if encoding is 1:
        return img
    else:
        return img ** (1 / float(encoding))


def fold_array(array, axis=1):
    """Fold an array in half over the given axis and averages.

    Parameters
    ----------
    array : `numpy.ndarray`
        ndarray
    axis : `int`, optional
        axis to fold over

    Returns
    -------
    `numpy.ndarray`
        folded array

    """
    xs, ys = array.shape
    if axis is 1:
        xh = xs // 2
        left_chunk = array[:, :xh]
        right_chunk = array[:, xh:]
        folded_array = m.concatenate((right_chunk[:, :, m.newaxis],
                                       m.flip(m.flip(left_chunk, axis=1),
                                               axis=0)[:, :, m.newaxis]),
                                      axis=2)
    else:
        yh = ys // 2
        top_chunk = array[:yh, :]
        bottom_chunk = array[yh:, :]
        folded_array = m.concatenate((bottom_chunk[:, :, m.newaxis],
                                       m.flip(m.flip(top_chunk, axis=1),
                                               axis=0)[:, :, m.newaxis]),
                                      axis=2)
    return folded_array.mean(axis=2)


def share_fig_ax(fig=None, ax=None, numax=1, sharex=False, sharey=False):
    """Reurns the given figure and/or axis if given one.  If they are None, creates a new fig/ax.

    Parameters
    ----------
    fig : `matplotlib.figure.Figure`, optional
        figure
    ax : `matplotlib.axes.Axis`
        axis or array of axes
    numax : `int`
        number of axes in the desired figure, 1 for most plots, 3 for plot_fourier_chain
    sharex : `bool`, optional
        whether to share the x axis
    sharey : `bool`, optional
        whether to share the y axis

    Returns
    -------
    `matplotlib.figure.Figure`
        A figure object
    `matplotlib.axes.Axis`
        An axis object

    """
    if fig is None and ax is None:
        fig, ax = plt.subplots(nrows=1, ncols=numax, sharex=sharex, sharey=sharey)
    elif ax is None:
        ax = fig.gca()

    return fig, ax


def rms(array):
    """Return the RMS value of the valid elements of an array.

    Parameters
    ----------
    array : `numpy.ndarray`
        array of values

    Returns
    -------
    `float`
        RMS of the array

    """
    non_nan = m.isfinite(array)
    return m.sqrt((array[non_nan] ** 2).mean())


def guarantee_array(variable):
    """Guarantee that a varaible is a numpy ndarray and supports -, *, +, and other operators.

    Parameters
    ----------
    variable : `number` or `numpy.ndarray`
        variable to coalesce

    Returns
    -------
    `object`
        an object that  supports * / and other operations with ndarrays

    Raises
    ------
    ValueError
        non-numeric type

    """
    if type(variable) in [float, m.ndarray, m.int32, m.int64, m.float32, m.float64, m.complex64, m.complex128]:
        return variable
    elif type(variable) is int:
        return float(variable)
    elif type(variable) is list:
        return m.asarray(variable)
    else:
        raise ValueError(f'variable is of invalid type {type(variable)}')


def ecdf(x):
    """Compute the empirical cumulative distribution function of a dataset.

    Parameters
    ----------
    x : `iterable`
        Data

    Returns
    -------
    xs : `numpy.ndarray`
        sorted data
    ys : `numpy.ndarray`
        cumulative distribution function of the data

    """
    xs = m.sort(x)
    ys = m.arange(1, len(xs) + 1) / float(len(xs))
    return xs, ys


def sort_xy(x, y):
    """Sorts a pair of x and y iterables, returning arrays in order of ascending x.

    Parameters
    ----------
    x : `iterable`
        a list, numpy ndarray, or other iterable to sort by
    y : `iterable`
        a list, numpy ndarray, or other iterable that is y=f(x)

    Returns
    -------
    sorted_x : iterable
        an iterable containing the sorted x elements
    sorted_y : iterable
        an iterable containing the sorted y elements

    """
    # zip x and y, sort by the 0th element (x) of each tuple in zip()
    _ = sorted(zip(x, y), key=itemgetter(0))
    sorted_x, sorted_y = zip(*_)
    return sorted_x, sorted_y
