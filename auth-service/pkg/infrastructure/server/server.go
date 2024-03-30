package server_auth_server

import (
	"context"

	requestmodell_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	"github.com/Vajid-Hussain/HiperHive/auth-service/pkg/pb"
	interface_usecase_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userUseCase  interface_usecase_auth_server.IUserUseCase
	adminUseCase interface_usecase_auth_server.IAdminUseCase
}

func NewAuthServer(userUseCase interface_usecase_auth_server.IUserUseCase, adminUseCase interface_usecase_auth_server.IAdminUseCase) *AuthServer {
	return &AuthServer{userUseCase: userUseCase,
		adminUseCase: adminUseCase}
}

func (u *AuthServer) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	var UserDetails requestmodell_auth_server.UserSignup

	UserDetails.ConfirmPassword = req.ConfirmPassword
	UserDetails.Email = req.Email
	UserDetails.Password = req.Password
	UserDetails.UserName = req.UserName
	UserDetails.Name = req.Name

	userReq, err := u.userUseCase.Signup(UserDetails)
	if err != nil {
		return nil, err
	}

	return &pb.SignupResponse{
		UserID:          userReq.ID,
		UserName:        userReq.UserName,
		Name:            userReq.Name,
		Email:           userReq.Email,
		TemperveryToken: userReq.TemperveryToken,
	}, nil
}

func (u *AuthServer) VerifyUser(ctx context.Context, req *pb.UserVerifyRequest) (*emptypb.Empty, error) {
	err := u.userUseCase.VerifyUserSignup(req.Email, req.Token)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *AuthServer) ConfirmSignup(ctx context.Context, req *pb.ConfirmSignupRequest) (*pb.ConfirmSignupResponse, error) {
	result, err := u.userUseCase.ConfirmSignup(req.TemperveryToken)
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmSignupResponse{
		AccessToken: result.AccesToken,
		RefresToken: result.RefreshToken,
	}, nil
}

func (u *AuthServer) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	result, err := u.userUseCase.UserLogin(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UserLoginResponse{
		AccessToken: result.AccesToken,
		RefresToken: result.RefreshToken,
	}, nil
}

func (u *AuthServer) UpdateProfilePhoto(ctx context.Context, req *pb.UpdateprofilePhotoRequest) (*pb.UpdateProfilePhotoResponse, error) {
	url, err := u.userUseCase.UpdateProfilePhoto(req.UserID, req.Image)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfilePhotoResponse{
		Url: url,
	}, nil
}

func (u *AuthServer) UpdateCoverPhoto(ctx context.Context, req *pb.UpdateCoverPhotoRequest) (*pb.UpdateCoverPhotoResponse, error) {
	url, err := u.userUseCase.UpdateCoverPhoto(req.UserID, req.CoverPhoto)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCoverPhotoResponse{
		Url: url,
	}, nil
}

//Admin

func (u *AuthServer) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	token, err := u.adminUseCase.AdminLogin(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.AdminLoginResponse{
		AdminToken: token,
	}, nil
}

// Auth Middlewire

func (u *AuthServer) ValidateUserToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	accessToken := req.AccessToken
	refreshToken := req.RefreshToken
	id, err := u.userUseCase.VerifyUserToken(accessToken, refreshToken)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateTokenResponse{
		UserID: id,
	}, nil
}

func (u *AuthServer) ValidateAdminToken(ctx context.Context, req *pb.ValidateAdminTokenRequest) (*emptypb.Empty, error) {
	err := u.adminUseCase.VerifyAdminToken(req.Token)
	return nil, err
}
