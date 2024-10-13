package interactions

import (
	"elysium.com/applications/interactions/pkg/interaction_service"
	pb "elysium.com/pb/interactions"
)

type Handler struct {
	*pb.UnimplementedInteractionServiceServer
	interactionService interaction_service.Service
}

func NewHandler(service interaction_service.Service) *Handler {
	return &Handler{
		interactionService: service,
	}
}
