package gapi

import (
	"context"
	"paypal-webhook-listener/pb"
)

func (server *Server) HealthCheck(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	rsp := &pb.PingResponse{
		Message: "Pong",
	}
	return rsp, nil
}
