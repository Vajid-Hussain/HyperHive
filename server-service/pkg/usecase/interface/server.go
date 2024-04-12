package interface_useCase_server_service

import (
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
)

type IServerUseCase interface {
	CreateServer(*requestmodel_server_service.Server) (*responsemodel_server_service.Server, error)
}
