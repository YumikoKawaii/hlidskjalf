package authenticator

import (
	"elysium.com/applications/authenticator/pkg/api_key"
	"elysium.com/applications/authenticator/pkg/auth"
	"elysium.com/applications/authenticator/pkg/discovery"
	"elysium.com/applications/authenticator/pkg/jwt"
	pb "elysium.com/pb/authenticator"
)

type Handler struct {
	*pb.UnimplementedAuthenticatorServer
	authService    auth.Service
	jwtResolver    jwt.Resolver
	apiKeyResolver api_key.Resolver
	storage        discovery.Storage
}

func NewService(service auth.Service, jwtResolver jwt.Resolver, apiKeyResolver api_key.Resolver, storage discovery.Storage) *Handler {
	return &Handler{
		authService:    service,
		jwtResolver:    jwtResolver,
		apiKeyResolver: apiKeyResolver,
		storage:        storage,
	}
}
