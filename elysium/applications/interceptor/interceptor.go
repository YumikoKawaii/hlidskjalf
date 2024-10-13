package interceptor

import (
	"context"
	"elysium.com/applications/utils"
	"github.com/sirupsen/logrus"
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
		id, authorized, err := i.auth.Verify(ctx, info.FullMethod)
		if err != nil {
			logrus.Errorf("error authenticator: %s", err.Error())
		}
		if err != nil || !authorized {
			return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
		}

		return handler(context.WithValue(ctx, utils.UserIdKey, id), req)
	}
}
