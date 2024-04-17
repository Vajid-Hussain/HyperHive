package interface_useCase_server_service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
)

type IServerUseCase interface {
	CreateServer(*requestmodel_server_service.Server) (*responsemodel_server_service.Server, error)
	CreateChannel(*requestmodel_server_service.CreateChannel) error
	CreateCategory(*requestmodel_server_service.CreateCategory) error
	JoinToServer(*requestmodel_server_service.JoinToServer) error
	GetServerCategory(string) ([]*responsemodel_server_service.FullServerChannel, error)
	GetChannels(string) ([]*responsemodel_server_service.FullServerChannel, error)
	GetUserServers(string) ([]*responsemodel_server_service.UserServerList, error)
	GetServer(string) (*responsemodel_server_service.Server, error)
	KafkaConsumerServerMessage() error
	GetChannelMessages(string, requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMessage, error)
	UpdateServerPhoto(*requestmodel_server_service.ServerImages) error
	UpdateServerDiscription(*requestmodel_server_service.Description) error
	GetServerMembers(string, requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMembers, error)
	UpdateMemberRole(requestmodel_server_service.UpdateMemberRole) error
	RemoveUserFromServer(*requestmodel_server_service.RemoveUser) error
	LeftFromServer(string,  string) error
	DeleteServer(string,  string) error 
}
