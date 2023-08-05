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

from .property_batch_operation import PropertyBatchOperation


class CheckExistsPropertyBatchOperation(PropertyBatchOperation):
    """Represents a PropertyBatchOperation that compares the Boolean existence of
    a property with the Exists argument.
    The PropertyBatchOperation operation fails if the property's existence is
    not equal to the Exists argument.
    The CheckExistsPropertyBatchOperation is generally used as a precondition
    for the write operations in the batch.
    Note that if one PropertyBatchOperation in a PropertyBatch fails,
    the entire batch fails and cannot be committed in a transactional manner.

    All required parameters must be populated in order to send to Azure.

    :param property_name: Required. The name of the Service Fabric property.
    :type property_name: str
    :param kind: Required. Constant filled by server.
    :type kind: str
    :param exists: Required. Whether or not the property should exist for the
     operation to pass.
    :type exists: bool
    """

    _validation = {
        'property_name': {'required': True},
        'kind': {'required': True},
        'exists': {'required': True},
    }

    _attribute_map = {
        'property_name': {'key': 'PropertyName', 'type': 'str'},
        'kind': {'key': 'Kind', 'type': 'str'},
        'exists': {'key': 'Exists', 'type': 'bool'},
    }

    def __init__(self, **kwargs):
        super(CheckExistsPropertyBatchOperation, self).__init__(**kwargs)
        self.exists = kwargs.get('exists', None)
        self.kind = 'CheckExists'
