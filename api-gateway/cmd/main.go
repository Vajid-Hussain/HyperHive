package main

import (
	"log"

	di_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/di"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	di_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/di"
	di_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/di"
	"github.com/labstack/echo/v4"
)

func main() {
	serverError := "error at initial setup "
	
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	engin := echo.New()

	middlewire, err := di_auth_svc.InitAuthClind(engin, config)
	if err != nil {
		log.Fatal(serverError, err)
	}

	err = di_friend_svc.InitFriendClind(engin, config, middlewire)
	if err != nil {
		log.Fatal(serverError, err)
	}

	err = di_server_svc.InitServerClind(engin, config, middlewire)
	if err != nil {
		log.Fatal(serverError, err)
	}

	engin.Logger.Fatal(engin.Start(config.PORT))
}
