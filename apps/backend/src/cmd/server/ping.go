package server

import (
	"context"

	pb "rockstart/src/_generated/proto/v1"
	sqlc "rockstart/src/_generated/sqlc"
)

type PingServer struct {
	pb.UnimplementedPingServiceServer
	Queries *sqlc.Queries
}

func (s *PingServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "Pong"}, nil
}
