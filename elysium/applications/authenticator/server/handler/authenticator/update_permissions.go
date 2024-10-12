package authenticator

import (
	"context"
	"elysium.com/applications/authenticator/pkg/repository"
	pb "elysium.com/pb/authenticator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const updatePermissionRoute = "/api/v1/permissions"

func (s *Handler) UpdatePermissions(ctx context.Context, request *pb.UpdatePermissionsRequest) (*pb.UpdatePermissionsResponse, error) {

	// 1. check api key
	if !s.apiKeyResolver.Verify(updatePermissionRoute, request.ApiKey) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid api key")
	}

	// 2. Update permissions
	if err := s.authService.UpsertPermissions(ctx, s.transformProtoToPermissions(request.Id, request.Permissions)); err != nil {
		return nil, err
	}

	return &pb.UpdatePermissionsResponse{
		Code:    http.StatusOK,
		Message: "Success",
	}, nil
}

func (s *Handler) transformProtoToPermissions(id string, permissions []string) []repository.Permission {
	ps := make([]repository.Permission, 0)
	for _, p := range permissions {
		ps = append(ps, repository.Permission{
			AccountId: id,
			Route:     p,
		})
	}
	return ps
}
