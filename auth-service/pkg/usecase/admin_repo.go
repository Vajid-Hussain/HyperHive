package usecasel_auth_server

import (
	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	interface_repo_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository/interface"
	interface_usecase_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase/interface"
	utils_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/utils"
)

type AdminUseCase struct {
	AdminRepo interface_repo_auth_server.IAdminRepository
	token     configl_auth_server.Token
}

func NewAdminUseCase(repo interface_repo_auth_server.IAdminRepository, token configl_auth_server.Token) interface_usecase_auth_server.IAdminUseCase {
	return &AdminUseCase{token: token, AdminRepo: repo}
}

func (d *AdminUseCase) AdminLogin(email, password string) (token string, err error) {
	storedPassword, err := d.AdminRepo.FetchPaswordUsingEmail(email)
	if err != nil {
		return "", err
	}

	err = utils_auth_server.CompairPassword(storedPassword, password)
	if err != nil {
		return "", err
	}

	token, err = utils_auth_server.GenerateRefreshToken(d.token.AdminSecurityKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
