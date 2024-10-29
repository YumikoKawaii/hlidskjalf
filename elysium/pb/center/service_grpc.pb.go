// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: center/service.proto

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

// CenterServiceClient is the client API for CenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CenterServiceClient interface {
	DiscoveryPosts(ctx context.Context, in *DiscoveryPostsRequest, opts ...grpc.CallOption) (*DiscoveryPostsResponse, error)
	GetPostsDetail(ctx context.Context, in *GetPostsDetailRequest, opts ...grpc.CallOption) (*GetPostsDetailResponse, error)
	UpsertPost(ctx context.Context, in *UpsertPostRequest, opts ...grpc.CallOption) (*UpsertPostResponse, error)
}

type centerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCenterServiceClient(cc grpc.ClientConnInterface) CenterServiceClient {
	return &centerServiceClient{cc}
}

func (c *centerServiceClient) DiscoveryPosts(ctx context.Context, in *DiscoveryPostsRequest, opts ...grpc.CallOption) (*DiscoveryPostsResponse, error) {
	out := new(DiscoveryPostsResponse)
	err := c.cc.Invoke(ctx, "/center.api.CenterService/DiscoveryPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) GetPostsDetail(ctx context.Context, in *GetPostsDetailRequest, opts ...grpc.CallOption) (*GetPostsDetailResponse, error) {
	out := new(GetPostsDetailResponse)
	err := c.cc.Invoke(ctx, "/center.api.CenterService/GetPostsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) UpsertPost(ctx context.Context, in *UpsertPostRequest, opts ...grpc.CallOption) (*UpsertPostResponse, error) {
	out := new(UpsertPostResponse)
	err := c.cc.Invoke(ctx, "/center.api.CenterService/UpsertPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CenterServiceServer is the server API for CenterService service.
// All implementations must embed UnimplementedCenterServiceServer
// for forward compatibility
type CenterServiceServer interface {
	DiscoveryPosts(context.Context, *DiscoveryPostsRequest) (*DiscoveryPostsResponse, error)
	GetPostsDetail(context.Context, *GetPostsDetailRequest) (*GetPostsDetailResponse, error)
	UpsertPost(context.Context, *UpsertPostRequest) (*UpsertPostResponse, error)
	mustEmbedUnimplementedCenterServiceServer()
}

// UnimplementedCenterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCenterServiceServer struct {
}

func (UnimplementedCenterServiceServer) DiscoveryPosts(context.Context, *DiscoveryPostsRequest) (*DiscoveryPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoveryPosts not implemented")
}
func (UnimplementedCenterServiceServer) GetPostsDetail(context.Context, *GetPostsDetailRequest) (*GetPostsDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsDetail not implemented")
}
func (UnimplementedCenterServiceServer) UpsertPost(context.Context, *UpsertPostRequest) (*UpsertPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertPost not implemented")
}
func (UnimplementedCenterServiceServer) mustEmbedUnimplementedCenterServiceServer() {}

// UnsafeCenterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CenterServiceServer will
// result in compilation errors.
type UnsafeCenterServiceServer interface {
	mustEmbedUnimplementedCenterServiceServer()
}

func RegisterCenterServiceServer(s grpc.ServiceRegistrar, srv CenterServiceServer) {
	s.RegisterService(&CenterService_ServiceDesc, srv)
}

func _CenterService_DiscoveryPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).DiscoveryPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/center.api.CenterService/DiscoveryPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).DiscoveryPosts(ctx, req.(*DiscoveryPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_GetPostsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).GetPostsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/center.api.CenterService/GetPostsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).GetPostsDetail(ctx, req.(*GetPostsDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_UpsertPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).UpsertPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/center.api.CenterService/UpsertPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).UpsertPost(ctx, req.(*UpsertPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CenterService_ServiceDesc is the grpc.ServiceDesc for CenterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CenterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "center.api.CenterService",
	HandlerType: (*CenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoveryPosts",
			Handler:    _CenterService_DiscoveryPosts_Handler,
		},
		{
			MethodName: "GetPostsDetail",
			Handler:    _CenterService_GetPostsDetail_Handler,
		},
		{
			MethodName: "UpsertPost",
			Handler:    _CenterService_UpsertPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "center/service.proto",
}
