#! /usr/bin/python
# -*- coding: utf-8 -*-


import os.path
import sys

path_to_script = os.path.dirname(os.path.abspath(__file__))
sys.path.append(os.path.join(path_to_script, "../extern/sed3"))

import numpy as np

import logging
logger = logging.getLogger(__name__)

import scipy
import scipy.ndimage

from . import dili

def select_labels(segmentation, labels, slab=None):
    """
    return ndimage with zeros and ones based on input labels

    :param segmentation: 3D ndimage
    :param labels: labels to select
    :param slab: dictionary{string_label: numeric_label}. Allow to use
    string labels if it is defined
    :return:
    """

    if slab is not None:
        labels = get_nlabels(slab, labels)

    if type(labels) not in (list, np.ndarray):
        labels = [labels]

    ds = np.zeros(segmentation.shape, np.bool)
    for lab in labels:
        dadd = (segmentation == lab)

        ds = ds | dadd

    return ds


def get_nlabels(slab, labels, labels_meta=None, return_mode="num", return_first=False):
    """
    Get one or more labels, create a new one if necessary and return its numeric value.

    Look at the get_nlabel function for more details.

    :param slab:
    :param labels:
    :param labels_meta:
    :param return_mode: "num" or "str" or "both". Both means (numlabel, strlabel).
    :param return_first: Return just first found label
    :return:
    """


    if type(labels) not in (list, np.ndarray):
        labels = [labels]
        labels_meta = [labels_meta]
        return_first = True

    if labels_meta is None:
        labels_meta = [None] * len(labels)

    nlabels = []
    for label, label_meta in zip(labels, labels_meta):
        nlab = get_nlabel(slab, label, label_meta, return_mode=return_mode)
        nlabels.append(nlab)

    if return_first:
        nlabels=nlabels[0]
    return nlabels


def get_nlabel(slab, label, label_meta=None, return_mode="num"):

    """
    Add label if it is necessery and return its numeric value.

    If "new" keyword is used and no other information is provided, the max + 1 label is created.
    If "new" keyword is used and additional numeric info is provided, the number is used also as a key.
    :param return_mode: Set requested label return type. "int", "num", "numeric" or "str" or "both".
    "both" means (numlabel, strlabel).
    :param label: string, number or "new"
    :param label_meta: string, number or "new
    :return:
    """
    numlabel = None
    strlabel = None
    if type(label) == str:
        if label_meta is None:
            if label not in slab.keys():
                free_numeric_label = np.max(list(slab.values())) + 1
                if label == "new":
                    label = str(free_numeric_label)
                slab[label] = free_numeric_label
                strlabel = label
                numlabel = slab[label]
            else:
                strlabel = label
                numlabel = slab[label]
        else:
            if label == "new":
                label = str(label_meta)
            update_slab(slab, label_meta, label)
            strlabel = label
            numlabel = label_meta
    else:
        # it is numeric
        if label_meta is None:
            if label not in list(slab.values()):
                update_slab(slab, label, str(label))
                strlabel = str(label)
            else:
                strlabel = dili.dict_find_key(slab, label)

            numlabel = label

        else:
            if label_meta == "new":
                label_meta = str(label)
            update_slab(slab, label, label_meta)
            strlabel = label_meta
            numlabel = label
            # return label

    if return_mode in ("num", "int", "numeric"):
        return numlabel
    elif return_mode == "str":
        return strlabel
    elif return_mode == "both":
        return numlabel, strlabel
    else:
        logger.error("Unknown return_mode: " + str(return_mode))


def update_slab(slab, numeric_label, string_label):
    """ Add label to segmentation label dictionary if it is not there yet.

    :param numeric_label:
    :param string_label:
    :return:
    """

    slab_tmp = {string_label: numeric_label}
    slab.update(slab_tmp)
    # slab = slab_tmp
    logger.debug('self.slab')
    logger.debug(str(slab))

def add_slab_label_carefully2(slab, numeric_label, string_label):
    """ Add label to slab if it is not there yet.

    :param numeric_label:
    :param string_label:
    :return:
    """
    # todo implement
    # if numeric_label in
    pass


def add_missing_labels(segmentation, slab):
    labels = np.unique(segmentation)
    get_nlabels(slab, labels)



def add_slab_label_carefully(slab, numeric_label, string_label):
    """ Add label to slab if it is not there yet.

    :param numeric_label:
    :param string_label:
    :return:
    """
    slab_tmp = {string_label: numeric_label}
    slab_tmp.update(slab)
    slab = slab_tmp
    logger.debug('self.slab')
    logger.debug(str(slab))

def add_missing_labels(segmentation, slab):
    labels = np.unique(segmentation)
    get_nlabels(slab, labels)



class SparseMatrix():
    def __init__(self, ndarray):
        self.coordinates = ndarray.nonzero()
        self.shape = ndarray.shape
        self.values = ndarray[self.coordinates]
        self.dtype = ndarray.dtype
        self.sparse = True

    def todense(self):
        dense = np.zeros(self.shape, dtype=self.dtype)
        dense[self.coordinates[:]] = self.values
        return dense


def isSparseMatrix(obj):
    if obj.__class__.__name__ == 'SparseMatrix':
        return True
    else:
        return False


# import sed3

def manualcrop(data):  # pragma: no cover

    try:
        from pysegbase import seed_editor_qt
    except:
        logger.warning("Deprecated of pyseg_base as submodule")
        import seed_editor_qt

    pyed = seed_editor_qt.QTSeedEditor(data, mode='crop')
    pyed.exec_()
    # pyed = sed3.sed3(data)
    # pyed.show()
    nzs = pyed.seeds.nonzero()
    crinfo = [
        [np.min(nzs[0]), np.max(nzs[0])],
        [np.min(nzs[1]), np.max(nzs[1])],
        [np.min(nzs[2]), np.max(nzs[2])],
        ]
    data = crop(data, crinfo)
    return data, crinfo


def crop(data, crinfo):
    """
    Crop the data.

    crop(data, crinfo)

    :param crinfo: min and max for each axis - [[minX, maxX], [minY, maxY], [minZ, maxZ]]

    """
    crinfo = fix_crinfo(crinfo)
    return data[
        __int_or_none(crinfo[0][0]):__int_or_none(crinfo[0][1]),
        __int_or_none(crinfo[1][0]):__int_or_none(crinfo[1][1]),
        __int_or_none(crinfo[2][0]):__int_or_none(crinfo[2][1])
        ]


def __int_or_none(number):
    if number is not None:
        number = int(number)
    return number


def combinecrinfo(crinfo1, crinfo2):
    """
    Combine two crinfos. First used is crinfo1, second used is crinfo2.
    """
    crinfo1 = fix_crinfo(crinfo1)
    crinfo2 = fix_crinfo(crinfo2)

    crinfo = [
        [crinfo1[0][0] + crinfo2[0][0], crinfo1[0][0] + crinfo2[0][1]],
        [crinfo1[1][0] + crinfo2[1][0], crinfo1[1][0] + crinfo2[1][1]],
        [crinfo1[2][0] + crinfo2[2][0], crinfo1[2][0] + crinfo2[2][1]]
        ]

    return crinfo


def crinfo_from_specific_data(data, margin=0):
    """
    Create crinfo of minimum orthogonal nonzero block in input data.

    :param data: input data
    :param margin: add margin to minimum block
    :return:
    """
    # hledáme automatický ořez, nonzero dá indexy
    logger.debug('crinfo')
    logger.debug(str(margin))
    nzi = np.nonzero(data)
    logger.debug(str(nzi))

    if np.isscalar(margin):
        margin = [margin] * 3

    x1 = np.min(nzi[0]) - margin[0]
    x2 = np.max(nzi[0]) + margin[0] + 1
    y1 = np.min(nzi[1]) - margin[0]
    y2 = np.max(nzi[1]) + margin[0] + 1
    z1 = np.min(nzi[2]) - margin[0]
    z2 = np.max(nzi[2]) + margin[0] + 1

# ošetření mezí polí
    if x1 < 0:
        x1 = 0
    if y1 < 0:
        y1 = 0
    if z1 < 0:
        z1 = 0

    if x2 > data.shape[0]:
        x2 = data.shape[0] - 1
    if y2 > data.shape[1]:
        y2 = data.shape[1] - 1
    if z2 > data.shape[2]:
        z2 = data.shape[2] - 1

# ořez
    crinfo = [[x1, x2], [y1, y2], [z1, z2]]
    return crinfo


def uncrop(data, crinfo, orig_shape, resize=False):
    """

    :param data: input data
    :param crinfo: array with minimum and maximum index along each axis
        [[minX. maxX],[minY, maxY],[minZ, maxZ]]
    :param orig_shape: shape of uncropped image
    :param resize:
    :return:
    """

    crinfo = fix_crinfo(crinfo)
    data_out = np.zeros(orig_shape, dtype=data.dtype)

    # print 'uncrop ', crinfo
    # print orig_shape
    # print data.shape
    if resize:
        data = resize_to_shape(data, crinfo[:, 1] - crinfo[:, 0])

    startx = np.round(crinfo[0][0]).astype(int)
    starty = np.round(crinfo[1][0]).astype(int)
    startz = np.round(crinfo[2][0]).astype(int)

    data_out[
        # np.round(crinfo[0][0]).astype(int):np.round(crinfo[0][1]).astype(int)+1,
        # np.round(crinfo[1][0]).astype(int):np.round(crinfo[1][1]).astype(int)+1,
        # np.round(crinfo[2][0]).astype(int):np.round(crinfo[2][1]).astype(int)+1
        startx:startx + data.shape[0],
        starty:starty + data.shape[1],
        startz:startz + data.shape[2]
        ] = data

    return data_out


def fix_crinfo(crinfo, to='axis'):
    """
    Function recognize order of crinfo and convert it to proper format.
    """

    crinfo = np.asarray(crinfo)
    if crinfo.shape[0] == 2:
        crinfo = crinfo.T

    return crinfo



def get_one_biggest_object(data):
    """ Return biggest object """
    lab, num = scipy.ndimage.label(data)
    # print ("bum = "+str(num))

    maxlab = max_area_index(lab, num)

    data = (lab == maxlab)
    return data


def max_area_index(labels, num):
    """
    Return index of maxmum labeled area
    """
    mx = 0
    mxi = -1
    for l in range(1, num + 1):
        mxtmp = np.sum(labels == l)
        if mxtmp > mx:
            mx = mxtmp
            mxi = l

    return mxi


def resize_to_shape(data, shape, zoom=None, mode='nearest', order=0):
    """Resize input data to specific shape.

    :param data: input 3d array-like data
    :param shape: shape of output data
    :param zoom: zoom is used for back compatibility
    :mode: default is 'nearest'
    """
    # @TODO remove old code in except part

    if np.array_equal(data.shape, shape):
        return data

    try:
        # rint 'pred vyjimkou'
        # aise Exception ('test without skimage')
        # rint 'za vyjimkou'
        import skimage
        import skimage.transform
        # Now we need reshape  seeds and segmentation to original size

        segm_orig_scale = skimage.transform.resize(
            data, shape, order=order,
            preserve_range=True
        )

        segmentation = segm_orig_scale
        logger.debug('resize to orig with skimage')
    except Exception:
        logger.warning("Resize by scipy will be removed in the future")
        import scipy
        import scipy.ndimage
        dtype = data.dtype
        if zoom is None:
            zoom = np.asarray(data.shape).astype(np.double) / shape

        segm_orig_scale = scipy.ndimage.zoom(
            data,
            1.0 / zoom,
            mode=mode,
            order=order
        ).astype(dtype)
        logger.debug('resize to orig with scipy.ndimage')

        # @TODO odstranit hack pro oříznutí na stejnou velikost
        # v podstatě je to vyřešeno, ale nechalo by se to dělat elegantněji v zoom
        # tam je bohužel patrně bug
        # rint 'd3d ', self.data3d.shape
        # rint 's orig scale shape ', segm_orig_scale.shape
        shp = [
            np.min([segm_orig_scale.shape[0], shape[0]]),
            np.min([segm_orig_scale.shape[1], shape[1]]),
            np.min([segm_orig_scale.shape[2], shape[2]]),
        ]
        # elf.data3d = self.data3d[0:shp[0], 0:shp[1], 0:shp[2]]
        # mport ipdb; ipdb.set_trace() # BREAKPOINT

        segmentation = np.zeros(shape, dtype=dtype)
        segmentation[
        0:shp[0],
        0:shp[1],
        0:shp[2]] = segm_orig_scale[0:shp[0], 0:shp[1], 0:shp[2]]

        del segm_orig_scale
    return segmentation

def resize_to_mm(data3d, voxelsize_mm, new_voxelsize_mm, mode='nearest', order=1):
    """
    Function can resize data3d or segmentation to specifed voxelsize_mm
    :new_voxelsize_mm: requested voxelsize. List of 3 numbers, also
        can be a string 'orig', 'orig*2' and 'orig*4'.

    :voxelsize_mm: size of voxel
    :mode: default is 'nearest'
    """

    if new_voxelsize_mm is 'orig':
        new_voxelsize_mm = np.array(voxelsize_mm)

    elif new_voxelsize_mm is 'orig*2':
        new_voxelsize_mm = np.array(voxelsize_mm) * 2
    elif new_voxelsize_mm is 'orig*4':
        new_voxelsize_mm = np.array(voxelsize_mm) * 4
        # vx_size = np.array(metadata['voxelsize_mm']) * 4

    zoom = voxelsize_mm / (1.0 * np.array(new_voxelsize_mm))
    # data3d_res = scipy.ndimage.zoom(
    #     data3d,
    #     zoom,
    #     mode=mode,
    #     order=order
    # ).astype(data3d.dtype)

    # probably better implementation
    new_shape = data3d.shape * zoom
    import skimage.transform
    # Now we need reshape  seeds and segmentation to original size

    data3d_res2 = skimage.transform.resize(
        data3d, new_shape, order=order,
        mode=mode,
        preserve_range=True
    ).astype(data3d.dtype)

    return data3d_res2
