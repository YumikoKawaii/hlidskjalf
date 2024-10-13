package user

import (
	"context"
	"elysium.com/applications/user/pkg/repository"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/user"
	"net/http"
)

func (s *Handler) UpsertUser(ctx context.Context, request *pb.UpsertUserRequest) (*pb.UpsertUserResponse, error) {

	idData, err := utils.ExtractValueFromContext(ctx, utils.UserIdKey)
	if err != nil {
		return &pb.UpsertUserResponse{
			Code: http.StatusUnauthorized,
		}, nil
	}

	id, ok := idData.(string)
	if !ok {
		return &pb.UpsertUserResponse{
			Code: http.StatusUnauthorized,
		}, nil
	}

	if err := s.userService.UpsertUser(ctx, &repository.User{
		Id:           id,
		Name:         request.Name,
		Alias:        request.Alias,
		Avatar:       request.Avatar,
		Introduction: request.Introduction,
		Workplace:    request.Workplace,
		Hometown:     request.Hometown,
	}); err != nil {
		return nil, err
	}

	return &pb.UpsertUserResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}, nil
}
