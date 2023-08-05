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

from .chaos_event import ChaosEvent


class ValidationFailedChaosEvent(ChaosEvent):
    """Chaos event corresponding to a failure during validation.

    All required parameters must be populated in order to send to Azure.

    :param time_stamp_utc: Required. The UTC timestamp when this Chaos event
     was generated.
    :type time_stamp_utc: datetime
    :param kind: Required. Constant filled by server.
    :type kind: str
    :param reason: Describes why the ValidationFailedChaosEvent was generated.
     This may happen because more than MaxPercentUnhealthyNodes are unhealthy
     for more than MaxClusterStabilizationTimeout. This reason will be in the
     Reason property of the ValidationFailedChaosEvent as a string.
    :type reason: str
    """

    _validation = {
        'time_stamp_utc': {'required': True},
        'kind': {'required': True},
    }

    _attribute_map = {
        'time_stamp_utc': {'key': 'TimeStampUtc', 'type': 'iso-8601'},
        'kind': {'key': 'Kind', 'type': 'str'},
        'reason': {'key': 'Reason', 'type': 'str'},
    }

    def __init__(self, **kwargs):
        super(ValidationFailedChaosEvent, self).__init__(**kwargs)
        self.reason = kwargs.get('reason', None)
        self.kind = 'ValidationFailed'
