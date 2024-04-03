// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/friend-svc/pb/friend.proto

package pb

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

// FriendServiceClient is the client API for FriendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendServiceClient interface {
	FriendRequest(ctx context.Context, in *FriendRequestRequest, opts ...grpc.CallOption) (*FriendRequestResponse, error)
	FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	GetReceivedFriendRequest(ctx context.Context, in *GetReceivedFriendRequestRequest, opts ...grpc.CallOption) (*GetReceivedFriendRequestResponse, error)
	GetSendFriendRequest(ctx context.Context, in *GetSendFriendRequestRequest, opts ...grpc.CallOption) (*GetSendFriendRequestResponse, error)
}

type friendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendServiceClient(cc grpc.ClientConnInterface) FriendServiceClient {
	return &friendServiceClient{cc}
}

func (c *friendServiceClient) FriendRequest(ctx context.Context, in *FriendRequestRequest, opts ...grpc.CallOption) (*FriendRequestResponse, error) {
	out := new(FriendRequestResponse)
	err := c.cc.Invoke(ctx, "/FriendService/FriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, "/FriendService/FriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) GetReceivedFriendRequest(ctx context.Context, in *GetReceivedFriendRequestRequest, opts ...grpc.CallOption) (*GetReceivedFriendRequestResponse, error) {
	out := new(GetReceivedFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/FriendService/GetReceivedFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) GetSendFriendRequest(ctx context.Context, in *GetSendFriendRequestRequest, opts ...grpc.CallOption) (*GetSendFriendRequestResponse, error) {
	out := new(GetSendFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/FriendService/GetSendFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServiceServer is the server API for FriendService service.
// All implementations must embed UnimplementedFriendServiceServer
// for forward compatibility
type FriendServiceServer interface {
	FriendRequest(context.Context, *FriendRequestRequest) (*FriendRequestResponse, error)
	FriendList(context.Context, *FriendListRequest) (*FriendListResponse, error)
	GetReceivedFriendRequest(context.Context, *GetReceivedFriendRequestRequest) (*GetReceivedFriendRequestResponse, error)
	GetSendFriendRequest(context.Context, *GetSendFriendRequestRequest) (*GetSendFriendRequestResponse, error)
	mustEmbedUnimplementedFriendServiceServer()
}

// UnimplementedFriendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFriendServiceServer struct {
}

func (UnimplementedFriendServiceServer) FriendRequest(context.Context, *FriendRequestRequest) (*FriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendRequest not implemented")
}
func (UnimplementedFriendServiceServer) FriendList(context.Context, *FriendListRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendList not implemented")
}
func (UnimplementedFriendServiceServer) GetReceivedFriendRequest(context.Context, *GetReceivedFriendRequestRequest) (*GetReceivedFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceivedFriendRequest not implemented")
}
func (UnimplementedFriendServiceServer) GetSendFriendRequest(context.Context, *GetSendFriendRequestRequest) (*GetSendFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendFriendRequest not implemented")
}
func (UnimplementedFriendServiceServer) mustEmbedUnimplementedFriendServiceServer() {}

// UnsafeFriendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServiceServer will
// result in compilation errors.
type UnsafeFriendServiceServer interface {
	mustEmbedUnimplementedFriendServiceServer()
}

func RegisterFriendServiceServer(s grpc.ServiceRegistrar, srv FriendServiceServer) {
	s.RegisterService(&FriendService_ServiceDesc, srv)
}

func _FriendService_FriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).FriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/FriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).FriendRequest(ctx, req.(*FriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/FriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).FriendList(ctx, req.(*FriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_GetReceivedFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceivedFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).GetReceivedFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/GetReceivedFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).GetReceivedFriendRequest(ctx, req.(*GetReceivedFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_GetSendFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).GetSendFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FriendService/GetSendFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).GetSendFriendRequest(ctx, req.(*GetSendFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendService_ServiceDesc is the grpc.ServiceDesc for FriendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FriendService",
	HandlerType: (*FriendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FriendRequest",
			Handler:    _FriendService_FriendRequest_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _FriendService_FriendList_Handler,
		},
		{
			MethodName: "GetReceivedFriendRequest",
			Handler:    _FriendService_GetReceivedFriendRequest_Handler,
		},
		{
			MethodName: "GetSendFriendRequest",
			Handler:    _FriendService_GetSendFriendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/friend-svc/pb/friend.proto",
}
