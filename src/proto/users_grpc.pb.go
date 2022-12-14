// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: users.proto

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

// FunctionsClient is the client API for Functions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FunctionsClient interface {
	// Send a user id and get his information
	GetInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
	// Send a user id and change in manner variable in the server
	ChangeManner(ctx context.Context, in *MannerRequest, opts ...grpc.CallOption) (*MannerReply, error)
}

type functionsClient struct {
	cc grpc.ClientConnInterface
}

func NewFunctionsClient(cc grpc.ClientConnInterface) FunctionsClient {
	return &functionsClient{cc}
}

func (c *functionsClient) GetInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/users.Functions/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *functionsClient) ChangeManner(ctx context.Context, in *MannerRequest, opts ...grpc.CallOption) (*MannerReply, error) {
	out := new(MannerReply)
	err := c.cc.Invoke(ctx, "/users.Functions/ChangeManner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FunctionsServer is the server API for Functions service.
// All implementations must embed UnimplementedFunctionsServer
// for forward compatibility
type FunctionsServer interface {
	// Send a user id and get his information
	GetInfo(context.Context, *UserRequest) (*UserReply, error)
	// Send a user id and change in manner variable in the server
	ChangeManner(context.Context, *MannerRequest) (*MannerReply, error)
	mustEmbedUnimplementedFunctionsServer()
}

// UnimplementedFunctionsServer must be embedded to have forward compatible implementations.
type UnimplementedFunctionsServer struct {
}

func (UnimplementedFunctionsServer) GetInfo(context.Context, *UserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedFunctionsServer) ChangeManner(context.Context, *MannerRequest) (*MannerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeManner not implemented")
}
func (UnimplementedFunctionsServer) mustEmbedUnimplementedFunctionsServer() {}

// UnsafeFunctionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FunctionsServer will
// result in compilation errors.
type UnsafeFunctionsServer interface {
	mustEmbedUnimplementedFunctionsServer()
}

func RegisterFunctionsServer(s grpc.ServiceRegistrar, srv FunctionsServer) {
	s.RegisterService(&Functions_ServiceDesc, srv)
}

func _Functions_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Functions/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServer).GetInfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Functions_ChangeManner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunctionsServer).ChangeManner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Functions/ChangeManner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunctionsServer).ChangeManner(ctx, req.(*MannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Functions_ServiceDesc is the grpc.ServiceDesc for Functions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Functions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Functions",
	HandlerType: (*FunctionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _Functions_GetInfo_Handler,
		},
		{
			MethodName: "ChangeManner",
			Handler:    _Functions_ChangeManner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
