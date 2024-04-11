package usecase_server_service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	interface_Repository_Server_Service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository/interface"
)

type ServerUsecase struct {
	useCase interface_Repository_Server_Service.IRepositoryServer
}

func NewServerUseCase(repo interface_Repository_Server_Service.IRepositoryServer) {
	return &ServerUsecase{useCase: repo}
}

func (r *ServerUsecase) CreateServer(server *requestmodel_server_service.Server) (*respose)