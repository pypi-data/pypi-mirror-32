import os
import csv
from bobo.util import remkdir, isstring, filetail, filebase, quietprint, tolist, islist
from bobo.viset import janus_cs2
import itertools
from bobo.geometry import BoundingBox
from janus.template import GalleryTemplate


def newkey(im, version=3):
    if version in ['4a', 4.1]:
        return os.path.join('Janus_CS2_frontalized_3d_v4a', 'frontalized', '%s_frontalized_wsym.png' % (im.attributes['SIGHTING_ID']))
    elif version in ['4b', 4.2]:
        return os.path.join('Janus_CS2_frontalized_3d_v4b', 'frontalized', '%s_frontalized_wsym.png' % (im.attributes['SIGHTING_ID']))
    else:
        raise ValueError('Undefined 3D frontalization version "%s"' % str(version))

def canonicalize(im, version=3):
    imf = im.rekey(newkey(im, version))
    if imf.isvalid():
        imf.bbox = BoundingBox(centroid=(250,250), height=500, width=350)
        return imf
    else:
        return None

def iscanonicalized(im):    
    return im.rekey(newkey(im)).isvalid()
    
def rdd(sparkContext, dataset='train', split=None, version=3):
    """RENAME ME: this should be rdd_as_media"""
    rdd = janus_cs2.rdd(sparkContext, dataset, split)
    return rdd.map(lambda im: canonicalize(im, version) if im.mediatype()=='image' else im.mapFrames(lambda fr: canonicalize(fr)))

def rdd_as_templates(sparkContext, dataset='train', split=None, version=3):
    rdd = janus_cs2.rdd_as_templates(sparkContext, dataset, split)
    return rdd.map(lambda tmpl: GalleryTemplate(media=[canonicalize(im, version) if canonicalize(im, version) is not None else im for im in tmpl], mediakeys=tmpl.mediakeys))    

def rdd_as_frontalized(sparkContext, dataset='train', split=None, version=3):
    """RENAME ME: this should be rdd_as_images"""
    rdd = janus_cs2.rdd_as_templates(sparkContext, dataset, split)
    return rdd.flatMap(lambda tmpl: [canonicalize(im, version) for im in tmpl])
    
def splits(sparkContext, split=None, rdd=rdd):
    split = range(1,11) if split is None else split
    return [ [rdd(sparkContext, dataset=d, split=s) for d in ['train', 'gallery', 'probe']] for s in tolist(split) ]

def stream(dataset='train', split=None, version=3):
    s = janus_cs2.stream(dataset, split)
    s_map = itertools.imap(lambda im: canonicalize(im, version) if im.mediatype()=='image' else im.mapFrames(lambda fr: canonicalize(fr, version)), s)
    return s_map

