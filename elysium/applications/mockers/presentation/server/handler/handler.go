package handler

import (
	"elysium.com/applications/mockers/presentation/config"
	"elysium.com/applications/mockers/presentation/pkg/adapters/echo"
	"elysium.com/applications/mockers/presentation/pkg/adapters/faux"
	"elysium.com/applications/mockers/presentation/pkg/adapters/fictio"
	"elysium.com/applications/mockers/presentation/pkg/adapters/mimic"
	"elysium.com/applications/mockers/presentation/pkg/adapters/virtual"
	pb "elysium.com/pb/mockers/presentation"
)

type Handler struct {
	pb.UnimplementedPresentationServiceServer
	echoClient    echo.Client
	fauxClient    faux.Client
	fictioClient  fictio.Client
	mimicClient   mimic.Client
	virtualClient virtual.Client
	cfg           config.Application
}

func NewHandler(cfg config.Application, client echo.Client, client2 faux.Client, client3 fictio.Client, client4 mimic.Client, client5 virtual.Client) *Handler {
	return &Handler{
		cfg:           cfg,
		echoClient:    client,
		fauxClient:    client2,
		fictioClient:  client3,
		mimicClient:   client4,
		virtualClient: client5,
	}
}
