package authenticator

import (
	"context"
	"elysium.com/applications/authenticator/pkg/jwt"
	pb "elysium.com/pb/authenticator"
	"net/http"
)

func (s *Handler) Signup(ctx context.Context, request *pb.SignupRequest) (*pb.SignupResponse, error) {

	// 1. Create account
	account, err := s.authService.Signup(ctx, request.Email, request.Password)
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

	return &pb.SignupResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Token:   token,
	}, nil
}
