package di_friend_server

import (
	"fmt"

	config_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/config"
	clind_srv_on_friend_service "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/clind/auth-clind"
	db_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/db"
	server_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/server"
	repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository"
	interface_repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository/interface"
	usecase_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/usecase"
	"github.com/robfig/cron"
)

func InitFriendService(config *config_friend_server.Config) (*server_friend_server.FriendServer, error) {
	DB, mongoCollection, err := db_friend_server.InitDB(&config.DB, &config.Mongo)
	if err != nil {
		return nil, err
	}

	authClind, err := clind_srv_on_friend_service.InitAuthClind(config.Auth.Auth_Service_port)
	if err != nil {
		return nil, err
	}

	friendRepository := repository_friend_server.NewFriendRepository(DB, mongoCollection)
	friendUseCase := usecase_friend_server.NewFriendUseCase(friendRepository, authClind, config.Kafka)
	friendServer := server_friend_server.NewFriendServer(friendUseCase)

	cronJobDeleteRevokeFriendShip(friendRepository)

	go friendUseCase.MessageConsumer()

	return friendServer, nil
}

func cronJobDeleteRevokeFriendShip(repo interface_repository_friend_server.IFriendRepository) {
	c:= cron.New()
	c.AddFunc("0 0 * * *" ,func ()  {
		err := repo.DeleteFriendShipOfStatusRejectRevoke()
		if err != nil {
			fmt.Println("error at cron job", err)
		}
	})
	c.Start()
}
