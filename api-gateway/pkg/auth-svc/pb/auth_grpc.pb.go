// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/auth-svc/pb/auth.proto

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	VerifyUser(ctx context.Context, in *UserVerifyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ConfirmSignup(ctx context.Context, in *ConfirmSignupRequest, opts ...grpc.CallOption) (*ConfirmSignupResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	// admin
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyUser(ctx context.Context, in *UserVerifyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/VerifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ConfirmSignup(ctx context.Context, in *ConfirmSignupRequest, opts ...grpc.CallOption) (*ConfirmSignupResponse, error) {
	out := new(ConfirmSignupResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ConfirmSignup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	out := new(AdminLoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/AdminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	VerifyUser(context.Context, *UserVerifyRequest) (*emptypb.Empty, error)
	ConfirmSignup(context.Context, *ConfirmSignupRequest) (*ConfirmSignupResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	// admin
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedAuthServiceServer) VerifyUser(context.Context, *UserVerifyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUser not implemented")
}
func (UnimplementedAuthServiceServer) ConfirmSignup(context.Context, *ConfirmSignupRequest) (*ConfirmSignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmSignup not implemented")
}
func (UnimplementedAuthServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedAuthServiceServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/VerifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyUser(ctx, req.(*UserVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ConfirmSignup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmSignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ConfirmSignup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ConfirmSignup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ConfirmSignup(ctx, req.(*ConfirmSignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/AdminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "VerifyUser",
			Handler:    _AuthService_VerifyUser_Handler,
		},
		{
			MethodName: "ConfirmSignup",
			Handler:    _AuthService_ConfirmSignup_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _AuthService_UserLogin_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _AuthService_AdminLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/auth-svc/pb/auth.proto",
}
