package jwt

import (
	"context"
	"elysium.com/applications/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	resolver Resolver
}

func NewInterceptor(resolver Resolver) Interceptor {
	return Interceptor{
		resolver: resolver,
	}
}

func (i *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		claim, _, err := i.authenticate(ctx)
		if err != nil {
			return nil, err
		}

		return handler(context.WithValue(ctx, utils.UserIdKey, claim.Id), req)
	}
}

func (i *Interceptor) authenticate(ctx context.Context) (*Claim, string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, "", status.Error(codes.Unauthenticated, "Unauthorized")
	}

	values := md["authorization"]
	if md["authorization"] == nil {
		values = md["Authorization"]
	}

	if len(values) == 0 {
		return nil, "", status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	token := values[0]
	claim, err := i.resolver.VerifyToken(token)
	if err != nil {
		return nil, "", status.Errorf(codes.Unauthenticated, "Unauthorized")
	}
	return claim, token, nil
}
