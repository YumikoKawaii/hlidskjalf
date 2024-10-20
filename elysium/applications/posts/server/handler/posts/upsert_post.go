package posts

import (
	"context"
	"elysium.com/applications/posts/pkg/repository"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/posts"
	"net/http"
)

func (s *Handler) UpsertPost(ctx context.Context, request *pb.UpsertPostRequest) (*pb.UpsertPostResponse, error) {

	post := s.transformProtoToPost(request)
	err := s.postService.UpsertPost(ctx, post)
	if err != nil {
		return nil, err
	}

	return &pb.UpsertPostResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Id:      *post.Id,
	}, nil
}

func (s *Handler) transformProtoToPost(request *pb.UpsertPostRequest) *repository.Post {
	return &repository.Post{
		Id:      utils.ProtoToUInt32Pointer(request.Id),
		Author:  request.Author,
		Content: request.Content,
	}
}
