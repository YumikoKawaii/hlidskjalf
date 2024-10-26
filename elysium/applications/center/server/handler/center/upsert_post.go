package center

import (
	"context"
	"elysium.com/applications/center/pkg/adapter/posts"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/center"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (c *Handler) UpsertPost(ctx context.Context, request *pb.UpsertPostRequest) (*pb.UpsertPostResponse, error) {

	userIdData, err := utils.ExtractValueFromContext(ctx, utils.UserIdKey)
	if err != nil {
		logrus.Error(err.Error())
		return &pb.UpsertPostResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}, nil
	}

	userId, ok := userIdData.(string)
	if !ok {
		logrus.Errorf("error invalid userId convert")
		return &pb.UpsertPostResponse{
			Code:    http.StatusUnauthorized,
			Message: "invalid userId convert",
		}, nil
	}

	resp, err := c.postClient.UpsertPost(ctx, posts.UpsertPostRequest{
		Id:      utils.ProtoToUInt32Pointer(request.Id),
		Author:  userId,
		Content: request.Content,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpsertPostResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Id:      resp.Id,
	}, nil
}
