package interface_repo_auth_server

import (
	"time"

	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
)

type IUserRepository interface {
	Signup(requestmodel_auth_server.UserSignup) (*responsemodel_auth_server.UserSignup, error)
	VerifyUserSignup(string, string) error
	ConfirmSignup(string) (int, error)
	IsUserIDExist(string) (int, error)
	EmailIsExist(string) (int, error)
	FetchMailUsingUserID(string) (string, error)
	UserNameIsExist(string) (int, error)
	GetUserPasswordUsingEmail(string) (string, error)
	FetchUserIDUsingMail(string) (string, error)
	CreateOtp(string, string, time.Time) error
	FetchOtp(string,time.Time) (string, error)
	ForgotPassword(string, string) error
	DeleteUnverifiedUsers() 

	//profile
	UpdateUserProfilePhoto(string, string) error
	UpdateCoverPhoto(string, string) error
	UpdateOrCreateUserStatus(requestmodel_auth_server.UserProfileStatus) error
	UpdateOrCreateUserDescription(string, string) error
	GetUserProfile(string) (*responsemodel_auth_server.UserProfile, error)
	DeleteUserAcoount(string) error
	DeleteExpiredStatus( time.Time)
	DeleteProfilePhoto( string) error
	DeleteCoverPhoto( string) error
}

type IAdminRepository interface {
	FetchPaswordUsingEmail(string) (string, error)
	UnBlockUserAccount(string) (*responsemodel_auth_server.AbstractUserDetails, error)
	BlockUserAccount(userID string) (*responsemodel_auth_server.AbstractUserDetails, error)
}
