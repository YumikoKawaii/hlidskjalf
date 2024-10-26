package interceptor

import (
	"context"
	"elysium.com/applications/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	auth Authenticator
}

func NewInterceptor(host string, usingGRPC bool) *Interceptor {
	var auth Authenticator
	if usingGRPC {
		auth = NewGRPCAuthenticator(host)
	} else {
		auth = NewHttpAuthenticator(host)
	}

	return &Interceptor{
		auth: auth,
	}
}

func (i *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		tokenData, err := utils.ExtractValueFromIncomingContext(ctx, authorizationHeader)
		if err != nil {
			logrus.Errorf("error token: %s, method: %s", err.Error(), info.FullMethod)
			return nil, xerrors.Errorf("error token: %s, method: %s", err.Error(), info.FullMethod)
		}

		token, ok := tokenData.(string)
		if !ok {
			return nil, xerrors.Errorf("invalid token: %v, method: %s", tokenData, info.FullMethod)
		}

		id, authorized, err := i.auth.Verify(ctx, token, info.FullMethod)
		if err != nil {
			logrus.Errorf("error authenticator: %s, userId: %s, method: %s", err.Error(), id, info.FullMethod)
		}
		if err != nil || !authorized {
			return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
		}

		return handler(context.WithValue(ctx, utils.UserIdKey, id), req)
	}
}
