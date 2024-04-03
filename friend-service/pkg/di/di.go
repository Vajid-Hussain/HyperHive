package di_friend_server

import (
	config_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/config"
	clind_srv_on_friend_service "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/clind/auth-clind"
	db_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/db"
	server_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/server"
	repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository"
	usecase_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/usecase"
)

func InitFriendService(config *config_friend_server.Config) (*server_friend_server.FriendServer, error) {
	DB, err := db_friend_server.InitDB(&config.DB)
	if err != nil {
		return nil, err
	}

	authClind, err := clind_srv_on_friend_service.InitAuthClind(config.Auth.Auth_Service_port)
	if err != nil {
		return nil, err
	}

	friendRepository := repository_friend_server.NewAdminRepository(DB)
	friendUseCase := usecase_friend_server.NewFriendUseCase(friendRepository, authClind)

	return server_friend_server.NewFriendServer(friendUseCase), nil
}
