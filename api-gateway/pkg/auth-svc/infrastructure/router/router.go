package router_auth_svc

import (
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func UserRoutes(engin *echo.Group, userHandler *handler_auth_svc.AuthHanlder) {
	engin.POST("/signup", userHandler.Signup)
}
