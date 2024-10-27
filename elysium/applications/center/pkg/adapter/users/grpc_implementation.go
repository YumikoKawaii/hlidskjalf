package users

import (
	"context"
	pb "elysium.com/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	rpc pb.UserServiceClient
}

func NewRpcClient(config Config) Client {

	conn, err := grpc.NewClient(config.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	rpc := pb.NewUserServiceClient(conn)
	return &grpcClient{
		rpc: rpc,
	}
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
			Page      uint32     `json:"page,omitempty"`
			PageSize  uint32     `json:"pageSize,omitempty"`
		}{
			UsersInfo: c.convertProtosToUsers(resp.Data.UsersInfo),
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
