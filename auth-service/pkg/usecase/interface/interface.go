package interface_usecase_auth_server

import (
	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
)

type IUserUseCase interface {
	Signup(userDetails requestmodel_auth_server.UserSignup) (*responsemodel_auth_server.UserSignup, error)
	VerifyUserSignup(string, string) error
	ReSendVerificationMail( string) (string, error) 
	ConfirmSignup(string) (*responsemodel_auth_server.AuthenticationResponse, error)
	UserLogin(email, password string) (*responsemodel_auth_server.AuthenticationResponse, error)
	VerifyUserToken(string, string) (string, error) 
	SendOtp( string) (string, error)
	ForgotPassword( requestmodel_auth_server.ForgotPassword) error
	
	//profile
	UpdateProfilePhoto( string,  []byte) ( string,  error) 
	UpdateCoverPhoto(string,  []byte) ( string,  error) 
	UpdateStatusOfUser(requestmodel_auth_server.UserProfileStatus,  float32) error 
	UpdateDescriptionOfUser(userID, description string) error
	GetUserProfile( string) (*responsemodel_auth_server.UserProfile, error) 
	DeleteAccount( string) error 
	DeletePhotoInProfile(string,  string) error
	SerchUsers( string,  requestmodel_auth_server.Pagination) ( *[]responsemodel_auth_server.UserProfile,  error)
}

type IAdminUseCase interface {
	AdminLogin(string, string) (string, error)
	VerifyAdminToken( string) error 
	UnBlockUserAccount( string) (*responsemodel_auth_server.AbstractUserDetails, error) 
	BlockUserAccount( string) (*responsemodel_auth_server.AbstractUserDetails, error)
}
