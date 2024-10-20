package users

import (
	"context"
	pb "elysium.com/pb/user"
)

type grpcClient struct {
	rpc pb.UserServiceClient
}

func (c *grpcClient) UpsertUser(ctx context.Context, request UpsertUserRequest) (UpsertUserResponse, error) {
	resp, err := c.rpc.UpsertUser(ctx, &pb.UpsertUserRequest{
		Name:         request.Name,
		Alias:        request.Alias,
		Avatar:       request.Avatar,
		Introduction: request.Introduction,
		Workplace:    request.Workplace,
		Hometown:     request.Hometown,
	})

	if err != nil {
		return UpsertUserResponse{}, err
	}

	return UpsertUserResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Id:      resp.Id,
	}, err
}

func (c *grpcClient) GetUsers(ctx context.Context, request GetUsersRequest) (GetUsersResponse, error) {
	resp, err := c.rpc.GetUsersInfo(ctx, &pb.GetUsersInfoRequest{
		Ids:      request.Ids,
		Page:     request.Page,
		PageSize: request.PageSize,
	})

	if err != nil {
		return GetUsersResponse{}, err
	}

	return GetUsersResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Data: struct {
			UsersInfo []UserInfo `json:"usersInfo,omitempty"`
			Page      int32      `json:"page,omitempty"`
			PageSize  int32      `json:"pageSize,omitempty"`
		}{
			UsersInfo: nil,
			Page:      resp.Data.Page,
			PageSize:  resp.Data.PageSize,
		},
	}, err
}

func (c *grpcClient) convertProtosToUsers(protos []*pb.UserInfo) []UserInfo {
	infos := make([]UserInfo, 0)
	for _, proto := range protos {
		infos = append(infos, UserInfo{
			Id:           proto.Id,
			Name:         proto.Name,
			Alias:        proto.Alias,
			Avatar:       proto.Avatar,
			Introduction: proto.Introduction,
			Workplace:    proto.Workplace,
			Hometown:     proto.Hometown,
		})
	}
	return infos
}
