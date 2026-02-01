package acoustics

import (
	"context"

	api "github.com/YumikoKawaii/rpc.com/protobuf/acoustics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	client api.AcousticsClient
}

func NewGRPCClient(endpoint string) (Client, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &grpcClient{
		client: api.NewAcousticsClient(conn),
	}, nil
}

func (c *grpcClient) Entry(ctx context.Context, request EntryRequest) (EntryResponse, error) {
	_, err := c.client.Entry(ctx, &api.EntryRequest{})
	if err != nil {
		return EntryResponse{}, err
	}
	return EntryResponse{}, nil
}
