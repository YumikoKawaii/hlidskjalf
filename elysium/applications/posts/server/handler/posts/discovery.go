package posts

import (
	"context"
	"elysium.com/applications/posts/pkg/repository"
	pb "elysium.com/pb/posts"
	"net/http"
)

func (s *Handler) Discovery(ctx context.Context, request *pb.DiscoveryRequest) (*pb.DiscoveryResponse, error) {

	params := &repository.GetPostsParams{
		Author:   request.Author,
		Page:     request.Page,
		PageSize: request.PageSize,
	}

	posts, err := s.postService.GetPosts(ctx, params)
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, 0)
	for _, post := range posts {
		ids = append(ids, *post.Id)
	}

	return &pb.DiscoveryResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.DiscoveryResponse_Data{
			PostIds:  ids,
			Page:     request.Page,
			PageSize: request.PageSize,
		},
	}, nil
}
