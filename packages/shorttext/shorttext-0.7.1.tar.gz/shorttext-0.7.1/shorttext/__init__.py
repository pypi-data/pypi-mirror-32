import os
import sys

thisdir, _ = os.path.split(__file__)
sys.path.append(thisdir)

from . import utils
from . import data
from . import classifiers
from . import generators
from . import stack
from .smartload import smartload_compact_model
from . import metrics
from . import spell

__version__ = '0.5.8'