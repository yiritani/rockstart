package server

import (
	pb "rockstart/src/_generated/proto/v1"
	sqlc "rockstart/src/_generated/sqlc"

	"google.golang.org/grpc"
)

type Server struct {
	GRPCServer *grpc.Server
	Queries    *sqlc.Queries
}

func NewServer(queries *sqlc.Queries) *Server {
	server := &Server{
		GRPCServer: grpc.NewServer(),
		Queries:    queries,
	}
	registerServices(server)
	return server
}

func registerServices(s *Server) {
	pb.RegisterPingServiceServer(s.GRPCServer, &PingServer{Queries: s.Queries})
	pb.RegisterInstrumentServiceServer(s.GRPCServer, &InstrumentServer{Queries: s.Queries})
}
