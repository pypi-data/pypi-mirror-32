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

from .fabric_event_py3 import FabricEvent


class NodeEvent(FabricEvent):
    """Represents the base for all Node Events.

    You probably want to use the sub-classes and not this class directly. Known
    sub-classes are: NodeAbortedEvent, NodeAbortingEvent, NodeAddedEvent,
    NodeCloseEvent, NodeClosingEvent, NodeDeactivateCompleteEvent,
    NodeDeactivateStartEvent, NodeDownEvent, NodeHealthReportCreatedEvent,
    NodeHealthReportExpiredEvent, NodeOpenedSuccessEvent, NodeOpenFailedEvent,
    NodeOpeningEvent, NodeRemovedEvent, NodeUpEvent,
    ChaosRestartNodeFaultCompletedEvent, ChaosRestartNodeFaultScheduledEvent

    All required parameters must be populated in order to send to Azure.

    :param event_instance_id: Required. The identifier for the FabricEvent
     instance.
    :type event_instance_id: str
    :param time_stamp: Required. The time event was logged.
    :type time_stamp: datetime
    :param has_correlated_events: Shows there is existing related events
     available.
    :type has_correlated_events: bool
    :param kind: Required. Constant filled by server.
    :type kind: str
    :param node_name: Required. The name of a Service Fabric node.
    :type node_name: str
    """

    _validation = {
        'event_instance_id': {'required': True},
        'time_stamp': {'required': True},
        'kind': {'required': True},
        'node_name': {'required': True},
    }

    _attribute_map = {
        'event_instance_id': {'key': 'EventInstanceId', 'type': 'str'},
        'time_stamp': {'key': 'TimeStamp', 'type': 'iso-8601'},
        'has_correlated_events': {'key': 'HasCorrelatedEvents', 'type': 'bool'},
        'kind': {'key': 'Kind', 'type': 'str'},
        'node_name': {'key': 'NodeName', 'type': 'str'},
    }

    _subtype_map = {
        'kind': {'NodeAborted': 'NodeAbortedEvent', 'NodeAborting': 'NodeAbortingEvent', 'NodeAdded': 'NodeAddedEvent', 'NodeClose': 'NodeCloseEvent', 'NodeClosing': 'NodeClosingEvent', 'NodeDeactivateComplete': 'NodeDeactivateCompleteEvent', 'NodeDeactivateStart': 'NodeDeactivateStartEvent', 'NodeDown': 'NodeDownEvent', 'NodeHealthReportCreated': 'NodeHealthReportCreatedEvent', 'NodeHealthReportExpired': 'NodeHealthReportExpiredEvent', 'NodeOpenedSuccess': 'NodeOpenedSuccessEvent', 'NodeOpenFailed': 'NodeOpenFailedEvent', 'NodeOpening': 'NodeOpeningEvent', 'NodeRemoved': 'NodeRemovedEvent', 'NodeUp': 'NodeUpEvent', 'ChaosRestartNodeFaultCompleted': 'ChaosRestartNodeFaultCompletedEvent', 'ChaosRestartNodeFaultScheduled': 'ChaosRestartNodeFaultScheduledEvent'}
    }

    def __init__(self, *, event_instance_id: str, time_stamp, node_name: str, has_correlated_events: bool=None, **kwargs) -> None:
        super(NodeEvent, self).__init__(event_instance_id=event_instance_id, time_stamp=time_stamp, has_correlated_events=has_correlated_events, **kwargs)
        self.node_name = node_name
        self.kind = 'NodeEvent'
