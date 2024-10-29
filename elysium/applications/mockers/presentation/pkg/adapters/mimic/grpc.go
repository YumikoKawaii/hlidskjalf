package mimic

import (
	"context"
	pb "elysium.com/pb/mockers/mimic"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type grpcImpl struct {
	client pb.MimicServiceClient
}

func NewRpcClient(host string) Client {

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	conn.Connect()
	return &grpcImpl{
		client: pb.NewMimicServiceClient(conn),
	}
}

func (c *grpcImpl) Mimic(ctx context.Context) (Response, error) {

	start := time.Now()
	resp, err := c.client.Mimic(ctx, &pb.MimicRequest{})
	if err != nil {
		return Response{}, err
	}

	logrus.Infof("[Mimic] - Cost: %d ms", time.Since(start).Milliseconds())

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