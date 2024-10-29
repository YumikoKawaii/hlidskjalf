package handler

import (
	"elysium.com/applications/mockers/echo/config"
	pb "elysium.com/pb/mockers/echo"
)

type Handler struct {
	pb.UnimplementedEchoServiceServer
	cfg config.Application
}

func NewHandler(cfg config.Application) *Handler {
	return &Handler{
		cfg: cfg,
	}
}
