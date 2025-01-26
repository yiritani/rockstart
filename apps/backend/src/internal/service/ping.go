package service

import (
	"context"
	"fmt"
	proto_v1 "rockstart/src/_generated/proto/v1"
)

func PingService(ctx context.Context, req *proto_v1.PingRequest) (*proto_v1.PingResponse, error) {
	msg := req.Message
	return &proto_v1.PingResponse{
		Message: fmt.Sprintf("Pong from gRPC backend: %s", msg),
	}, nil
}
