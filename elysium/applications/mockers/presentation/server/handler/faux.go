package handler

import (
	"context"
	pb "elysium.com/pb/mockers/presentation"
)

func (c *Handler) Faux(ctx context.Context, request *pb.Request) (*pb.Response, error) {

	resp, err := c.fauxClient.Faux(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}
