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

func NewService() *Service {
	return &Service{}
}

func (s *Service) Entry(ctx context.Context, request *pb.EntryRequest) (*pb.EntryResponse, error) {

	if err := s.publisher.Publish(ctx, topics.EntryTopic, []byte(request.Payload)); err != nil {
		return nil, err
	}

	return &pb.EntryResponse{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
