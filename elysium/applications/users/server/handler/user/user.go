package user

import (
	"elysium.com/applications/users/pkg/user_service"
	pb "elysium.com/pb/user"
)

type Handler struct {
	*pb.UnimplementedUserServiceServer
	userService user_service.Service
}

func NewHandler(service user_service.Service) *Handler {
	return &Handler{
		userService: service,
	}
}
