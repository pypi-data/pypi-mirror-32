VERSION = (0, 4, 0)
__version__ = ".".join(map(str, VERSION))

try:
    from .fields import *  # noqa
except ImportError:  # pragma: no cover
    pass
