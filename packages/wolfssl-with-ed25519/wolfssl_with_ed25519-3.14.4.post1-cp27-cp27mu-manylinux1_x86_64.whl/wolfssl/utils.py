# -*- coding: utf-8 -*-
#
# utils.py
#
# Copyright (C) 2006-2017 wolfSSL Inc.
#
# This file is part of wolfSSL. (formerly known as CyaSSL)
#
# wolfSSL is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# wolfSSL is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA

# pylint: disable=missing-docstring, unused-import, undefined-variable

import sys
from binascii import hexlify as b2h, unhexlify as h2b  # noqa: F401

_PY3 = sys.version_info[0] == 3
_TEXT_TYPE = str if _PY3 else unicode  # noqa: F821
_BINARY_TYPE = bytes if _PY3 else str


def t2b(string):
    """
    Converts text to bynary.
    """
    if isinstance(string, _BINARY_TYPE):
        return string
    return _TEXT_TYPE(string).encode("utf-8")
