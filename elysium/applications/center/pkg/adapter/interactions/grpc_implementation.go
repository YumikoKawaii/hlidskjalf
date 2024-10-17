package interactions

import (
	"context"
	"elysium.com/applications/utils"
	pb "elysium.com/pb/interactions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcClient struct {
	rpc pb.InteractionServiceClient
}

func NewRpcClient(config Config) Client {

	conn, err := grpc.NewClient(config.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewInteractionServiceClient(conn)
	return &grpcClient{
		rpc: client,
	}
}

func (c *grpcClient) UpsertInteraction(ctx context.Context, request UpsertInteractionRequest) (UpsertInteractionResponse, error) {

	resp, err := c.rpc.UpsertInteraction(ctx, &pb.UpsertInteractionRequest{
		Id:      utils.UInt32PointerToProto(request.Id),
		PostId:  request.PostId,
		Author:  request.Author,
		Type:    request.Type,
		Content: request.Content,
	})
	if err != nil {
		return UpsertInteractionResponse{}, err
	}

	return UpsertInteractionResponse{
		Code:          resp.Code,
		Message:       resp.Message,
		InteractionId: resp.InteractionId,
	}, nil
}

func (c *grpcClient) GetInteractions(ctx context.Context, request GetInteractionRequest) (GetInteractionResponse, error) {

	resp, err := c.rpc.GetInteractions(ctx, &pb.GetInteractionsRequest{
		PostIds:  request.PostIds,
		Page:     request.Page,
		PageSize: request.PageSize,
	})
	if err != nil {
		return GetInteractionResponse{}, err
	}

	return GetInteractionResponse{
		Code:    resp.Code,
		Message: resp.Message,
		Data: struct {
			Interactions []Interaction `json:"interactions,omitempty"`
			Page         int32         `json:"page,omitempty"`
			PageSize     int32         `json:"pageSize,omitempty"`
		}{
			Interactions: nil,
			Page:         resp.Data.Page,
			PageSize:     resp.Data.PageSize,
		},
	}, err
}

func (c *grpcClient) convertInteractionsProtoToInteractions(protos []*pb.Interaction) []Interaction {
	interactions := make([]Interaction, 0)
	for _, proto := range protos {
		interactions = append(interactions, Interaction{
			Id:        proto.Id,
			PostId:    proto.PostId,
			Author:    proto.Author,
			Type:      proto.Type,
			Content:   proto.Content,
			CreatedAt: proto.CreatedAt,
			UpdatedAt: proto.UpdatedAt,
		})
	}

	return interactions
}
