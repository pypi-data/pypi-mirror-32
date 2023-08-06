# -*- coding: utf-8 -*-
from __future__ import absolute_import

from .serialization import EncryptedSerializer
from kombu.serialization import registry

__author__ = 'Bryan Shelton'
__email__ = 'bryan@rover.com'
__version__ = '0.1.0'


def setup_encrypted_serializer(key=None, serializer='pickle'):
    encrypted_serializer = EncryptedSerializer(key=key, serializer=serializer)
    name = "encrypted_{0}".format(serializer)
    registry.register(
        "encrypted_{0}".format(serializer),
        encrypted_serializer.serialize,
        encrypted_serializer.deserialize,
        content_type="application/x-encrypted-{0}".format(serializer),
        content_encoding='utf-8',
    )
    return name
