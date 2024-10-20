package posts

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/posts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	rpc pb.PostServiceClient
}

func NewRpcClient(config Config) Client {

	conn, err := grpc.NewClient(config.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	rpc := pb.NewPostServiceClient(conn)
	return &grpcClient{
		rpc: rpc,
	}
}

func (c *grpcClient) UpsertPost(ctx context.Context, request UpsertPostRequest) (UpsertPostResponse, error) {

	resp, err := c.rpc.UpsertPost(ctx, &pb.UpsertPostRequest{
		Id:      utils.UInt32PointerToProto(request.Id),
		Author:  request.Author,
		Content: request.Content,
		Medias:  request.Medias,
	})
	if err != nil {
		return UpsertPostResponse{}, err
	}

	return UpsertPostResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Id:      resp.Id,
	}, nil
}

func (c *grpcClient) GetPosts(ctx context.Context, request GetPostsRequest) (GetPostsResponse, error) {
	resp, err := c.rpc.GetPosts(ctx, &pb.GetPostsRequest{
		Ids:      request.Ids,
		Page:     request.Page,
		PageSize: request.PageSize,
	})
	if err != nil {
		return GetPostsResponse{}, err
	}

	return GetPostsResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Data: struct {
			Posts    []Post `json:"posts,omitempty"`
			Page     int32  `json:"page,omitempty"`
			PageSize int32  `json:"pageSize,omitempty"`
		}{
			Posts:    c.convertProtosToPosts(resp.Data.Posts),
			Page:     resp.Data.Page,
			PageSize: resp.Data.PageSize,
		},
	}, nil
}

func (c *grpcClient) convertProtosToPosts(protos []*pb.Post) []Post {
	posts := make([]Post, 0)
	for _, proto := range protos {
		posts = append(posts, Post{
			Id:        proto.Id,
			Author:    proto.Author,
			Content:   proto.Content,
			CreatedAt: proto.CreatedAt,
			UpdatedAt: proto.UpdatedAt,
		})
	}
	return posts
}
