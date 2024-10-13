package user

import (
	"context"
	"elysium.com/applications/users/pkg/repository"
	pb "elysium.com/pb/user"
	"net/http"
)

func (s *Handler) GetUsersInfo(ctx context.Context, request *pb.GetUsersInfoRequest) (*pb.GetUsersInfoResponse, error) {

	users, err := s.userService.GetUsersInfoByIds(ctx, request.Ids)
	if err != nil {
		return nil, err
	}

	return &pb.GetUsersInfoResponse{
		Code:      http.StatusOK,
		Message:   "Success",
		UsersInfo: s.transformUsersToProto(users),
	}, nil
}

func (s *Handler) transformUsersToProto(users []repository.User) []*pb.UserInfo {

	protos := make([]*pb.UserInfo, 0)
	for _, user := range users {
		protos = append(protos, &pb.UserInfo{
			Id:           user.Id,
			Name:         user.Name,
			Alias:        user.Alias,
			Avatar:       user.Avatar,
			Introduction: user.Introduction,
			Workplace:    user.Workplace,
			Hometown:     user.Hometown,
		})
	}
	return protos
}
