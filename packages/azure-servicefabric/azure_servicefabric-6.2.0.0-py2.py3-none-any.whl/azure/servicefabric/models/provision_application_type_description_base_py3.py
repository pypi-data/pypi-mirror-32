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


class ProvisionApplicationTypeDescriptionBase(Model):
    """Represents the type of registration or provision requested, and if the
    operation needs to be asynchronous or not. Supported types of provision
    operations are from either image store or external store.

    You probably want to use the sub-classes and not this class directly. Known
    sub-classes are: ProvisionApplicationTypeDescription,
    ExternalStoreProvisionApplicationTypeDescription

    All required parameters must be populated in order to send to Azure.

    :param async_property: Required. Indicates whether or not provisioning
     should occur asynchronously. When set to true, the provision operation
     returns when the request is accepted by the system, and the provision
     operation continues without any timeout limit. The default value is false.
     For large application packages, we recommend setting the value to true.
    :type async_property: bool
    :param kind: Required. Constant filled by server.
    :type kind: str
    """

    _validation = {
        'async_property': {'required': True},
        'kind': {'required': True},
    }

    _attribute_map = {
        'async_property': {'key': 'Async', 'type': 'bool'},
        'kind': {'key': 'Kind', 'type': 'str'},
    }

    _subtype_map = {
        'kind': {'ImageStorePath': 'ProvisionApplicationTypeDescription', 'ExternalStore': 'ExternalStoreProvisionApplicationTypeDescription'}
    }

    def __init__(self, *, async_property: bool, **kwargs) -> None:
        super(ProvisionApplicationTypeDescriptionBase, self).__init__(**kwargs)
        self.async_property = async_property
        self.kind = None
