import pytest
from pikepdf import (parse_content_stream, Page, Pdf, Stream, Operator, Object,
        Dictionary)
from pikepdf._qpdf import StreamParser
import os
from subprocess import run, PIPE
import shutil


class PrintParser(StreamParser):
    def __init__(self):
        super().__init__()

    def handle_object(self, obj):
        print(repr(obj))

    def handle_eof(self):
        print("--EOF--")


class ExceptionParser(StreamParser):
    def __init__(self):
        super().__init__()

    def handle_object(self, obj):
        raise ValueError("I take exception to this")

    def handle_eof(self):
        print("--EOF--")


def test_open_pdf(resources):
    pdf = Pdf.open(resources / 'graph.pdf')
    stream = pdf.pages[0]['/Contents']
    Object._parse_stream(stream, PrintParser())


def test_parser_exception(resources):
    pdf = Pdf.open(resources / 'graph.pdf')
    stream = pdf.pages[0]['/Contents']
    with pytest.raises(ValueError):
        Object._parse_stream(stream, ExceptionParser())


@pytest.mark.skipif(
    shutil.which('pdftotext') is None,
    reason="poppler not installed")
def test_text_filter(resources, outdir):
    input_pdf = resources / 'veraPDF test suite 6-2-10-t02-pass-a.pdf'

    # Ensure the test PDF has detect we can find
    proc = run(['pdftotext', str(input_pdf), '-'],
        check=True, stdout=PIPE, encoding='utf-8')
    assert proc.stdout.strip() != '', "Need input test file that contains text"

    pdf = Pdf.open(input_pdf)
    stream = pdf.pages[0].Contents

    keep = []
    for operands, command in parse_content_stream(stream):
        if command == Operator('Tj'):
            print("skipping Tj")
            continue
        keep.append((operands, command))

    new_stream = Stream(pdf, keep)
    print(new_stream.read_bytes())
    pdf.pages[0]['/Contents'] = new_stream
    pdf.pages[0]['/Rotate'] = 90

    pdf.save(outdir / 'notext.pdf', True)

    proc = run(['pdftotext', str(outdir / 'notext.pdf'), '-'],
        check=True, stdout=PIPE, encoding='utf-8')

    assert proc.stdout.strip() == '', "Expected text to be removed"


def test_invalid_stream_object():
    with pytest.raises(TypeError):
        parse_content_stream(Dictionary({"/Hi": 3}))


@pytest.mark.parametrize("test_file,expected", [
    ("fourpages.pdf", True),
    ("graph.pdf", False),
    ("veraPDF test suite 6-2-10-t02-pass-a.pdf", True),
    ("veraPDF test suite 6-2-3-3-t01-fail-c.pdf", False),
    ('sandwich.pdf', True)
])
def test_has_text(resources, test_file, expected):
    pdf = Pdf.open(resources / test_file)
    for p in pdf.pages:
        page = Page(p)
        assert page.has_text() == expected
