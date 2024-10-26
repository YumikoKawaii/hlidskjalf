package authenticator

import (
	"context"
	pb "elysium.com/pb/authenticator"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	authorizationHeader = "authorization"
)

func (s *Handler) Verify(ctx context.Context, request *pb.VerifyRequest) (*pb.VerifyResponse, error) {

	claim, err := s.jwtResolver.VerifyToken(request.Token)
	if err != nil {
		return nil, err
	}

	userId := claim.GetInfo().Id

	// verify route
	authorized, err := s.storage.IsAccessible(ctx, userId, request.Route)
	if err != nil {
		logrus.Errorf("error verify: %s, id: %s, route: %s", err.Error(), userId, request.Route)
	}

	if !authorized {
		return &pb.VerifyResponse{
			Code: http.StatusUnauthorized,
			Id:   userId,
		}, nil
	}

	return &pb.VerifyResponse{
		Code: http.StatusOK,
		Id:   claim.GetInfo().Id,
	}, nil
}
