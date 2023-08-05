# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.
#
# Copyright (C) 2017, James R. Barlow (https://github.com/jbarlow83/)


from collections import namedtuple
from pkg_resources import get_distribution, DistributionNotFound

import os

try:
    __version__ = get_distribution(__name__).version
except DistributionNotFound:
    __version__ = "Not installed"
    pass
try:
    from . import _qpdf
except ImportError:
    raise ImportError("pikepdf's extension library failed to import")

from ._qpdf import (Object, ObjectType, PdfError, Pdf, PasswordError,
        ObjectStreamMode, StreamDataMode)

from ._objects import (Boolean, Integer, Real, Name, String, Array, Dictionary,
        Stream, Operator, Null)

from ._pdfimage import PdfImage, UnsupportedImageTypeError

__libqpdf_version__ = _qpdf.qpdf_version()



class _OperandGrouper(_qpdf.StreamParser):
    """Parse a PDF content stream into a sequence of instructions.

    Helper class for parsing PDF content streams into instructions. Semantics
    are a little weird since it is subclassed from C++.

    """

    PdfInstruction = namedtuple('PdfInstruction', ('operands', 'operator'))

    def __init__(self):
        super().__init__()
        self.instructions = []
        self._tokens = []

    def handle_object(self, obj):
        if obj.type_code == ObjectType.operator:
            instruction = self.PdfInstruction(
                operands=self._tokens, operator=obj)
            self.instructions.append(instruction)
            self._tokens = []
        else:
            self._tokens.append(obj)

    def handle_eof(self):
        if self._tokens:
            raise EOFError("Unexpected end of stream")


def parse_content_stream(stream):
    """Parse a PDF content stream into a sequence of instructions.

    A PDF content stream is list of instructions that describe where to render
    the text and graphics in a PDF. This is the starting point for analyzing
    PDFs.

    Each instruction contains at least one operator and zero or more operands.

    The *stream* object may be either a :class:`pikepdf.Stream` or an array
    of streams.

    >>> pdf = pikepdf.Pdf.open(input_pdf)
    >>> stream = pdf.pages[0].Contents
    >>> for operands, command in parse_content_stream(stream):
    >>>     print(command)

    """

    if not isinstance(stream, Object):
        raise TypeError("stream must a PDF object")

    if stream.type_code != ObjectType.stream:
        raise TypeError("parse_content_stream called on non-stream Object")

    grouper = _OperandGrouper()
    try:
        Object._parse_stream(stream, grouper)
    except PdfError as e:
        if 'parseContentStream called on non-stream' in str(e):  # qpdf 6.x
            raise TypeError("parse_content_stream called on non-stream Object")
        elif 'ignoring non-stream while parsing' in str(e):  # qpdf 7.0
            raise TypeError("parse_content_stream called on non-stream Object")
        raise e from e

    return grouper.instructions


class Page:
    def __init__(self, obj):
        self.obj = obj

    def __getattr__(self, item):
        return getattr(self.obj, item)

    def __setattr__(self, item, value):
        if item == 'obj':
            object.__setattr__(self, item, value)
        elif hasattr(self.obj, item):
            setattr(self.obj, item, value)
        else:
            raise AttributeError(item)

    def __repr__(self):
        return repr(self.obj).replace(
            'pikepdf.Dictionary', 'pikepdf.Page', 1)

    @property
    def mediabox(self):
        return self.obj.MediaBox

    def has_text(self):
        """Check if this page print text

        Search the content stream for any of the four text showing operators.
        We ignore text positioning operators because some editors might
        generate maintain these even if text is deleted etc.

        This cannot detect raster text (text in a bitmap), text rendered as
        curves. It also cannot determine if the text is visible to the user.

        :return: True if there is text
        """
        text_showing_operators = {Operator(op) for op in
                                  ('Tj', 'TJ', '"', "'")}
        for _, operator in parse_content_stream(self.obj.Contents):
            if operator in text_showing_operators:
                return True
        return False


def open(*args, **kwargs):
    "Alias for :func:`pikepdf.Pdf.open`."
    return Pdf.open(*args, **kwargs)


class PdfMatrix:
    """Support class for PDF content stream matrices

    PDF content stream matrices are 3x3 matrices summarized by a shorthand
    (a, b, c, d, e, f) which correspond to the first two column vectors. The
    final column vector is always (0, 0, 1) since this is using homogenous
    coordinates.

    PDF uses row vectors.  That is vr @ A' gives the effect of transforming
    a row vector vr=(x, y, 1) by the matrix A'.  It's more common to use
    A @ vc where vc is a column vector = (x, y, 1)'.

    Addition is not implemented. If needed it would be necessary to normalize
    with division by self.values[2][2].

    """

    def __init__(self, other):
        if isinstance(other, PdfMatrix):
            self.values = other.values
        elif len(other) == 6:
            a, b, c, d, e, f = map(float, other)
            self.values = ((a, b, 0),
                           (c, d, 0),
                           (e, f, 1))
        elif len(other) == 3 and len(other[0]) == 3:
            self.values = (tuple(other[0]),
                           tuple(other[1]),
                           tuple(other[2]))
        else:
            raise ValueError('arguments')

    @staticmethod
    def identity():
        return PdfMatrix((1, 0, 0, 1, 0, 0))

    def __matmul__(self, other):
        a = self.values
        b = other.values
        return PdfMatrix(
                [[sum([float(i) * float(j)
                       for i, j in zip(row, col)]
                     ) for col in zip(*b)]
                  for row in a]
        )

    @property
    def a(self):
        return self.values[0][0]

    @a.setter
    def a(self, value):
        self.values[0][0] = value

    @property
    def b(self):
        return self.values[0][1]

    @b.setter
    def b(self, value):
        self.values[0][1] = value

    @property
    def c(self):
        return self.values[1][0]

    @c.setter
    def c(self, value):
        self.values[1][0] = value

    @property
    def d(self):
        return self.values[1][1]

    @d.setter
    def d(self, value):
        self.values[1][1] = value

    @property
    def e(self):
        return self.values[2][0]

    @e.setter
    def e(self, value):
        self.values[2][0] = value

    @property
    def f(self):
        return self.values[2][1]

    @f.setter
    def f(self, value):
        self.values[2][1] = value

    def encode(self):
        return '{:.6f} {:.6f} {:.6f} {:.6f} {:.6f} {:.6f}'.format(
            self.a, self.b, self.c, self.d, self.e, self.f
        ).encode()
