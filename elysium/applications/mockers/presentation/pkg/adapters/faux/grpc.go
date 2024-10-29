package faux

import (
	"context"
	pb "elysium.com/pb/mockers/faux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type grpcImpl struct {
	client pb.FauxServiceClient
}

func NewRpcClient(host string) Client {

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	conn.Connect()
	return &grpcImpl{
		client: pb.NewFauxServiceClient(conn),
	}
}

func (c *grpcImpl) Faux(ctx context.Context) (Response, error) {

	start := time.Now()
	resp, err := c.client.Faux(ctx, &pb.FauxRequest{})
	if err != nil {
		return Response{}, err
	}

	logrus.Infof("[Faux] - Cost: %d ms", time.Since(start).Milliseconds())

	return Response{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}
