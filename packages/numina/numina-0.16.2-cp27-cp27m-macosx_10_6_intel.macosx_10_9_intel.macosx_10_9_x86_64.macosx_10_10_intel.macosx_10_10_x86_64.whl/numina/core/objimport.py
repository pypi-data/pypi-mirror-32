#
# Copyright 2011-2014 Universidad Complutense de Madrid
#
# This file is part of Numina
#
# SPDX-License-Identifier: GPL-3.0+
# License-Filename: LICENSE.txt
#

"""Import objects by name"""

import importlib
import inspect


def import_object(path):
    """Import an object given its fully qualified name."""
    spl = path.split('.')
    if len(spl) == 1:
        return importlib.import_module(path)
    # avoid last part for the moment
    cls = spl[-1]
    mods = '.'.join(spl[:-1])

    mm = importlib.import_module(mods)
    # try to get the last part as an attribute
    try:
        obj = getattr(mm, cls)
        return obj
    except AttributeError:
        pass

    # Try to import the last part
    rr = importlib.import_module(path)
    return rr


def fully_qualified_name(obj, sep='.'):
    if inspect.isclass(obj):
        return obj.__module__ + sep + obj.__name__
    else:
        return obj.__module__ + sep + obj.__class__.__name__



