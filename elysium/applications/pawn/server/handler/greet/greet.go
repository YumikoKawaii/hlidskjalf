package greet

import (
	"context"
	pb "elysium.com/pb/pawn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand/v2"
)

type ServiceServer struct {
	pb.UnimplementedGreetServiceServer
}

func NewServiceServer() *ServiceServer {
	return &ServiceServer{}
}

// Greet 10% request will get error
func (s *ServiceServer) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {

	number := rand.IntN(10)
	if number == 1 {
		return nil, status.Error(codes.Internal, "Unlucky")
	}

	return &pb.GreetResponse{
		Code:    int32(codes.OK),
		Message: "Hi",
	}, nil
}
