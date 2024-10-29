package handler

import (
	"elysium.com/applications/mockers/fictio/config"
	pb "elysium.com/pb/mockers/fictio"
)

type Handler struct {
	pb.UnimplementedFictioServiceServer
	cfg config.Application
}

func NewHandler(cfg config.Application) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
