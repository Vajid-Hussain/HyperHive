package interface_usecase_auth_server

import (
	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
)

type IUserUseCase interface {
	Signup(userDetails requestmodel_auth_server.UserSignup) (*responsemodel_auth_server.UserSignup, error)
	VerifyUserSignup(string, string) error
	ConfirmSignup(string) (*responsemodel_auth_server.AuthenticationResponse, error)
	UserLogin(email, password string) (*responsemodel_auth_server.AuthenticationResponse, error)
}

type IAdminUseCase interface {
	AdminLogin(string, string) (string, error)
}
