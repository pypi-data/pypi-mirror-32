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

from .backup_storage_description_py3 import BackupStorageDescription


class AzureBlobBackupStorageDescription(BackupStorageDescription):
    """Describes the parameters for Azure blob store used for storing and
    enumerating backups.

    All required parameters must be populated in order to send to Azure.

    :param friendly_name: Friendly name for this backup storage.
    :type friendly_name: str
    :param storage_kind: Required. Constant filled by server.
    :type storage_kind: str
    :param connection_string: Required. The connection string to connect to
     the Azure blob store.
    :type connection_string: str
    :param container_name: Required. The name of the container in the blob
     store to store and enumerate backups from.
    :type container_name: str
    """

    _validation = {
        'storage_kind': {'required': True},
        'connection_string': {'required': True},
        'container_name': {'required': True},
    }

    _attribute_map = {
        'friendly_name': {'key': 'FriendlyName', 'type': 'str'},
        'storage_kind': {'key': 'StorageKind', 'type': 'str'},
        'connection_string': {'key': 'ConnectionString', 'type': 'str'},
        'container_name': {'key': 'ContainerName', 'type': 'str'},
    }

    def __init__(self, *, connection_string: str, container_name: str, friendly_name: str=None, **kwargs) -> None:
        super(AzureBlobBackupStorageDescription, self).__init__(friendly_name=friendly_name, **kwargs)
        self.connection_string = connection_string
        self.container_name = container_name
        self.storage_kind = 'AzureBlobStore'
