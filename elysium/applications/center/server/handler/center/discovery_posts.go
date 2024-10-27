package center

import (
	"context"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/center"
)

func (c *Handler) DiscoveryPosts(ctx context.Context, request *pb.DiscoveryPostsRequest) (*pb.DiscoveryPostsResponse, error) {

	resp, err := c.postClient.Discovery(ctx, posts.DiscoveryRequest{
		Author:    request.Author,
		SortOrder: utils.ToSortOrder(int32(request.SortOrder)),
		Page:      request.Page,
		PageSize:  request.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DiscoveryPostsResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Data: &pb.DiscoveryPostsResponse_Data{
			PostIds:  resp.PostIds,
			Page:     request.Page,
			PageSize: request.PageSize,
		},
	}, nil
}
