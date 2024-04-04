package di_auth_svc

import (
	clind_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/clind"
	handler_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	router_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/router"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitAuthClind(engin *echo.Echo, config *config.Config) (*middlewire_auth_svc.Middlewire, error) {
	clind, err := clind_auth_svc.InitClind(config.Auth_service_port)
	if err != nil {
		return nil, err
	}

	engin.Use(middleware.Logger())
	middleWire := middlewire_auth_svc.NewAuthMiddlewire(clind)

	UserHandler := handler_auth_svc.NewAuthHandler(clind)
	AdminHandler := handler_auth_svc.NewAdminAuthHandler(clind)

	router_auth_svc.UserRoutes(engin.Group(""), UserHandler, middleWire)
	router_auth_svc.AdminRoutes(engin.Group("admin"), AdminHandler, middleWire)

	return middleWire, nil
}

