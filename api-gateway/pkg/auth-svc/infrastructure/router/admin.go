package router_auth_svc

import (
	handler_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(engin *echo.Group, AdminHandler *handler_auth_svc.AdminAuthHanlder) {
	engin.POST("/login", AdminHandler.AdminLogin)
}
