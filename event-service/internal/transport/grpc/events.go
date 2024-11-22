package grpc

import "mzhn/event-service/pb/espb"

// Events implements espb.EventServiceServer.
func (s *Server) Events(espb.EventService_EventsServer) error {
	panic("unimplemented")
}
