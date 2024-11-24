// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.3
// source: proto/event-service.proto

package espb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	Load(ctx context.Context, opts ...grpc.CallOption) (EventService_LoadClient, error)
	Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	Events(ctx context.Context, opts ...grpc.CallOption) (EventService_EventsClient, error)
	Sport(ctx context.Context, in *SportRequest, opts ...grpc.CallOption) (*SportResponse, error)
	Sports(ctx context.Context, opts ...grpc.CallOption) (EventService_SportsClient, error)
	GetUpcomingEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (EventService_GetUpcomingEventsClient, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Load(ctx context.Context, opts ...grpc.CallOption) (EventService_LoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[0], "/events.EventService/Load", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceLoadClient{stream}
	return x, nil
}

type EventService_LoadClient interface {
	Send(*LoadRequest) error
	CloseAndRecv() (*LoadResponse, error)
	grpc.ClientStream
}

type eventServiceLoadClient struct {
	grpc.ClientStream
}

func (x *eventServiceLoadClient) Send(m *LoadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServiceLoadClient) CloseAndRecv() (*LoadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LoadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) Event(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/Event", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Events(ctx context.Context, opts ...grpc.CallOption) (EventService_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[1], "/events.EventService/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceEventsClient{stream}
	return x, nil
}

type EventService_EventsClient interface {
	Send(*EventRequest) error
	Recv() (*EventResponse, error)
	grpc.ClientStream
}

type eventServiceEventsClient struct {
	grpc.ClientStream
}

func (x *eventServiceEventsClient) Send(m *EventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServiceEventsClient) Recv() (*EventResponse, error) {
	m := new(EventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) Sport(ctx context.Context, in *SportRequest, opts ...grpc.CallOption) (*SportResponse, error) {
	out := new(SportResponse)
	err := c.cc.Invoke(ctx, "/events.EventService/Sport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Sports(ctx context.Context, opts ...grpc.CallOption) (EventService_SportsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[2], "/events.EventService/Sports", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceSportsClient{stream}
	return x, nil
}

type EventService_SportsClient interface {
	Send(*SportRequest) error
	Recv() (*SportResponse, error)
	grpc.ClientStream
}

type eventServiceSportsClient struct {
	grpc.ClientStream
}

func (x *eventServiceSportsClient) Send(m *SportRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServiceSportsClient) Recv() (*SportResponse, error) {
	m := new(SportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) GetUpcomingEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (EventService_GetUpcomingEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[3], "/events.EventService/GetUpcomingEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServiceGetUpcomingEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventService_GetUpcomingEventsClient interface {
	Recv() (*UpcomingEventResponse, error)
	grpc.ClientStream
}

type eventServiceGetUpcomingEventsClient struct {
	grpc.ClientStream
}

func (x *eventServiceGetUpcomingEventsClient) Recv() (*UpcomingEventResponse, error) {
	m := new(UpcomingEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	Load(EventService_LoadServer) error
	Event(context.Context, *EventRequest) (*EventResponse, error)
	Events(EventService_EventsServer) error
	Sport(context.Context, *SportRequest) (*SportResponse, error)
	Sports(EventService_SportsServer) error
	GetUpcomingEvents(*emptypb.Empty, EventService_GetUpcomingEventsServer) error
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) Load(EventService_LoadServer) error {
	return status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedEventServiceServer) Event(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Event not implemented")
}
func (UnimplementedEventServiceServer) Events(EventService_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (UnimplementedEventServiceServer) Sport(context.Context, *SportRequest) (*SportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sport not implemented")
}
func (UnimplementedEventServiceServer) Sports(EventService_SportsServer) error {
	return status.Errorf(codes.Unimplemented, "method Sports not implemented")
}
func (UnimplementedEventServiceServer) GetUpcomingEvents(*emptypb.Empty, EventService_GetUpcomingEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUpcomingEvents not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_Load_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).Load(&eventServiceLoadServer{stream})
}

type EventService_LoadServer interface {
	SendAndClose(*LoadResponse) error
	Recv() (*LoadRequest, error)
	grpc.ServerStream
}

type eventServiceLoadServer struct {
	grpc.ServerStream
}

func (x *eventServiceLoadServer) SendAndClose(m *LoadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServiceLoadServer) Recv() (*LoadRequest, error) {
	m := new(LoadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventService_Event_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Event(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/Event",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Event(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).Events(&eventServiceEventsServer{stream})
}

type EventService_EventsServer interface {
	Send(*EventResponse) error
	Recv() (*EventRequest, error)
	grpc.ServerStream
}

type eventServiceEventsServer struct {
	grpc.ServerStream
}

func (x *eventServiceEventsServer) Send(m *EventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServiceEventsServer) Recv() (*EventRequest, error) {
	m := new(EventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventService_Sport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Sport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.EventService/Sport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Sport(ctx, req.(*SportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Sports_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).Sports(&eventServiceSportsServer{stream})
}

type EventService_SportsServer interface {
	Send(*SportResponse) error
	Recv() (*SportRequest, error)
	grpc.ServerStream
}

type eventServiceSportsServer struct {
	grpc.ServerStream
}

func (x *eventServiceSportsServer) Send(m *SportResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServiceSportsServer) Recv() (*SportRequest, error) {
	m := new(SportRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventService_GetUpcomingEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventServiceServer).GetUpcomingEvents(m, &eventServiceGetUpcomingEventsServer{stream})
}

type EventService_GetUpcomingEventsServer interface {
	Send(*UpcomingEventResponse) error
	grpc.ServerStream
}

type eventServiceGetUpcomingEventsServer struct {
	grpc.ServerStream
}

func (x *eventServiceGetUpcomingEventsServer) Send(m *UpcomingEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "events.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Event",
			Handler:    _EventService_Event_Handler,
		},
		{
			MethodName: "Sport",
			Handler:    _EventService_Sport_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Load",
			Handler:       _EventService_Load_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Events",
			Handler:       _EventService_Events_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Sports",
			Handler:       _EventService_Sports_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetUpcomingEvents",
			Handler:       _EventService_GetUpcomingEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/event-service.proto",
}
