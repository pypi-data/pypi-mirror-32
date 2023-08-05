# -*- coding: utf-8 -*-
from __future__ import unicode_literals


# http://momentjs.com/docs/#/parsing/string-format/
# http://docs.python.org/2/library/datetime.html#strftime-strptime-behavior
FORMAT_MAP = (
    ('DDD', r'%j'),
    ('DD', r'%d'),
    ('MMMM', r'%B'),
    ('MMM', r'%b'),
    ('MM', r'%m'),
    ('YYYY', r'%Y'),
    ('YY', r'%y'),
    ('HH', r'%H'),
    ('hh', r'%I'),
    ('mm', r'%M'),
    ('ss', r'%S'),
    ('a', r'%p'),
    ('ZZ', r'%z'),
)


def datetime_format_py2js(format):
    """
    Converts Python's :func:`~time.strftime` format to a JavaScript notation,
    eg. ``"%d"`` to ``"DD"`` or ``"%Y"`` to ``"YYYY"``.
    """
    for js, py in FORMAT_MAP:
        format = format.replace(py, js)
    return format


def datetime_format_js2py(format):
    """
    Converts a JavaScript datetime notation to Python's :func:`~time.strftime`
    format, eg. ``"DD"`` to ``"%d"`` or ``"YYYY"`` to ``"%Y"`` .
    """
    for js, py in FORMAT_MAP:
        format = format.replace(js, py)
    return format
