package proxy

import "context"

type contextKey struct{}

func WithOriginalDst(ctx context.Context, dst string) context.Context {
	return context.WithValue(ctx, contextKey{}, dst)
}

func OriginalDstFromContext(ctx context.Context) (string, bool) {
	dst, ok := ctx.Value(contextKey{}).(string)
	return dst, ok
}
