package entry

import (
	"context"
	pb "elysium.com/pb/valgrind"
	"net/http"
)

type Service struct {
	pb.UnimplementedEntryServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Entry(ctx context.Context, request *pb.EntryRequest) (*pb.EntryResponse, error) {
	return &pb.EntryResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
