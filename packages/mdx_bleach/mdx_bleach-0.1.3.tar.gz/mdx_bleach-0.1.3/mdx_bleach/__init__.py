# -*- coding: utf-8 -*-
from mdx_bleach.extension import BleachExtension


VERSION = '0.1.3'


def makeExtension(**kwargs):
    return BleachExtension(**kwargs)
