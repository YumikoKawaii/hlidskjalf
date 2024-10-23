package posts

import (
	"context"
	"elysium.com/applications/posts/pkg/repository"
	pb "elysium.com/pb/posts"
	"net/http"
)

func (s *Handler) GetPosts(ctx context.Context, request *pb.GetPostsRequest) (*pb.GetPostResponse, error) {
	params := &repository.GetPostsParams{
		Ids: request.Ids,
	}

	posts, err := s.postService.GetPosts(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.GetPostResponse_Data{
			Posts: s.transformPostToProto(posts),
		},
	}, nil
}

func (s *Handler) transformPostToProto(posts []repository.Post) []*pb.Post {
	protos := make([]*pb.Post, 0)
	for _, post := range posts {
		protos = append(protos, &pb.Post{
			Id:        *post.Id,
			Author:    post.Author,
			Content:   post.Content,
			CreatedAt: uint32(post.CreatedAt.UTC().Unix()),
			UpdatedAt: uint32(post.UpdatedAt.UTC().Unix()),
		})
	}
	return protos
}
