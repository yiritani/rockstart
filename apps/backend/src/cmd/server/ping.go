package server

import (
	"context"

	pb "rockstart/src/_generated/proto/v1"
)

type PingServer struct {
	pb.UnimplementedPingServiceServer
}

func (s *PingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "Pong"}, nil
}
