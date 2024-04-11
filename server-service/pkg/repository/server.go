package repository_server_service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	"gorm.io/gorm"
)

type ServerRepository struct {
	DB *gorm.DB
}

func NewServerRepository(db *gorm.DB) {
	return &ServerRepository{DB: db}
}

func (d *ServerRepository) CreateServer(server *requestmodel_server_service.Server) (res *responsemodel_server_service.Server, err error) {
	query := "INSERT INTO server (name) VALUES ($1)"
	result := d.DB.Raw(query, server.Name).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return res, nil
}
