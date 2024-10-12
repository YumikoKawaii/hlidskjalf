package authenticator

import (
	"context"
	"elysium.com/applications/authenticator/pkg/jwt"
	pb "elysium.com/pb/authenticator"
	"net/http"
)

func (s *Handler) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {

	account, err := s.authService.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	// 2. Generate token
	token, err := s.jwtResolver.GenerateToken(jwt.Info{
		Id:   account.Id,
		User: account.HashedEmail,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Token:   token,
	}, nil
}
