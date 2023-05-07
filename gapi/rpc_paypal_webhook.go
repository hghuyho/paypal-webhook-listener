package gapi

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"paypal-webhook-listener/pb"
)

func (server *Server) PayPalWebhook(ctx context.Context, req *pb.WebhookRequest) (*emptypb.Empty, error) {
	//md, _ := metadata.FromIncomingContext(ctx)
	// TODO: event header validation
	return &emptypb.Empty{}, nil
}
