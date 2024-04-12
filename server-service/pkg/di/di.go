package di_server_service

import (
	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	db_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/db"
	server_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/server"
	repository_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository"
	usecase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase"
)

func ServerInitialize(config *config_server_service.Config) (*server_server_service.ServerServer, error) {
	gormDB, err := db_server_service.DbInit(config.DB)
	if err != nil {
		return nil, err
	}

	serverRepository := repository_server_service.NewServerRepository(gormDB)
	serverUseCase := usecase_server_service.NewServerUseCase(serverRepository)
	serverServer := server_server_service.NewServerServer(serverUseCase)

	return serverServer, nil
}
