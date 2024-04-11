package interface_Repository_Server_Service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
)

type IRepositoryServer interface {
	CreateServer(*requestmodel_server_service.Server) (*responsemodel_server_service.Server, error)
}
