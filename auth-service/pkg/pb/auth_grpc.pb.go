// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/pb/auth.proto

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
	ValidateUserToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	ReSendVerificationEmail(ctx context.Context, in *ReSendVerificationEmailRequest, opts ...grpc.CallOption) (*ReSendVerificationEmailResponse, error)
	SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// User Profile
	UpdateProfilePhoto(ctx context.Context, in *UpdateprofilePhotoRequest, opts ...grpc.CallOption) (*UpdateProfilePhotoResponse, error)
	UpdateCoverPhoto(ctx context.Context, in *UpdateCoverPhotoRequest, opts ...grpc.CallOption) (*UpdateCoverPhotoResponse, error)
	DeletePhotoInProfile(ctx context.Context, in *DeletePhotoInProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserProfileStatus(ctx context.Context, in *UpdateUserProfileStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateUserProfileDescription(ctx context.Context, in *UpdateUserProfileDescriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SerchUsers(ctx context.Context, in *SerchUsersRequest, opts ...grpc.CallOption) (*SerchUsersResponse, error)
	// admin
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
	ValidateAdminToken(ctx context.Context, in *ValidateAdminTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BlockUse(ctx context.Context, in *BlockUseRequest, opts ...grpc.CallOption) (*BlockUseResponse, error)
	UnBlockUser(ctx context.Context, in *UnBlockUserRequest, opts ...grpc.CallOption) (*UnBlockUserResponse, error)
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

func (c *authServiceClient) ValidateUserToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ValidateUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ReSendVerificationEmail(ctx context.Context, in *ReSendVerificationEmailRequest, opts ...grpc.CallOption) (*ReSendVerificationEmailResponse, error) {
	out := new(ReSendVerificationEmailResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ReSendVerificationEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SendOtp(ctx context.Context, in *SendOtpRequest, opts ...grpc.CallOption) (*SendOtpResponse, error) {
	out := new(SendOtpResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SendOtp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateProfilePhoto(ctx context.Context, in *UpdateprofilePhotoRequest, opts ...grpc.CallOption) (*UpdateProfilePhotoResponse, error) {
	out := new(UpdateProfilePhotoResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateProfilePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateCoverPhoto(ctx context.Context, in *UpdateCoverPhotoRequest, opts ...grpc.CallOption) (*UpdateCoverPhotoResponse, error) {
	out := new(UpdateCoverPhotoResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateCoverPhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeletePhotoInProfile(ctx context.Context, in *DeletePhotoInProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeletePhotoInProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserProfileStatus(ctx context.Context, in *UpdateUserProfileStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateUserProfileStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUserProfileDescription(ctx context.Context, in *UpdateUserProfileDescriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateUserProfileDescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserProfile(ctx context.Context, in *UserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SerchUsers(ctx context.Context, in *SerchUsersRequest, opts ...grpc.CallOption) (*SerchUsersResponse, error) {
	out := new(SerchUsersResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SerchUsers", in, out, opts...)
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

func (c *authServiceClient) ValidateAdminToken(ctx context.Context, in *ValidateAdminTokenRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ValidateAdminToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) BlockUse(ctx context.Context, in *BlockUseRequest, opts ...grpc.CallOption) (*BlockUseResponse, error) {
	out := new(BlockUseResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/BlockUse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UnBlockUser(ctx context.Context, in *UnBlockUserRequest, opts ...grpc.CallOption) (*UnBlockUserResponse, error) {
	out := new(UnBlockUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UnBlockUser", in, out, opts...)
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
	ValidateUserToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	ReSendVerificationEmail(context.Context, *ReSendVerificationEmailRequest) (*ReSendVerificationEmailResponse, error)
	SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*emptypb.Empty, error)
	// User Profile
	UpdateProfilePhoto(context.Context, *UpdateprofilePhotoRequest) (*UpdateProfilePhotoResponse, error)
	UpdateCoverPhoto(context.Context, *UpdateCoverPhotoRequest) (*UpdateCoverPhotoResponse, error)
	DeletePhotoInProfile(context.Context, *DeletePhotoInProfileRequest) (*emptypb.Empty, error)
	UpdateUserProfileStatus(context.Context, *UpdateUserProfileStatusRequest) (*emptypb.Empty, error)
	UpdateUserProfileDescription(context.Context, *UpdateUserProfileDescriptionRequest) (*emptypb.Empty, error)
	UserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*emptypb.Empty, error)
	SerchUsers(context.Context, *SerchUsersRequest) (*SerchUsersResponse, error)
	// admin
	AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error)
	ValidateAdminToken(context.Context, *ValidateAdminTokenRequest) (*emptypb.Empty, error)
	BlockUse(context.Context, *BlockUseRequest) (*BlockUseResponse, error)
	UnBlockUser(context.Context, *UnBlockUserRequest) (*UnBlockUserResponse, error)
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
func (UnimplementedAuthServiceServer) ValidateUserToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUserToken not implemented")
}
func (UnimplementedAuthServiceServer) ReSendVerificationEmail(context.Context, *ReSendVerificationEmailRequest) (*ReSendVerificationEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReSendVerificationEmail not implemented")
}
func (UnimplementedAuthServiceServer) SendOtp(context.Context, *SendOtpRequest) (*SendOtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOtp not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) UpdateProfilePhoto(context.Context, *UpdateprofilePhotoRequest) (*UpdateProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfilePhoto not implemented")
}
func (UnimplementedAuthServiceServer) UpdateCoverPhoto(context.Context, *UpdateCoverPhotoRequest) (*UpdateCoverPhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoverPhoto not implemented")
}
func (UnimplementedAuthServiceServer) DeletePhotoInProfile(context.Context, *DeletePhotoInProfileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhotoInProfile not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserProfileStatus(context.Context, *UpdateUserProfileStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfileStatus not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUserProfileDescription(context.Context, *UpdateUserProfileDescriptionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfileDescription not implemented")
}
func (UnimplementedAuthServiceServer) UserProfile(context.Context, *UserProfileRequest) (*UserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserProfile not implemented")
}
func (UnimplementedAuthServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAuthServiceServer) SerchUsers(context.Context, *SerchUsersRequest) (*SerchUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SerchUsers not implemented")
}
func (UnimplementedAuthServiceServer) AdminLogin(context.Context, *AdminLoginRequest) (*AdminLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAuthServiceServer) ValidateAdminToken(context.Context, *ValidateAdminTokenRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateAdminToken not implemented")
}
func (UnimplementedAuthServiceServer) BlockUse(context.Context, *BlockUseRequest) (*BlockUseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUse not implemented")
}
func (UnimplementedAuthServiceServer) UnBlockUser(context.Context, *UnBlockUserRequest) (*UnBlockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBlockUser not implemented")
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

func _AuthService_ValidateUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ValidateUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateUserToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ReSendVerificationEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReSendVerificationEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ReSendVerificationEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ReSendVerificationEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ReSendVerificationEmail(ctx, req.(*ReSendVerificationEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SendOtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SendOtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SendOtp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SendOtp(ctx, req.(*SendOtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateProfilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateprofilePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateProfilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateProfilePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateProfilePhoto(ctx, req.(*UpdateprofilePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateCoverPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoverPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateCoverPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateCoverPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateCoverPhoto(ctx, req.(*UpdateCoverPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeletePhotoInProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhotoInProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeletePhotoInProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeletePhotoInProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeletePhotoInProfile(ctx, req.(*DeletePhotoInProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserProfileStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserProfileStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateUserProfileStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserProfileStatus(ctx, req.(*UpdateUserProfileStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUserProfileDescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileDescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUserProfileDescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateUserProfileDescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUserProfileDescription(ctx, req.(*UpdateUserProfileDescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserProfile(ctx, req.(*UserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SerchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SerchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SerchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SerchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SerchUsers(ctx, req.(*SerchUsersRequest))
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

func _AuthService_ValidateAdminToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateAdminTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateAdminToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ValidateAdminToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateAdminToken(ctx, req.(*ValidateAdminTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_BlockUse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).BlockUse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/BlockUse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).BlockUse(ctx, req.(*BlockUseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UnBlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnBlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UnBlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UnBlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UnBlockUser(ctx, req.(*UnBlockUserRequest))
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
			MethodName: "ValidateUserToken",
			Handler:    _AuthService_ValidateUserToken_Handler,
		},
		{
			MethodName: "ReSendVerificationEmail",
			Handler:    _AuthService_ReSendVerificationEmail_Handler,
		},
		{
			MethodName: "SendOtp",
			Handler:    _AuthService_SendOtp_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _AuthService_ForgotPassword_Handler,
		},
		{
			MethodName: "UpdateProfilePhoto",
			Handler:    _AuthService_UpdateProfilePhoto_Handler,
		},
		{
			MethodName: "UpdateCoverPhoto",
			Handler:    _AuthService_UpdateCoverPhoto_Handler,
		},
		{
			MethodName: "DeletePhotoInProfile",
			Handler:    _AuthService_DeletePhotoInProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfileStatus",
			Handler:    _AuthService_UpdateUserProfileStatus_Handler,
		},
		{
			MethodName: "UpdateUserProfileDescription",
			Handler:    _AuthService_UpdateUserProfileDescription_Handler,
		},
		{
			MethodName: "UserProfile",
			Handler:    _AuthService_UserProfile_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AuthService_DeleteAccount_Handler,
		},
		{
			MethodName: "SerchUsers",
			Handler:    _AuthService_SerchUsers_Handler,
		},
		{
			MethodName: "AdminLogin",
			Handler:    _AuthService_AdminLogin_Handler,
		},
		{
			MethodName: "ValidateAdminToken",
			Handler:    _AuthService_ValidateAdminToken_Handler,
		},
		{
			MethodName: "BlockUse",
			Handler:    _AuthService_BlockUse_Handler,
		},
		{
			MethodName: "UnBlockUser",
			Handler:    _AuthService_UnBlockUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth.proto",
}
