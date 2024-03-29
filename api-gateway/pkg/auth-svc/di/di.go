package di_auth_svc

import (
	clind_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/clind"
	handler_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	router_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/router"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitAuthClind(engin *echo.Echo, config *config.Config) error {
	clind, err := clind_auth_svc.InitClind(config.Auth_service_port)
	if err != nil {
		return err
	}

	engin.Use(middleware.Logger())

	UserHandler := handler_auth_svc.NewAuthHandler(clind)
	AdminHandler:= handler_auth_svc.NewAdminAuthHandler(clind)

	router_auth_svc.UserRoutes(engin.Group("user"), UserHandler)
	router_auth_svc.AdminRoutes(engin.Group("admin"), AdminHandler)

	return nil
}
