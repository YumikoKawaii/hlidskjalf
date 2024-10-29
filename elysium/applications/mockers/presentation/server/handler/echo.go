package handler

import (
	"context"
	pb "elysium.com/pb/mockers/presentation"
)

func (c *Handler) Echo(ctx context.Context, request *pb.Request) (*pb.Response, error) {

	resp, err := c.echoClient.Echo(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}
