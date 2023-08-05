from __future__ import absolute_import
from __future__ import division
from __future__ import print_function


__all__ = [
    'ChecksumMismatchError',
]


class ChecksumMismatchError(Exception):
    pass


class ChecksumNotSupportedError(Exception):
    pass
