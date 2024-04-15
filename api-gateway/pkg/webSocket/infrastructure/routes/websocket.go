package routes_websocket_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	handler_werbsocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func WebsocketRoutes(engin echo.Group, handler *handler_werbsocket_svc.WebsocketHandler, middlewire *middlewire_auth_svc.Middlewire) {
	engin.Use(middlewire.UserAuthMiddlewire)
	{
		engin.GET("", handler.WebSocketConnection)
	}
}
