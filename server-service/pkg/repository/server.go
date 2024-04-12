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
	query := "INSERT INTO channel_category (server_id, name) VALUES($1, $2) ON CONFLICT (server_id, name) DO UPDATE SET name=$2 RETURNING *"
	result := d.DB.Raw(query, serverID, name).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return res, nil
}

func (d *ServerRepository) CreateSuperAdmin(admin requestmodel_server_service.ServerAdmin) (res *responsemodel_server_service.ServerAdmin, err error) {
	query := "INSERT INTO server_admin (server_id , user_id, role) VALUES ($1, $2, $3) ON CONFLICT (server_id, user_id) DO UPDATE SET server_id=$1, user_id=$2 "
	result := d.DB.Raw(query, admin.ServerID, admin.UserID, admin.Role).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return res, nil
}
