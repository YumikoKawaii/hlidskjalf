package user

import (
	"elysium.com/applications/user/pkg/userservice"
	pb "elysium.com/pb/user"
)

type Handler struct {
	*pb.UnimplementedUserServiceServer
	userService userservice.Service
}

func NewHandler(service userservice.Service) *Handler {
	return &Handler{
		userService: service,
	}
}
