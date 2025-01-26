package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "rockstart/src/_generated/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"rockstart/src/cmd/server"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	// ctx := context.Background()
	// conn := config.NewDBConfig().ConnectDB(ctx)
	// defer conn.Close(ctx)

	// queries := sqlc.New(conn)
	// NewServer(queries)


	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server.PingServer{})
	log.Printf("server listening at %v", lis.Addr())

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}