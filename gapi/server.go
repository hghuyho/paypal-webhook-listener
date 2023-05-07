package gapi

import (
	"github.com/gin-gonic/gin"
	"paypal-webhook-listener/pb"
)

type Server struct {
	pb.UnimplementedPublicAppServer
	pb.UnimplementedWebhookServiceServer
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}
	return server, nil
}
