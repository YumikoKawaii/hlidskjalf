package echo

import (
	"context"
	pb "elysium.com/pb/mockers/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type grpcImpl struct {
	client pb.EchoServiceClient
}

func NewRpcClient(host string) Client {

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	conn.Connect()
	return &grpcImpl{
		client: pb.NewEchoServiceClient(conn),
	}
}

func (c *grpcImpl) Echo(ctx context.Context) (Response, error) {

	start := time.Now()
	resp, err := c.client.Echo(ctx, &pb.EchoRequest{})
	if err != nil {
		return Response{}, err
	}

	logrus.Infof("[Echo] - Cost: %d ms", time.Since(start).Milliseconds())

	return Response{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

//func (c *grpcImpl) protoStats(protos []*pb.Stat) []Stat {
//	stats := make([]Stat, 0)
//	for _, proto := range protos {
//		stats = append(stats, Stat{
//			First:   proto.First,
//			Second:  proto.Second,
//			Third:   proto.Third,
//			Fourth:  proto.Fourth,
//			Fifth:   proto.Fifth,
//			Sixth:   proto.Sixth,
//			Seventh: proto.Seventh,
//			Eighth:  proto.Eighth,
//			Ninth:   proto.Ninth,
//			Tenth:   proto.Tenth,
//		})
//	}
//	return stats
//}
