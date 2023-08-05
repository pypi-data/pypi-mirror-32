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


class RestoreProgressInfo(Model):
    """Describes the progress of a restore operation on a partition.

    :param restore_state: Represents the current state of the partition
     restore operation. Possible values include: 'Invalid', 'Accepted',
     'RestoreInProgress', 'Success', 'Failure', 'Timeout'
    :type restore_state: str or ~azure.servicefabric.models.RestoreState
    :param time_stamp_utc: Timestamp when operation succeeded or failed.
    :type time_stamp_utc: datetime
    :param restored_epoch: Describes the epoch at which the partition is
     restored.
    :type restored_epoch: ~azure.servicefabric.models.BackupEpoch
    :param restored_lsn: Restored LSN.
    :type restored_lsn: str
    :param failure_error: Denotes the failure encountered in performing
     restore operation.
    :type failure_error: ~azure.servicefabric.models.FabricErrorError
    """

    _attribute_map = {
        'restore_state': {'key': 'RestoreState', 'type': 'str'},
        'time_stamp_utc': {'key': 'TimeStampUtc', 'type': 'iso-8601'},
        'restored_epoch': {'key': 'RestoredEpoch', 'type': 'BackupEpoch'},
        'restored_lsn': {'key': 'RestoredLsn', 'type': 'str'},
        'failure_error': {'key': 'FailureError', 'type': 'FabricErrorError'},
    }

    def __init__(self, *, restore_state=None, time_stamp_utc=None, restored_epoch=None, restored_lsn: str=None, failure_error=None, **kwargs) -> None:
        super(RestoreProgressInfo, self).__init__(**kwargs)
        self.restore_state = restore_state
        self.time_stamp_utc = time_stamp_utc
        self.restored_epoch = restored_epoch
        self.restored_lsn = restored_lsn
        self.failure_error = failure_error
