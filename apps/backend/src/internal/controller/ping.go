package controller

import (
	"context"
	proto_v1 "rockstart/src/_generated/proto/v1"
	"rockstart/src/internal/service"

	"connectrpc.com/connect"
)

type PingServer struct {
}

func (p *PingServer) Ping(ctx context.Context, req *connect.Request[proto_v1.PingRequest]) (*connect.Response[proto_v1.PingResponse], error) {
	resp, err := service.PingService(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
