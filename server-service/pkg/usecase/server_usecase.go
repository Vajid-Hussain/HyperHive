package usecase_server_service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	interface_Repository_Server_Service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository/interface"
	interface_useCase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase/interface"
)

type ServerUsecase struct {
	repository interface_Repository_Server_Service.IRepositoryServer
}

func NewServerUseCase(repo interface_Repository_Server_Service.IRepositoryServer) interface_useCase_server_service.IServerUseCase {
	return &ServerUsecase{repository: repo}
}

func (r *ServerUsecase) CreateServer(server *requestmodel_server_service.Server) (*responsemodel_server_service.Server, error) {
	serverRes, err := r.repository.CreateServer(server)
	if err != nil {
		return nil, err
	}

	_, err = r.repository.CreateOrUpdateChannelCategory("General", serverRes.ServerID)
	if err != nil {
		return nil, err
	}

	_, err = r.repository.CreateSuperAdmin(requestmodel_server_service.ServerAdmin{UserID: server.UserID, ServerID: serverRes.ServerID, Role: "SuperAdmin"})
	if err != nil {
		return nil, err
	}
	return serverRes, nil
}
