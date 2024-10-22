package posts

import (
	"elysium.com/applications/posts/pkg/post_service"
	pb "elysium.com/pb/posts"
)

type Handler struct {
	*pb.UnimplementedPostServiceServer
	postService post_service.Service
}

func NewHandler(service post_service.Service) *Handler {
	return &Handler{
		postService: service,
	}
}
