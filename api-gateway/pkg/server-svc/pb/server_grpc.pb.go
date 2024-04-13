// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/server-svc/pb/server.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	JoinToServer(ctx context.Context, in *JoinToServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCategoryOfServer(ctx context.Context, in *GetCategoryOfServerRequest, opts ...grpc.CallOption) (*GetCategoryOfServerResponse, error)
	GetChannelsOfServer(ctx context.Context, in *GetChannelsOfServerRequest, opts ...grpc.CallOption) (*GetChannelsOfServerResponse, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerResponse, error) {
	out := new(CreateServerResponse)
	err := c.cc.Invoke(ctx, "/Server/CreateServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Server/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Server/CreateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) JoinToServer(ctx context.Context, in *JoinToServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Server/JoinToServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) GetCategoryOfServer(ctx context.Context, in *GetCategoryOfServerRequest, opts ...grpc.CallOption) (*GetCategoryOfServerResponse, error) {
	out := new(GetCategoryOfServerResponse)
	err := c.cc.Invoke(ctx, "/Server/GetCategoryOfServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) GetChannelsOfServer(ctx context.Context, in *GetChannelsOfServerRequest, opts ...grpc.CallOption) (*GetChannelsOfServerResponse, error) {
	out := new(GetChannelsOfServerResponse)
	err := c.cc.Invoke(ctx, "/Server/GetChannelsOfServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	CreateServer(context.Context, *CreateServerRequest) (*CreateServerResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*emptypb.Empty, error)
	CreateChannel(context.Context, *CreateChannelRequest) (*emptypb.Empty, error)
	JoinToServer(context.Context, *JoinToServerRequest) (*emptypb.Empty, error)
	GetCategoryOfServer(context.Context, *GetCategoryOfServerRequest) (*GetCategoryOfServerResponse, error)
	GetChannelsOfServer(context.Context, *GetChannelsOfServerRequest) (*GetChannelsOfServerResponse, error)
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) CreateServer(context.Context, *CreateServerRequest) (*CreateServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedServerServer) CreateCategory(context.Context, *CreateCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedServerServer) CreateChannel(context.Context, *CreateChannelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedServerServer) JoinToServer(context.Context, *JoinToServerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinToServer not implemented")
}
func (UnimplementedServerServer) GetCategoryOfServer(context.Context, *GetCategoryOfServerRequest) (*GetCategoryOfServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryOfServer not implemented")
}
func (UnimplementedServerServer) GetChannelsOfServer(context.Context, *GetChannelsOfServerRequest) (*GetChannelsOfServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelsOfServer not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/CreateServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).CreateServer(ctx, req.(*CreateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_JoinToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinToServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).JoinToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/JoinToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).JoinToServer(ctx, req.(*JoinToServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_GetCategoryOfServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryOfServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetCategoryOfServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/GetCategoryOfServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetCategoryOfServer(ctx, req.(*GetCategoryOfServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_GetChannelsOfServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelsOfServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetChannelsOfServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server/GetChannelsOfServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetChannelsOfServer(ctx, req.(*GetChannelsOfServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServer",
			Handler:    _Server_CreateServer_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Server_CreateCategory_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _Server_CreateChannel_Handler,
		},
		{
			MethodName: "JoinToServer",
			Handler:    _Server_JoinToServer_Handler,
		},
		{
			MethodName: "GetCategoryOfServer",
			Handler:    _Server_GetCategoryOfServer_Handler,
		},
		{
			MethodName: "GetChannelsOfServer",
			Handler:    _Server_GetChannelsOfServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/server-svc/pb/server.proto",
}
