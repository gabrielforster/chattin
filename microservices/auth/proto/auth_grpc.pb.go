// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/auth.proto

package proto

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

// SessionDataServiceClient is the client API for SessionDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionDataServiceClient interface {
	GetSessionData(ctx context.Context, in *GetSessionDataRequest, opts ...grpc.CallOption) (*GetSessionDataResponse, error)
}

type sessionDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionDataServiceClient(cc grpc.ClientConnInterface) SessionDataServiceClient {
	return &sessionDataServiceClient{cc}
}

func (c *sessionDataServiceClient) GetSessionData(ctx context.Context, in *GetSessionDataRequest, opts ...grpc.CallOption) (*GetSessionDataResponse, error) {
	out := new(GetSessionDataResponse)
	err := c.cc.Invoke(ctx, "/proto.SessionDataService/GetSessionData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionDataServiceServer is the server API for SessionDataService service.
// All implementations must embed UnimplementedSessionDataServiceServer
// for forward compatibility
type SessionDataServiceServer interface {
	GetSessionData(context.Context, *GetSessionDataRequest) (*GetSessionDataResponse, error)
	mustEmbedUnimplementedSessionDataServiceServer()
}

// UnimplementedSessionDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionDataServiceServer struct {
}

func (UnimplementedSessionDataServiceServer) GetSessionData(context.Context, *GetSessionDataRequest) (*GetSessionDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionData not implemented")
}
func (UnimplementedSessionDataServiceServer) mustEmbedUnimplementedSessionDataServiceServer() {}

// UnsafeSessionDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionDataServiceServer will
// result in compilation errors.
type UnsafeSessionDataServiceServer interface {
	mustEmbedUnimplementedSessionDataServiceServer()
}

func RegisterSessionDataServiceServer(s grpc.ServiceRegistrar, srv SessionDataServiceServer) {
	s.RegisterService(&SessionDataService_ServiceDesc, srv)
}

func _SessionDataService_GetSessionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionDataServiceServer).GetSessionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionDataService/GetSessionData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionDataServiceServer).GetSessionData(ctx, req.(*GetSessionDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionDataService_ServiceDesc is the grpc.ServiceDesc for SessionDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SessionDataService",
	HandlerType: (*SessionDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSessionData",
			Handler:    _SessionDataService_GetSessionData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}