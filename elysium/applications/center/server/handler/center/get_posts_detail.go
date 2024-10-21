package center

import (
	"elysium.com/applications/center/pkg/adapter/interactions"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/center/pkg/adapter/users"
	pb "elysium.com/pb/center"
	golang_set "github.com/deckarep/golang-set/v2"
	"golang.org/x/net/context"
)

func (c *Handler) GetPostsDetail(ctx context.Context, request *pb.GetPostsDetailRequest) (*pb.GetPostsDetailResponse, error) {

	// context will contains id
	// get posts
	postsResp, err := c.postClient.GetPosts(ctx, posts.GetPostsRequest{
		Ids:      request.PostIds,
		Page:     request.Page,
		PageSize: request.PageSize,
	})
	if err != nil {
		return nil, err
	}

	authorIds := make([]string, 0)
	postIds := make([]uint32, 0)
	for _, post := range postsResp.Data.Posts {
		postIds = append(postIds, post.Id)
		authorIds = append(authorIds, post.Author)
	}

	// fetch interactions
	interactionList := make([]interactions.Interaction, 0)
	page := uint32(1)
	pageSize := uint32(100)
	for {
		// fetch until done
		interactionsResp, err := c.interactionClient.GetInteractions(ctx, interactions.GetInteractionRequest{
			PostIds:  postIds,
			Page:     page,
			PageSize: pageSize,
		})
		if err != nil {
			return nil, err
		}

		interactionList = append(interactionList, interactionsResp.Data.Interactions...)

		if interactionsResp.Data.PageSize < pageSize {
			break
		}
		page += 1
	}

	// fetch users
	for _, interaction := range interactionList {
		authorIds = append(authorIds, interaction.Author)
	}

	authorIds = golang_set.NewSet[string](authorIds...).ToSlice()
	page = 1
	for {
		authorResp, err := c.userClient.GetUsers(ctx, users.GetUsersRequest{
			Ids:      authorIds,
			Page:     0,
			PageSize: 0,
		})
	}

	return nil, nil
}
