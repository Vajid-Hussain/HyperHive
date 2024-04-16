package di_server_service

import (
	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	clind_srv_on_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/clind"
	db_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/db"
	server_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/server"
	repository_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository"
	usecase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase"
)

func ServerInitialize(config *config_server_service.Config) (*server_server_service.ServerServer, error) {
	gormDB, mongoCollection, err := db_server_service.DbInit(config.DB, config.MongoDB)
	if err != nil {
		return nil, err
	}

	authClind, err := clind_srv_on_server_service.InitAuthClind(config.Auth.Auth_Service_port)
	if err != nil {
		return nil, err
	}

	serverRepository := repository_server_service.NewServerRepository(gormDB, mongoCollection)
	serverUseCase := usecase_server_service.NewServerUseCase(serverRepository, config.KafkaConsumer, authClind)
	serverServer := server_server_service.NewServerServer(serverUseCase)

	go serverUseCase.KafkaConsumerServerMessage()

	return serverServer, nil
}
