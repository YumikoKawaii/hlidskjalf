package user

import (
	"context"
	"elysium.com/applications/users/pkg/repository"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/user"
	"net/http"
)

func (s *Handler) UpsertUser(ctx context.Context, request *pb.UpsertUserRequest) (*pb.UpsertUserResponse, error) {

	if err := s.userService.UpsertUser(ctx, &repository.User{
		Id:           utils.ProtoToStringPointer(request.Id),
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
