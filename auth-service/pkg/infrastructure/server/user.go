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
	userUseCase interface_usecase_auth_server.IUserUseCase
}

func NewAuthServer(userUseCase interface_usecase_auth_server.IUserUseCase) *AuthServer {
	return &AuthServer{userUseCase: userUseCase}
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
		UserID:   userReq.ID,
		UserName: userReq.UserName,
		Name:     userReq.Name,
		Email:    userReq.Email,
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
