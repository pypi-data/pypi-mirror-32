# -*- coding: utf-8 -*-
__author__  = 'Laurent P. René de Cotret'
__email__   = 'laurent.renedecotret@mail.mcgill.ca'
__license__ = 'MIT'
__version__ = '5.0.1'

# Versioning should be compliant with PyPI guide 
# https://packaging.python.org/tutorials/distributing-packages/#choosing-a-versioning-scheme

from .raw import AbstractRawDataset, check_raw_bounds
from .mcgill import McGillRawDataset, LegacyMcGillRawDataset
from .dataset import DiffractionDataset, PowderDiffractionDataset
from .meta import ExperimentalParameter

from . import plugins
