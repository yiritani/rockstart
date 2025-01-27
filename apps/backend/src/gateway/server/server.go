package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "rockstart/src/_generated/proto/v1"
)

func NewServer(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption, grpcServerEndpoint string) {
	registerServices(ctx, mux, opts, grpcServerEndpoint)
}

func registerServices(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption, grpcServerEndpoint string) {
	err := pb.RegisterPingServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, nil)
	if err != nil {
		return
	}

	err = pb.RegisterInstrumentServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return
	}
}
