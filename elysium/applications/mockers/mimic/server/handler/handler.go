package handler

import (
	"elysium.com/applications/mockers/mimic/config"
	pb "elysium.com/pb/mockers/mimic"
)

type Handler struct {
	pb.UnimplementedMimicServiceServer
	cfg config.Application
}

func NewHandler(cfg config.Application) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
