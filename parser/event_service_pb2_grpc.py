# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import event_service_pb2 as event__service__pb2

GRPC_GENERATED_VERSION = "1.68.0"
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower

    _version_not_supported = first_version_is_lower(
        GRPC_VERSION, GRPC_GENERATED_VERSION
    )
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f"The grpc package installed is at version {GRPC_VERSION},"
        + f" but the generated code in event_service_pb2_grpc.py depends on"
        + f" grpcio>={GRPC_GENERATED_VERSION}."
        + f" Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}"
        + f" or downgrade your generated code using grpcio-tools<={GRPC_VERSION}."
    )


class EventServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Load = channel.stream_unary(
            "/events.EventService/Load",
            request_serializer=event__service__pb2.LoadRequest.SerializeToString,
            response_deserializer=event__service__pb2.LoadResponse.FromString,
            _registered_method=True,
        )
        self.Events = channel.stream_stream(
            "/events.EventService/Events",
            request_serializer=event__service__pb2.EventRequest.SerializeToString,
            response_deserializer=event__service__pb2.EventResponse.FromString,
            _registered_method=True,
        )


class EventServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Load(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")

    def Events(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_EventServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "Load": grpc.stream_unary_rpc_method_handler(
            servicer.Load,
            request_deserializer=event__service__pb2.LoadRequest.FromString,
            response_serializer=event__service__pb2.LoadResponse.SerializeToString,
        ),
        "Events": grpc.stream_stream_rpc_method_handler(
            servicer.Events,
            request_deserializer=event__service__pb2.EventRequest.FromString,
            response_serializer=event__service__pb2.EventResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "events.EventService", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers("events.EventService", rpc_method_handlers)


# This class is part of an EXPERIMENTAL API.
class EventService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Load(
        request_iterator,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.stream_unary(
            request_iterator,
            target,
            "/events.EventService/Load",
            event__service__pb2.LoadRequest.SerializeToString,
            event__service__pb2.LoadResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True,
        )

    @staticmethod
    def Events(
        request_iterator,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.stream_stream(
            request_iterator,
            target,
            "/events.EventService/Events",
            event__service__pb2.EventRequest.SerializeToString,
            event__service__pb2.EventResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True,
        )