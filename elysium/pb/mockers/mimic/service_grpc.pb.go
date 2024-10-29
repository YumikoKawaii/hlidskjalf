// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MimicServiceClient is the client API for MimicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MimicServiceClient interface {
	Mimic(ctx context.Context, in *MimicRequest, opts ...grpc.CallOption) (*MimicResponse, error)
}

type mimicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMimicServiceClient(cc grpc.ClientConnInterface) MimicServiceClient {
	return &mimicServiceClient{cc}
}

func (c *mimicServiceClient) Mimic(ctx context.Context, in *MimicRequest, opts ...grpc.CallOption) (*MimicResponse, error) {
	out := new(MimicResponse)
	err := c.cc.Invoke(ctx, "/mimic.api.MimicService/Mimic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MimicServiceServer is the server API for MimicService service.
// All implementations must embed UnimplementedMimicServiceServer
// for forward compatibility
type MimicServiceServer interface {
	Mimic(context.Context, *MimicRequest) (*MimicResponse, error)
	mustEmbedUnimplementedMimicServiceServer()
}

// UnimplementedMimicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMimicServiceServer struct {
}

func (UnimplementedMimicServiceServer) Mimic(context.Context, *MimicRequest) (*MimicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mimic not implemented")
}
func (UnimplementedMimicServiceServer) mustEmbedUnimplementedMimicServiceServer() {}

// UnsafeMimicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MimicServiceServer will
// result in compilation errors.
type UnsafeMimicServiceServer interface {
	mustEmbedUnimplementedMimicServiceServer()
}

func RegisterMimicServiceServer(s grpc.ServiceRegistrar, srv MimicServiceServer) {
	s.RegisterService(&_MimicService_serviceDesc, srv)
}

func _MimicService_Mimic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MimicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MimicServiceServer).Mimic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mimic.api.MimicService/Mimic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MimicServiceServer).Mimic(ctx, req.(*MimicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MimicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mimic.api.MimicService",
	HandlerType: (*MimicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mimic",
			Handler:    _MimicService_Mimic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockers/mimic/service.proto",
}