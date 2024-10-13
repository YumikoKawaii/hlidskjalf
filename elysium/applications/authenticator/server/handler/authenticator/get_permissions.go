package authenticator

import (
	"context"
	pb "elysium.com/pb/authenticator"
	"net/http"
)

func (s *Handler) GetPermissions(ctx context.Context, request *pb.GetPermissionsRequest) (*pb.GetPermissionsResponse, error) {
	permissions, err := s.storage.GetPermissionsById(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetPermissionsResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data: &pb.GetPermissionsResponse_Data{
			Id:          request.UserId,
			Permissions: permissions,
		},
	}, nil
}
