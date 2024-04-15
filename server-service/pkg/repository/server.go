package repository_server_service

import (
	"fmt"

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

	query := "INSERT INTO servers (name) SELECT $1 WHERE NOT EXISTS (SELECT 1 FROM servers WHERE name= $1 AND status='active') RETURNING *"
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
	query := "INSERT INTO channel_categories (server_id, name) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM channel_categories WHERE server_id= $1 AND name=$2  AND status='active') RETURNING *"
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
	query := "INSERT INTO server_members (server_id , user_id, role) VALUES ($1, $2, $3)  RETURNING *"
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
	query := "INSERT INTO channel_categories (server_id, name) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM channel_categories WHERE server_id=$1 AND name =$2 AND status='active') AND EXISTS(SELECT 1 FROM server_members WHERE server_id= $1 AND user_id= $3 AND role= 'SuperAdmin' AND status='active')"
	result := d.DB.Exec(query, req.ServerID, req.CategoryName, req.UserID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrcategoryExistOrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) CreateChannel(req *requestmodel_server_service.CreateChannel) error {
	query := "INSERT INTO channels (server_id, categoryid, name, type) SELECT $1, $2, $3, $4 WHERE NOT EXISTS (SELECT 1 FROM channels WHERE server_id= $1 AND name= $3  AND status='active') AND EXISTS(SELECT 1 FROM server_members WHERE server_id= $1 AND user_id= $5 AND role= 'SuperAdmin' AND status='active')"
	result := d.DB.Exec(query, req.ServerID, req.CategoryID, req.ChannelName, req.Type, req.UserID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrChannelExistOrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) JoinInServer(req *requestmodel_server_service.JoinToServer) error {
	query := "INSERT INTO server_members (server_id, user_id, role) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM server_members WHERE server_id= $1 AND user_id= $2  AND status='active')"
	result := d.DB.Exec(query, req.ServerID, req.UserID, req.Role)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrExistMemberJoin
	}
	return nil
}

func (d *ServerRepository) GetServerCategory(serverID string) ([]*responsemodel_server_service.FullServerChannel, error) {
	var res []*responsemodel_server_service.FullServerChannel
	query := "SELECT category_id, name FROM channel_categories WHERE server_id= $1 AND status='active'"
	result := d.DB.Raw(query, serverID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}
	return res, nil
}

func (d *ServerRepository) GetChannelUnderCategory(categoryID string) ([]*responsemodel_server_service.Channel, error) {
	var res []*responsemodel_server_service.Channel
	query := "SELECT * FROM channels WHERE categoryid= $1 AND status='active'"
	result := d.DB.Raw(query, categoryID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}
	return res, nil
}

func (d *ServerRepository) GetServer(serverID string) (*responsemodel_server_service.Server, error) {
	var res *responsemodel_server_service.Server
	query := "SELECT * FROM servers WHERE id=$1 AND status= 'active'"
	result := d.DB.Raw(query, serverID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}
	return res, nil
}

func (d *ServerRepository) UpdateServerCoverPhoto(serverID, url string) error {
	query := "UPDATE servers SET cover_photo=$1 WHERE id=$2 AND status='active' "
	result := d.DB.Exec(query, serverID, url)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrEmptyResponse
	}
	return nil
}

func (d *ServerRepository) UpdateServerIcon(serverID, url string) error {
	query := "UPDATE servers SET icon=$1 WHERE id=$2 AND status='active' "
	result := d.DB.Exec(query, serverID, url)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrEmptyResponse
	}
	return nil
}

func (d *ServerRepository) GetUserServers(userID string) ([]*responsemodel_server_service.UserServerList, error) {
	fmt.Println(userID)
	var res []*responsemodel_server_service.UserServerList
	query := "SELECT * FROM servers INNER JOIN server_members ON servers.id= server_members.server_id WHERE user_id= $1 AND server_members.status='active' AND servers.status='active'"
	result := d.DB.Raw(query, userID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}
	return res, nil
}

// func (d *ServerRepository) UpdateServerCoverPhoto(url string) error{
// 	query:= "UPDATE servers SET cover_photo= $1"
// 	result:=d.DB.Exec(query,url)
// 	if result.Error != nil {
// 		return responsemodel_server_service.ErrInternalServer
// 	}

// 	if result.RowsAffected == 0 {
// 		return responsemodel_server_service.ErrEmptyResponse
// 	}
// 	return nil
// }

// func (d *ServerRepository) UpdateServerIcon(url string) error{
// 	query:= "UPDATE servers SET icon= $1"
// 	result:=d.DB.Exec(query,url)
// 	if result.Error != nil {
// 		return responsemodel_server_service.ErrInternalServer
// 	}

// 	if result.RowsAffected == 0 {
// 		return responsemodel_server_service.ErrEmptyResponse
// 	}
// 	return nil
// }
