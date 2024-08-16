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

// EntryServiceClient is the client API for EntryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntryServiceClient interface {
	Entry(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (*EntryResponse, error)
}

type entryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEntryServiceClient(cc grpc.ClientConnInterface) EntryServiceClient {
	return &entryServiceClient{cc}
}

func (c *entryServiceClient) Entry(ctx context.Context, in *EntryRequest, opts ...grpc.CallOption) (*EntryResponse, error) {
	out := new(EntryResponse)
	err := c.cc.Invoke(ctx, "/valgrind.api.EntryService/Entry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntryServiceServer is the server API for EntryService service.
// All implementations must embed UnimplementedEntryServiceServer
// for forward compatibility
type EntryServiceServer interface {
	Entry(context.Context, *EntryRequest) (*EntryResponse, error)
	mustEmbedUnimplementedEntryServiceServer()
}

// UnimplementedEntryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEntryServiceServer struct {
}

func (UnimplementedEntryServiceServer) Entry(context.Context, *EntryRequest) (*EntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Entry not implemented")
}
func (UnimplementedEntryServiceServer) mustEmbedUnimplementedEntryServiceServer() {}

// UnsafeEntryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntryServiceServer will
// result in compilation errors.
type UnsafeEntryServiceServer interface {
	mustEmbedUnimplementedEntryServiceServer()
}

func RegisterEntryServiceServer(s grpc.ServiceRegistrar, srv EntryServiceServer) {
	s.RegisterService(&_EntryService_serviceDesc, srv)
}

func _EntryService_Entry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).Entry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/valgrind.api.EntryService/Entry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).Entry(ctx, req.(*EntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "valgrind.api.EntryService",
	HandlerType: (*EntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Entry",
			Handler:    _EntryService_Entry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "valgrind/service.proto",
}
