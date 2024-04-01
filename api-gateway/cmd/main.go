package main

import (
	"log"

	di_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/di"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	di_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/di"
	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	engin := echo.New()

	middlewire, err := di_auth_svc.InitAuthClind(engin, config)
	if err != nil {
		log.Fatal("error at initial setup", err)
	}

	err = di_friend_svc.InitFriendClind(engin, config, middlewire)
	if err != nil {
		log.Fatal("error at initial setup", err)
	}

	engin.Logger.Fatal(engin.Start(config.PORT))
}
