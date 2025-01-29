package main

import (
	"context"
	pb "elysium.com/pb/observer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	cc, err := grpc.NewClient("localhost:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	observerClient := pb.NewObserverClient(cc)
	ctx := context.Background()
	go func() {
		respStream, err := observerClient.ViewStream(context.Background(), &pb.ViewStreamRequest{
			Session: "123",
		})
		if err != nil {
			panic(err)
		}

		for {
			resp, err := respStream.Recv()
			if err != nil {
				panic(err)
			}
			logrus.Info(resp.View)
		}
	}()

	go func() {
		respStream, err := observerClient.ViewStream(context.Background(), &pb.ViewStreamRequest{
			Session: "123",
		})
		if err != nil {
			panic(err)
		}

		for {
			resp, err := respStream.Recv()
			if err != nil {
				panic(err)
			}
			logrus.Info(resp.View)
		}
	}()

	<-ctx.Done()
}
