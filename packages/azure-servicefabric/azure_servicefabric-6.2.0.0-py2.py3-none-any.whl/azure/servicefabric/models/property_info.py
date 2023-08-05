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


class PropertyInfo(Model):
    """Information about a Service Fabric property.

    All required parameters must be populated in order to send to Azure.

    :param name: Required. The name of the Service Fabric property.
    :type name: str
    :param value: Describes a Service Fabric property value.
    :type value: ~azure.servicefabric.models.PropertyValue
    :param metadata: Required. The metadata associated with a property,
     including the property's name.
    :type metadata: ~azure.servicefabric.models.PropertyMetadata
    """

    _validation = {
        'name': {'required': True},
        'metadata': {'required': True},
    }

    _attribute_map = {
        'name': {'key': 'Name', 'type': 'str'},
        'value': {'key': 'Value', 'type': 'PropertyValue'},
        'metadata': {'key': 'Metadata', 'type': 'PropertyMetadata'},
    }

    def __init__(self, **kwargs):
        super(PropertyInfo, self).__init__(**kwargs)
        self.name = kwargs.get('name', None)
        self.value = kwargs.get('value', None)
        self.metadata = kwargs.get('metadata', None)
