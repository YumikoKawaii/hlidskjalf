package virtual

import (
	"context"
	pb "elysium.com/pb/mockers/virtual"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type grpcImpl struct {
	client pb.VirtualServiceClient
}

func NewRpcClient(host string) Client {

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	conn.Connect()
	return &grpcImpl{
		client: pb.NewVirtualServiceClient(conn),
	}
}

func (c *grpcImpl) Virtual(ctx context.Context) (Response, error) {

	start := time.Now()
	resp, err := c.client.Virtual(ctx, &pb.VirtualRequest{})
	if err != nil {
		return Response{}, err
	}

	logrus.Infof("[Virtual] - Cost: %d ms", time.Since(start).Milliseconds())

	return Response{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}
