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


class RestorePartitionDescription(Model):
    """Specifies the parameters needed to trigger a restore of a specific
    partition.

    All required parameters must be populated in order to send to Azure.

    :param backup_id: Required. Unique backup ID.
    :type backup_id: str
    :param backup_location: Required. Location of the backup relative to the
     backup storage specified/ configured.
    :type backup_location: str
    :param backup_storage: Location of the backup from where the partition
     will be restored.
    :type backup_storage: ~azure.servicefabric.models.BackupStorageDescription
    """

    _validation = {
        'backup_id': {'required': True},
        'backup_location': {'required': True},
    }

    _attribute_map = {
        'backup_id': {'key': 'BackupId', 'type': 'str'},
        'backup_location': {'key': 'BackupLocation', 'type': 'str'},
        'backup_storage': {'key': 'BackupStorage', 'type': 'BackupStorageDescription'},
    }

    def __init__(self, *, backup_id: str, backup_location: str, backup_storage=None, **kwargs) -> None:
        super(RestorePartitionDescription, self).__init__(**kwargs)
        self.backup_id = backup_id
        self.backup_location = backup_location
        self.backup_storage = backup_storage
