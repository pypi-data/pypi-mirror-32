# -*- coding: utf-8 -*-
#
from __future__ import print_function

# https://github.com/pybind/pybind11/issues/1004
# pylint: disable=wildcard-import
from _accupy import *

from .__about__ import (
    __author__,
    __email__,
    __copyright__,
    __credits__,
    __license__,
    __version__,
    __maintainer__,
    __status__,
    )

from .dot import *
from .ill_cond import *
from .sums import *

try:
    import pipdate
except ImportError:
    pass
else:
    if pipdate.needs_checking(__name__):
        print(pipdate.check(__name__, __version__), end='')
