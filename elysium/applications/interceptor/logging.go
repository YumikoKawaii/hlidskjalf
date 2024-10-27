package interceptor

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Logger struct {
}

func (l *Logger) Unary(entity string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		logrus.Infof("[%s] user agent: %s, authority: %s, method: %s", entity, md["user-agent"][0], md[":authority"][0], info.FullMethod)
		return handler(ctx, req)
	}
}
