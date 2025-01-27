package service

import (
	"context"
	"database/sql"
	pb "rockstart/src/_generated/proto/v1"
	"rockstart/src/_generated/sqlc"
)

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
func GetInstrumentService(queries *sqlc.Queries, ctx context.Context, req *pb.GetInstrumentsRequest) (*pb.GetInstrumentsResponse, error) {
	instruments, err := queries.GetInstrument(ctx, toNullString(*req.SearchQuery))
	if err != nil {
		return nil, err
	}

	var instrumentResponses []*pb.Instrument
	for _, instrument := range instruments {
		instrumentResponses = append(instrumentResponses, &pb.Instrument{
			Id:   instrument.ID,
			Name: instrument.Name,
		})
	}

	return &pb.GetInstrumentsResponse{
		Instruments: instrumentResponses,
	}, nil
}
