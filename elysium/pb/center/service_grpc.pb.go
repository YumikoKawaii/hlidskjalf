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

// CenterServiceClient is the client API for CenterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CenterServiceClient interface {
	GetPostsDetail(ctx context.Context, in *GetPostsDetailRequest, opts ...grpc.CallOption) (*GetPostsDetailResponse, error)
}

type centerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCenterServiceClient(cc grpc.ClientConnInterface) CenterServiceClient {
	return &centerServiceClient{cc}
}

func (c *centerServiceClient) GetPostsDetail(ctx context.Context, in *GetPostsDetailRequest, opts ...grpc.CallOption) (*GetPostsDetailResponse, error) {
	out := new(GetPostsDetailResponse)
	err := c.cc.Invoke(ctx, "/center.api.CenterService/GetPostsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CenterServiceServer is the server API for CenterService service.
// All implementations must embed UnimplementedCenterServiceServer
// for forward compatibility
type CenterServiceServer interface {
	GetPostsDetail(context.Context, *GetPostsDetailRequest) (*GetPostsDetailResponse, error)
	mustEmbedUnimplementedCenterServiceServer()
}

// UnimplementedCenterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCenterServiceServer struct {
}

func (UnimplementedCenterServiceServer) GetPostsDetail(context.Context, *GetPostsDetailRequest) (*GetPostsDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsDetail not implemented")
}
func (UnimplementedCenterServiceServer) mustEmbedUnimplementedCenterServiceServer() {}

// UnsafeCenterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CenterServiceServer will
// result in compilation errors.
type UnsafeCenterServiceServer interface {
	mustEmbedUnimplementedCenterServiceServer()
}

func RegisterCenterServiceServer(s grpc.ServiceRegistrar, srv CenterServiceServer) {
	s.RegisterService(&_CenterService_serviceDesc, srv)
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

var _CenterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "center.api.CenterService",
	HandlerType: (*CenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPostsDetail",
			Handler:    _CenterService_GetPostsDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "center/service.proto",
}
