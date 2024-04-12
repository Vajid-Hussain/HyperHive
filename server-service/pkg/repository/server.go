package repository_server_service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	interface_Repository_Server_Service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type ServerRepository struct {
	DB *gorm.DB
}

func NewServerRepository(db *gorm.DB) interface_Repository_Server_Service.IRepositoryServer {
	return &ServerRepository{DB: db}
}

func (d *ServerRepository) CreateServer(server *requestmodel_server_service.Server) (*responsemodel_server_service.Server, error) {
	var res responsemodel_server_service.Server

	query := "INSERT INTO servers (name) SELECT $1 WHERE NOT EXISTS (SELECT 1 FROM servers WHERE name= $1) RETURNING *"
	result := d.DB.Raw(query, server.Name).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrNotUniqueServerName
	}
	return &res, nil
}

func (d *ServerRepository) CreateOrUpdateChannelCategory(name string, serverID string) (res *responsemodel_server_service.ChannelCategory, err error) {
	query := "INSERT INTO channel_categories (server_id, name) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM channel_categories WHERE server_id= $1 AND name=$2) RETURNING *"
	result := d.DB.Raw(query, serverID, name).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return res, nil
}

func (d *ServerRepository) CreateSuperAdmin(admin requestmodel_server_service.ServerAdmin) (*responsemodel_server_service.ServerAdmin, error) {
	var res responsemodel_server_service.ServerAdmin
	query := "INSERT INTO server_moderators (server_id , user_id, role) VALUES ($1, $2, $3)  RETURNING *"
	result := d.DB.Raw(query, admin.ServerID, admin.UserID, admin.Role).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return &res, nil
}

func (d *ServerRepository) CreateCategory(req *requestmodel_server_service.CreateCategory) error {
	query := "INSERT INTO channel_categories (server_id, name) SELECT $1, $2 WHERE NOT EXIST (SELECT 1 FROM channel_categories WHERE server_id=$1 AND name =$2) AND EXIST(SELECT 1 FROM server_moderators WHERE server_id= $1, AND user_id= $3 AND role= 'SuperAdmin')"
	result := d.DB.Raw(query, req.ServerID, req.CategoryName, req.UserID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrEmptyResponse
	}
	return nil
}

// func (d *ServerRepository) CreateChannel(req *requestmodel_server_service.CreateChannel) error {
// 	query := "INSET INTO "
// }
