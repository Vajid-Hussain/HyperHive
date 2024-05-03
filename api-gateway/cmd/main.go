package main

import (
	"log"

	"github.com/Vajid-Hussain/HiperHive/api-gateway/docs"
	di_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/di"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	di_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/di"
	di_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/di"
	di_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title          HyperHive
// @version        1.0
// @description    This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @host     hyperhive.vajid.tech
// @BasePath /

// @securityDefinitions.apikey UserAuthorization
// @in                         header
// @name                       AccessToken
// @securityDefinitions.apikey UserConfirmToken
// @in                         header
// @name                       ConfirmToken
// @securityDefinitions.apikey AdminAutherisation
// @in                         header
// @name                       AccessToken

func main() {
	docs.SwaggerInfo.Host = "dev.hyperhive.vajid.tech"
	serverError := "error at initial setup "

	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	engin := echo.New()

	// engin.Use(middleware.CORS())

	engin.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	engin.GET("/swagger/*", echoSwagger.WrapHandler)

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

	di_websocket_svc.InitWebSocket(engin, config, middlewire)
	if err != nil {
		log.Fatal(serverError, err)
	}

	engin.Logger.Fatal(engin.Start(config.PORT))
}
