package interactions

import (
	"context"
	"elysium.com/applications/interactions/pkg/repository"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/interactions"
	"net/http"
)

func (s *Handler) GetInteractions(ctx context.Context, request *pb.GetInteractionsRequest) (*pb.GetInteractionResponse, error) {
	params := &repository.GetInteractionsParams{
		PostId:   request.PostId,
		Order:    utils.ToSortOrder(int32(request.SortOrder)),
		Page:     request.Page,
		PageSize: request.PageSize,
	}

	interactions, err := s.interactionService.GetInteractions(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.GetInteractionResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.GetInteractionResponse_Data{
			Interactions: s.transformInteractionsToProtos(interactions),
			Page:         params.Page,
			PageSize:     params.PageSize,
		},
	}, nil
}

func (s *Handler) transformInteractionsToProtos(interactions []repository.Interaction) []*pb.Interaction {
	protos := make([]*pb.Interaction, 0)
	for _, i := range interactions {
		protos = append(protos, &pb.Interaction{
			Id:        *i.Id,
			PostId:    i.PostId,
			Author:    i.Author,
			Type:      i.Type,
			Content:   i.Content,
			CreatedAt: uint32(i.CreatedAt.Unix()),
			UpdatedAt: uint32(i.UpdatedAt.Unix()),
		})
	}
	return protos
}
