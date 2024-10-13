package authenticator

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/authenticator"
	"golang.org/x/xerrors"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

func (s *Handler) Verify(ctx context.Context, request *pb.VerifyRequest) (*pb.VerifyResponse, error) {

	// parse token
	data, err := utils.ExtractValueFromContext(ctx, authorizationHeader)
	if err != nil {
		return nil, err
	}

	token, ok := data.(string)
	if !ok {
		return nil, xerrors.Errorf("unavailable")
	}

	claim, err := s.jwtResolver.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	userId := claim.GetInfo().Id

	// verify route
	permissions, err := s.storage.GetPermissionsById(ctx, userId)
	if err != nil {
		return nil, err
	}

	authorized := false
	for _, p := range permissions {
		if p == request.Route {
			authorized = true
			break
		}
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
