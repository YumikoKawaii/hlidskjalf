package handler

import (
	"elysium.com/applications/mockers/virtual/config"
	pb "elysium.com/pb/mockers/virtual"
)

type Handler struct {
	pb.UnimplementedVirtualServiceServer
	cfg config.Application
}

func NewHandler(cfg config.Application) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
