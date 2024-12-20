// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: mockers/virtual/service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VirtualServiceClient is the client API for VirtualService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VirtualServiceClient interface {
	Virtual(ctx context.Context, in *VirtualRequest, opts ...grpc.CallOption) (*VirtualResponse, error)
}

type virtualServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtualServiceClient(cc grpc.ClientConnInterface) VirtualServiceClient {
	return &virtualServiceClient{cc}
}

func (c *virtualServiceClient) Virtual(ctx context.Context, in *VirtualRequest, opts ...grpc.CallOption) (*VirtualResponse, error) {
	out := new(VirtualResponse)
	err := c.cc.Invoke(ctx, "/virtual.api.VirtualService/Virtual", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualServiceServer is the server API for VirtualService service.
// All implementations must embed UnimplementedVirtualServiceServer
// for forward compatibility
type VirtualServiceServer interface {
	Virtual(context.Context, *VirtualRequest) (*VirtualResponse, error)
	mustEmbedUnimplementedVirtualServiceServer()
}

// UnimplementedVirtualServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVirtualServiceServer struct {
}

func (UnimplementedVirtualServiceServer) Virtual(context.Context, *VirtualRequest) (*VirtualResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Virtual not implemented")
}
func (UnimplementedVirtualServiceServer) mustEmbedUnimplementedVirtualServiceServer() {}

// UnsafeVirtualServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VirtualServiceServer will
// result in compilation errors.
type UnsafeVirtualServiceServer interface {
	mustEmbedUnimplementedVirtualServiceServer()
}

func RegisterVirtualServiceServer(s grpc.ServiceRegistrar, srv VirtualServiceServer) {
	s.RegisterService(&VirtualService_ServiceDesc, srv)
}

func _VirtualService_Virtual_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualServiceServer).Virtual(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/virtual.api.VirtualService/Virtual",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualServiceServer).Virtual(ctx, req.(*VirtualRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VirtualService_ServiceDesc is the grpc.ServiceDesc for VirtualService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VirtualService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "virtual.api.VirtualService",
	HandlerType: (*VirtualServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Virtual",
			Handler:    _VirtualService_Virtual_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockers/virtual/service.proto",
}
