package interface_Repository_Server_Service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
)

type IRepositoryServer interface {
	CreateServer(*requestmodel_server_service.Server) (*responsemodel_server_service.Server, error)
	CreateOrUpdateChannelCategory(string, string) (*responsemodel_server_service.ChannelCategory, error)
	CreateSuperAdmin(requestmodel_server_service.ServerAdmin) (*responsemodel_server_service.ServerAdmin, error)
	CreateCategory(*requestmodel_server_service.CreateCategory) error
	CreateChannel(*requestmodel_server_service.CreateChannel) error
	JoinInServer(*requestmodel_server_service.JoinToServer) error
	GetServerCategory(string) ([]*responsemodel_server_service.FullServerChannel, error)
	GetChannelUnderCategory(string) ([]*responsemodel_server_service.Channel, error)
	GetServer(string) (*responsemodel_server_service.Server, error)
	UpdateServerCoverPhoto(*requestmodel_server_service.ServerImages) error
	UpdateServerIcon(*requestmodel_server_service.ServerImages) error
	GetUserServers(string) ([]*responsemodel_server_service.UserServerList, error)
	KeepMessageInDB(requestmodel_server_service.ServerMessage) error
	GetServers( string, requestmodel_server_service.Pagination) ( []*responsemodel_server_service.Server,  error)

	GetChannelMessages(string, requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMessage, error)
	UpdateServerDiscription(*requestmodel_server_service.Description) error
	GetServerMembers(string, requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMembers, error)
	ChangeMemberRole(*requestmodel_server_service.UpdateMemberRole) error
	RemoveUserFromServer(*requestmodel_server_service.RemoveUser) error
	LeftFromServer(string, string) error
	DeleteServer(string, string) error

	// forum channel
	InsertForumPost(requestmodel_server_service.ForumPost) error
	InsertForumCommand(requestmodel_server_service.FormCommand) error
	GetForumPost(string, requestmodel_server_service.Pagination) ([]*responsemodel_server_service.ForumPost, error)
	GetForumCommands(string, requestmodel_server_service.Pagination) ([]*responsemodel_server_service.ForumCommand, error)
	GetFormSinglePost(string) (*responsemodel_server_service.ForumPost, error)
}
