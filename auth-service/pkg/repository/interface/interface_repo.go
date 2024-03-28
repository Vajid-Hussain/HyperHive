package interface_repo_auth_server

import (
	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
)

type IUserRepository interface {
	Signup(requestmodel_auth_server.UserSignup) (*responsemodel_auth_server.UserSignup, error)
}
