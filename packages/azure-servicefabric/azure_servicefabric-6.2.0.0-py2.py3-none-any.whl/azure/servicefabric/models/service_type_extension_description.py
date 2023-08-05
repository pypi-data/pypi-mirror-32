# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------

from msrest.serialization import Model


class ServiceTypeExtensionDescription(Model):
    """Describes extension of a service type defined in the service manifest.

    :param key: The name of the extension.
    :type key: str
    :param value: The extension value.
    :type value: str
    """

    _attribute_map = {
        'key': {'key': 'Key', 'type': 'str'},
        'value': {'key': 'Value', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ServiceTypeExtensionDescription, self).__init__(**kwargs)
        self.key = kwargs.get('key', None)
        self.value = kwargs.get('value', None)
