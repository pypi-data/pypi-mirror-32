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


class ContainerApiRequestBody(Model):
    """parameters for making container API call.

    All required parameters must be populated in order to send to Azure.

    :param http_verb: HTTP verb of container REST API, defaults to "GET"
    :type http_verb: str
    :param uri_path: Required. URI path of container REST API
    :type uri_path: str
    :param content_type: Content type of container REST API request, defaults
     to "application/json"
    :type content_type: str
    :param body: HTTP request body of container REST API
    :type body: str
    """

    _validation = {
        'uri_path': {'required': True},
    }

    _attribute_map = {
        'http_verb': {'key': 'HttpVerb', 'type': 'str'},
        'uri_path': {'key': 'UriPath', 'type': 'str'},
        'content_type': {'key': 'Content-Type', 'type': 'str'},
        'body': {'key': 'Body', 'type': 'str'},
    }

    def __init__(self, *, uri_path: str, http_verb: str=None, content_type: str=None, body: str=None, **kwargs) -> None:
        super(ContainerApiRequestBody, self).__init__(**kwargs)
        self.http_verb = http_verb
        self.uri_path = uri_path
        self.content_type = content_type
        self.body = body
