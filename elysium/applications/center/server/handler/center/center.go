package center

import (
	"elysium.com/applications/center/pkg/adapter/interactions"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/center/pkg/adapter/users"
	pb "elysium.com/pb/center"
)

type Handler struct {
	pb.CenterServiceServer
	postClient        posts.Client
	interactionClient interactions.Client
	userClient        users.Client
}

func NewHandler(postClient posts.Client, interactionClient interactions.Client, userClient users.Client) *Handler {
	return &Handler{
		postClient:        postClient,
		interactionClient: interactionClient,
		userClient:        userClient,
	}
}
