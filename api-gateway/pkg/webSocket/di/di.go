package di_websocket_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	handler_werbsocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/handler"
	routes_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/routes"
	usecase_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/useCase"
	"github.com/labstack/echo/v4"
)

func InitWebSocket(engin *echo.Echo, config *config.Config, middlewire *middlewire_auth_svc.Middlewire) {
	usecase := usecase_websocket_svc.NewWebSocketUseCase(*config)
	handler := handler_werbsocket_svc.NewWebSocketHandler(usecase)
	routes_websocket_svc.WebsocketRoutes(*engin.Group("/websocket"), handler, middlewire)
}
