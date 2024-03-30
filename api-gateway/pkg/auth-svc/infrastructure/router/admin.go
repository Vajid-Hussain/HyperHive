package router_auth_svc

import (
	handler_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(engin *echo.Group, AdminHandler *handler_auth_svc.AdminAuthHanlder, middleWire *middlewire_auth_svc.Middlewire) {
	engin.POST("/login", AdminHandler.AdminLogin)

	engin.Use(middleWire.AdminAuthMiddlewire)
	{
		UserManagement := engin.Group("/user")
		{
			UserManagement.PATCH("/block", AdminHandler.BlockUserAccount)
			UserManagement.PATCH("/unblock", AdminHandler.UnBlockUserAccount)
		}
	}
}
