package center

import (
	"elysium.com/applications/center/pkg/adapter/interactions"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/center/pkg/adapter/users"
	pb "elysium.com/pb/center"
	golang_set "github.com/deckarep/golang-set/v2"
	"golang.org/x/net/context"
	"net/http"
)

const (
	defaultGetInteractionPage     = 1
	defaultGetInteractionPageSize = 100
)

func (c *Handler) GetPostsDetail(ctx context.Context, request *pb.GetPostsDetailRequest) (*pb.GetPostsDetailResponse, error) {

	postsResp, err := c.postClient.GetPosts(ctx, posts.GetPostsRequest{
		Ids: []uint32{request.PostId},
	})
	if err != nil {
		return nil, err
	}
	post := postsResp.Data.Posts[0]

	authorIds := make([]string, 0)
	postIds := make([]uint32, 0)
	for _, post := range postsResp.Data.Posts {
		postIds = append(postIds, post.Id)
		authorIds = append(authorIds, post.Author)
	}

	// fetch interactions
	interactionsResp, err := c.interactionClient.GetInteractions(ctx, interactions.GetInteractionRequest{
		PostId:   request.PostId,
		Page:     defaultGetInteractionPage,
		PageSize: defaultGetInteractionPageSize,
	})
	if err != nil {
		return nil, err
	}

	// fetch users
	for _, interaction := range interactionsResp.Data.Interactions {
		authorIds = append(authorIds, interaction.Author)
	}

	authorIds = golang_set.NewSet[string](authorIds...).ToSlice()
	authorResp, err := c.userClient.GetUsers(ctx, users.GetUsersRequest{
		Ids:      authorIds,
		Page:     1,
		PageSize: uint32(len(authorIds)),
	})
	if err != nil {
		return nil, err
	}

	// create author map
	authorMap := make(map[string]*pb.Author)
	for _, author := range authorResp.Data.UsersInfo {
		authorMap[author.Id] = &pb.Author{
			Id:     author.Id,
			Name:   author.Name,
			Alias:  author.Alias,
			Avatar: author.Avatar,
		}
	}

	protoInteractions := make([]*pb.Interaction, 0)
	for _, interaction := range interactionsResp.Data.Interactions {
		protoInteractions = append(protoInteractions, &pb.Interaction{
			Id:        interaction.Id,
			Author:    authorMap[interaction.Author],
			Type:      interaction.Type,
			Content:   interaction.Content,
			CreatedAt: interaction.CreatedAt,
			UpdatedAt: interaction.UpdatedAt,
		})
	}

	return &pb.GetPostsDetailResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.GetPostsDetailResponse_Data{
			PostsDetail: &pb.PostDetail{
				Id:           post.Id,
				Author:       authorMap[post.Author],
				Content:      post.Content,
				CreatedAt:    post.CreatedAt,
				UpdatedAt:    post.UpdatedAt,
				Interactions: protoInteractions,
			},
		},
	}, nil
}
