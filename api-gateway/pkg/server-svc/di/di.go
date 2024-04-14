package di_server_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	clind_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/clind"
	handler_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/handler"
	router_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/router"
	"github.com/labstack/echo/v4"
)

func InitServerClind(engin *echo.Echo, config *config.Config, middleWire *middlewire_auth_svc.Middlewire) error {
	serverClind, err := clind_server_svc.InitClind(config.Server_service_port)
	if err != nil {
		return err
	}

	serverHandler := handler_server_svc.NewServerService(serverClind.Clind)
	// engin.GET("/socket.io/", serverHandler.SoketIO)

	router_server_svc.ServerRouter(engin.Group("/server"), serverHandler, middleWire)

	return nil
}
