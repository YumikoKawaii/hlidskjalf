package entry

import (
	"context"
	"elysium.com/applications/valgrind/pkg/publish"
	pb "elysium.com/pb/valgrind"
	"elysium.com/shared/topics"
	"encoding/base64"
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

	// decode payload
	payload, err := base64.StdEncoding.DecodeString(request.Payload)
	if err != nil {
		return &pb.EntryResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid payload",
		}, nil
	}

	if err := s.publisher.Publish(ctx, topics.EntryTopic, payload); err != nil {
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
