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


class ContainerApiResult(Model):
    """Container API result.

    All required parameters must be populated in order to send to Azure.

    :param status: Required. HTTP status code returned by the target container
     API
    :type status: int
    :param content_type: HTTP content type
    :type content_type: str
    :param content_encoding: HTTP content encoding
    :type content_encoding: str
    :param body: container API result body
    :type body: str
    """

    _validation = {
        'status': {'required': True},
    }

    _attribute_map = {
        'status': {'key': 'Status', 'type': 'int'},
        'content_type': {'key': 'Content-Type', 'type': 'str'},
        'content_encoding': {'key': 'Content-Encoding', 'type': 'str'},
        'body': {'key': 'Body', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ContainerApiResult, self).__init__(**kwargs)
        self.status = kwargs.get('status', None)
        self.content_type = kwargs.get('content_type', None)
        self.content_encoding = kwargs.get('content_encoding', None)
        self.body = kwargs.get('body', None)
