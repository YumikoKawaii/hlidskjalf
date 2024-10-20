package interactions

import (
	"context"
	"elysium.com/applications/interactions/pkg/repository"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/interactions"
	"net/http"
)

func (s *Handler) UpsertInteraction(ctx context.Context, request *pb.UpsertInteractionRequest) (*pb.UpsertInteractionResponse, error) {
	interaction := s.transformProtoToInteraction(request)
	err := s.interactionService.UpsertInteraction(ctx, interaction)
	if err != nil {
		return nil, err
	}

	return &pb.UpsertInteractionResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}, nil
}

func (s *Handler) transformProtoToInteraction(request *pb.UpsertInteractionRequest) *repository.Interaction {
	return &repository.Interaction{
		Id:      utils.ProtoToUInt32Pointer(request.Id),
		PostId:  request.PostId,
		Author:  request.Author,
		Type:    request.Type,
		Content: request.Content,
	}
}
