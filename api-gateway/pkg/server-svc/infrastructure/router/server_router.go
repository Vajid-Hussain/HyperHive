package router_server_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	handler_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func ServerRouter(engin *echo.Group, handler *handler_server_svc.ServerService, authMiddleWire *middlewire_auth_svc.Middlewire) {
	engin.Use(authMiddleWire.UserAuthMiddlewire)
	{
		engin.POST("", handler.CreateServer)
	}
}
