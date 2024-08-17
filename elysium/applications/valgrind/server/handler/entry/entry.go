package entry

import (
	"context"
	"elysium.com/applications/valgrind/pkg/publish"
	pb "elysium.com/pb/valgrind"
	"elysium.com/shared/topics"
	"net/http"
)

type Service struct {
	pb.UnimplementedEntryServiceServer
	publisher publish.Publisher
}

func NewService(publisher publish.Publisher) *Service {
	return &Service{
		publisher: publisher,
	}
}

func (s *Service) Entry(ctx context.Context, request *pb.EntryRequest) (*pb.EntryResponse, error) {

	if len(request.Payload) == 0 {
		return &pb.EntryResponse{
			Code:    http.StatusBadRequest,
			Message: "empty payload",
		}, nil
	}

	if err := s.publisher.Publish(ctx, topics.EntryTopic, request.Payload); err != nil {
		return &pb.EntryResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	return &pb.EntryResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
