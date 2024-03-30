package router_auth_svc

import (
	handler_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/labstack/echo/v4"
)

func UserRoutes(engin *echo.Group, userHandler *handler_auth_svc.AuthHanlder, middlewire *middlewire_auth_svc.Middlewire) {
	engin.POST("/signup", userHandler.Signup)
	engin.GET("/verify", userHandler.MailVerificationCallback)
	engin.POST("/confirm", userHandler.ConfirmSignup)
	engin.POST("/login", userHandler.UserLogin)

	engin.Use(middlewire.UserAuthMiddlewire)
	{
		engin.POST("/profilephoto", userHandler.UpdateProfilePhoto)
		engin.POST("/coverphoto", userHandler.UpdateCoverPhoto)
		engin.PATCH("/profilephoto", userHandler.UpdateProfilePhoto)
		engin.PATCH("/coverphoto", userHandler.UpdateCoverPhoto)
	}
}
