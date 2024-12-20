// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: authenticator/service.proto

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

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticatorClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UpdatePermissions(ctx context.Context, in *UpdatePermissionsRequest, opts ...grpc.CallOption) (*UpdatePermissionsResponse, error)
	GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type authenticatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticatorClient(cc grpc.ClientConnInterface) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/authenticator.api.Authenticator/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/authenticator.api.Authenticator/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) UpdatePermissions(ctx context.Context, in *UpdatePermissionsRequest, opts ...grpc.CallOption) (*UpdatePermissionsResponse, error) {
	out := new(UpdatePermissionsResponse)
	err := c.cc.Invoke(ctx, "/authenticator.api.Authenticator/UpdatePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) GetPermissions(ctx context.Context, in *GetPermissionsRequest, opts ...grpc.CallOption) (*GetPermissionsResponse, error) {
	out := new(GetPermissionsResponse)
	err := c.cc.Invoke(ctx, "/authenticator.api.Authenticator/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/authenticator.api.Authenticator/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
// All implementations must embed UnimplementedAuthenticatorServer
// for forward compatibility
type AuthenticatorServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	UpdatePermissions(context.Context, *UpdatePermissionsRequest) (*UpdatePermissionsResponse, error)
	GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedAuthenticatorServer()
}

// UnimplementedAuthenticatorServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorServer struct {
}

func (UnimplementedAuthenticatorServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedAuthenticatorServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticatorServer) UpdatePermissions(context.Context, *UpdatePermissionsRequest) (*UpdatePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePermissions not implemented")
}
func (UnimplementedAuthenticatorServer) GetPermissions(context.Context, *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}
func (UnimplementedAuthenticatorServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthenticatorServer) mustEmbedUnimplementedAuthenticatorServer() {}

// UnsafeAuthenticatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticatorServer will
// result in compilation errors.
type UnsafeAuthenticatorServer interface {
	mustEmbedUnimplementedAuthenticatorServer()
}

func RegisterAuthenticatorServer(s grpc.ServiceRegistrar, srv AuthenticatorServer) {
	s.RegisterService(&Authenticator_ServiceDesc, srv)
}

func _Authenticator_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticator.api.Authenticator/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticator.api.Authenticator/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_UpdatePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).UpdatePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticator.api.Authenticator/UpdatePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).UpdatePermissions(ctx, req.(*UpdatePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticator.api.Authenticator/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).GetPermissions(ctx, req.(*GetPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authenticator.api.Authenticator/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authenticator_ServiceDesc is the grpc.ServiceDesc for Authenticator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authenticator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authenticator.api.Authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _Authenticator_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authenticator_Login_Handler,
		},
		{
			MethodName: "UpdatePermissions",
			Handler:    _Authenticator_UpdatePermissions_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _Authenticator_GetPermissions_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Authenticator_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authenticator/service.proto",
}
