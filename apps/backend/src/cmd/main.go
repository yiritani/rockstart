package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "rockstart/src/_generated/proto/v1"

	sqlc "rockstart/src/_generated/sqlc"
	"rockstart/src/cmd/server"
	config "rockstart/src/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	ctx := context.Background()
	conn := config.NewDBConfig().ConnectSQLiteDB(ctx)
	defer conn.Close()

	queries := sqlc.New(conn)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server.NewServer(queries)
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server.PingServer{})
	pb.RegisterInstrumentServiceServer(s, &server.InstrumentServer{Queries: queries})
	log.Printf("server listening at %v", lis.Addr())

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
