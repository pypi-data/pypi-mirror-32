# -*- coding: utf-8 -*-
# @author: Federico Cerchiari <federicocerchiari@gmail.com>
import sys
from .tools import render_template
from .elements import Tag, VoidTag, Css, Content
from .tempyrepr import TempyREPR
from .places import TempyPlace
from .t import T
from .tempy import Escaped

__version__ = '1.2.0'
VERSION = tuple(map(int, __version__.split('.')))

if sys.version_info < (3, 3):
    raise RuntimeError('You need Python 3.3+ for this module.')
