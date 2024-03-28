package server_auth_server

import (
	"context"

	requestmodell_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	"github.com/Vajid-Hussain/HiperHive/auth-service/pkg/pb"
	interface_usecase_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase/interface"
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
	// UserDetails.ProfilePhoto = req.ProfilePhoto

	userReq, err := u.userUseCase.Signup(UserDetails)
	if err != nil {
		return nil, err
	}

	return &pb.SignupResponse{
		UserID:          userReq.ID,
		UserName:        userReq.UserName,
		Name:            userReq.Name,
		Email:           userReq.Email,
		// ProfilePhotoUrl: userReq.ProfilePhotoUrl,
		// CoverPhotoUrl:   userReq.CoverPhotoUrl,
	}, nil
}
