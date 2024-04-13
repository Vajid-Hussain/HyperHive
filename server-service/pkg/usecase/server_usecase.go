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

func (r *ServerUsecase) CreateCategory(req *requestmodel_server_service.CreateCategory) error {
	return r.repository.CreateCategory(req)
}

func (r *ServerUsecase) CreateChannel(req *requestmodel_server_service.CreateChannel) error {
	if req.Type != "voice" && req.Type != "text" && req.Type != "forum" {
		return responsemodel_server_service.ErrChannelTypeIsNotMatch
	}

	return r.repository.CreateChannel(req)
}

func (r *ServerUsecase) JoinToServer(req *requestmodel_server_service.JoinToServer) error {
	req.Role = "member"
	return r.repository.JoinInServer(req)
}

func (r *ServerUsecase) GetServerCategory(serverID string) ([]*responsemodel_server_service.FullServerChannel, error) {
	return r.repository.GetServerCategory(serverID)
}

func (r *ServerUsecase) GetChannels(serverID string) ([]*responsemodel_server_service.FullServerChannel, error) {
	category, err := r.repository.GetServerCategory(serverID)
	if err != nil {
		return nil, err
	}

	for i, val := range category {
		category[i].Channel, _ = r.repository.GetChannelUnderCategory(val.CategoryID)
	}
	return category, nil
}

func (r *ServerUsecase) GetUserServers(userID string) ([]*responsemodel_server_service.UserServerList, error) {
	return r.repository.GetUserServers(userID)
}

func (r *ServerUsecase) GetServer(serverID string) (*responsemodel_server_service.Server, error){
	return r.repository.GetServer(serverID)
}
