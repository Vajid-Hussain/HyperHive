package server_auth_server

import (
	"context"
	"fmt"

	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
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

	fmt.Println("====", UserDetails)

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

	return new(emptypb.Empty), nil
}

func (u *AuthServer) SendOtp(ctx context.Context, req *pb.SendOtpRequest) (*pb.SendOtpResponse, error) {
	token, err := u.userUseCase.SendOtp(req.Emain)
	if err != nil {
		return nil, err
	}
	return &pb.SendOtpResponse{Token: token}, nil
}

func (u *AuthServer) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*emptypb.Empty, error) {
	var forgotPasswordReq requestmodel_auth_server.ForgotPassword
	forgotPasswordReq.Password = req.Password
	forgotPasswordReq.Token = req.Token
	forgotPasswordReq.Otp = req.Otp

	err := u.userUseCase.ForgotPassword(forgotPasswordReq)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (u *AuthServer) ReSendVerificationEmail(ctx context.Context, req *pb.ReSendVerificationEmailRequest) (*pb.ReSendVerificationEmailResponse, error) {
	token, err := u.userUseCase.ReSendVerificationMail(req.Token)
	if err != nil {
		return nil, err
	}

	return &pb.ReSendVerificationEmailResponse{
		Token: token,
	}, nil
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

func (u *AuthServer) DeletePhotoInProfile(ctx context.Context, req *pb.DeletePhotoInProfileRequest) (*emptypb.Empty, error) {
	err := u.userUseCase.DeletePhotoInProfile(req.UserID, req.Types)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (u *AuthServer) UpdateUserProfileStatus(ctx context.Context, req *pb.UpdateUserProfileStatusRequest) (*emptypb.Empty, error) {
	var statusReq requestmodel_auth_server.UserProfileStatus
	statusReq.UserID = req.UserID
	statusReq.Status = req.Status

	err := u.userUseCase.UpdateStatusOfUser(statusReq, req.Duration)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (u *AuthServer) UpdateUserProfileDescription(ctx context.Context, req *pb.UpdateUserProfileDescriptionRequest) (*emptypb.Empty, error) {
	fmt.Println("===", req.Description)
	err := u.userUseCase.UpdateDescriptionOfUser(req.UserID, req.Description)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (u *AuthServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*emptypb.Empty, error) {
	err := u.userUseCase.DeleteAccount(req.UserID)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), nil
}

func (u *AuthServer) UserProfile(ctx context.Context, req *pb.UserProfileRequest) (*pb.UserProfileResponse, error) {
	result, err := u.userUseCase.GetUserProfile(req.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.UserProfileResponse{
		UserID:       req.UserID,
		UserName:     result.UserName,
		Name:         result.Name,
		Email:        result.Email,
		ProfilePhoto: result.ProfilePhoto,
		CoverPhoto:   result.CoverPhoto,
		Description:  result.Description,
		Status:       result.Status,
		UserSince:    result.UserSince.String(),
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

func (u *AuthServer) BlockUse(ctx context.Context, req *pb.BlockUseRequest) (*pb.BlockUseResponse, error) {
	result, err := u.adminUseCase.BlockUserAccount(req.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.BlockUseResponse{
		UserID:   req.UserID,
		Name:     result.Name,
		UserName: result.UserName,
	}, nil
}

func (u *AuthServer) UnBlockUser(ctx context.Context, req *pb.UnBlockUserRequest) (*pb.UnBlockUserResponse, error) {
	result, err := u.adminUseCase.UnBlockUserAccount(req.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.UnBlockUserResponse{
		UserID:   req.UserID,
		Name:     result.Name,
		UserName: result.UserName,
	}, nil
}

// Auth Middlewire

func (u *AuthServer) ValidateUserToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	accessToken := req.AccessToken
	// refreshToken := req.RefreshToken
	id, err := u.userUseCase.VerifyUserToken(accessToken)
	if err != nil {
		return nil, err
	}

	return &pb.ValidateTokenResponse{
		UserID: id,
	}, nil
}

func (u *AuthServer) ValidateAdminToken(ctx context.Context, req *pb.ValidateAdminTokenRequest) (*emptypb.Empty, error) {
	err := u.adminUseCase.VerifyAdminToken(req.Token)
	return new(emptypb.Empty), err
}

func (u *AuthServer) SerchUsers(ctx context.Context, req *pb.SerchUsersRequest) (*pb.SerchUsersResponse, error) {
	result, err := u.userUseCase.SerchUsers(req.UserName, requestmodell_auth_server.Pagination{Limit: req.Limit, OffSet: req.Offset})
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.UserProfileResponse
	for _, val := range *result {
		finalResult = append(finalResult, &pb.UserProfileResponse{UserID: val.UserID, UserName: val.UserName, Name: val.Name, ProfilePhoto: val.ProfilePhoto})
	}

	return &pb.SerchUsersResponse{Users: finalResult}, nil
}

func (u *AuthServer) SeperateUserIDFromAccessToken(ctx context.Context, req *pb.SeperateUserIDFromAccessTokenRequest) (*pb.SeperateUserIDFromAccessTokenResponse, error) {
	userID, err := u.userUseCase.SeperateUserIDFromAccessToken(req.AccessToken)
	return &pb.SeperateUserIDFromAccessTokenResponse{UserID: userID}, err
}

func (u *AuthServer) CreateAcceesTokenByValidatingRefreshToken(ctx context.Context, req *pb.CreateAcceesTokenByValidatingRefreshTokenRequest) (*pb.CreateAcceesTokenByValidatingRefreshTokenResponse, error) {
	accesstoken, err := u.userUseCase.CreateAcceesTokenByValidatingRefreshToken(req.RefreshToken)
	return &pb.CreateAcceesTokenByValidatingRefreshTokenResponse{AccessToken: accesstoken}, err
}
