package interface_repo_auth_server

import (
	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
)

type IUserRepository interface {
	Signup(requestmodel_auth_server.UserSignup) (*responsemodel_auth_server.UserSignup, error)
	VerifyUserSignup(string, string) error
	ConfirmSignup(string) (int, error)
	EmailIsExist(string) (int, error)
	UserNameIsExist(string) (int, error)
	GetUserPasswordUsingEmail(string) (string, error)
	FetchUserIDUsingMail(string) (string, error)

	//profile
	UpdateUserProfilePhoto(string, string) error
	UpdateCoverPhoto(string,  string) error 
}

type IAdminRepository interface {
	FetchPaswordUsingEmail(string) (string, error)
}
