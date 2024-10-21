package utils

import (
	"context"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/metadata"
)

const (
	UserIdKey = "id"
	XAPIKey   = "x-api-key"
)

func ExtractValueFromContext(ctx context.Context, key string) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, xerrors.Errorf("not found")
	}

	values := md[key]
	if values == nil {
		return nil, xerrors.Errorf("not found")
	}

	if len(values) == 0 {
		return nil, xerrors.Errorf("not found")
	}

	return values[0], nil
}
