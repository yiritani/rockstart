package server

import (
	"context"
	pb "rockstart/src/_generated/proto/v1"
	"rockstart/src/_generated/sqlc"
	service "rockstart/src/internal/service"
)

type InstrumentServer struct {
	pb.UnimplementedInstrumentServiceServer
	Queries *sqlc.Queries
}

func (s *InstrumentServer) GetInstruments(ctx context.Context, req *pb.GetInstrumentsRequest) (*pb.GetInstrumentsResponse, error) {
	res, err := service.GetInstrumentService(s.Queries, ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
