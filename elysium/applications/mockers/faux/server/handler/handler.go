package handler

import (
	"elysium.com/applications/mockers/faux/config"
	pb "elysium.com/pb/mockers/faux"
)

type Handler struct {
	pb.UnimplementedFauxServiceServer
	cfg config.Application
}

func NewHandler(cfg config.Application) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
