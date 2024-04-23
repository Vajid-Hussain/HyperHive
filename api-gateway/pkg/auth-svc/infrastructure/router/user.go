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
	engin.POST("/verifyemailresend", userHandler.ReSendVerificationEmail)
	engin.POST("/sendotp", userHandler.SendOtp)
	engin.POST("/forgotpassword", userHandler.ForgotPassword)

	token := engin.Group("token")
	{
		token.GET("", userHandler.SeperateUserIDFromAccessToken)
		token.POST("/accesstoken", userHandler.CreateAcceesTokenByValidatingRefreshToken)
	}

	engin.Use(middlewire.UserAuthMiddlewire)
	{
		account := engin.Group("/account")
		{
			account.DELETE("", userHandler.DeleteUserAcoount)
		}

		profile := engin.Group("/profile")
		{
			profile.GET("/", userHandler.GetUserProfile)
			profile.GET("/:userID", userHandler.GetUserProfile)

			profile.POST("/profilephoto", userHandler.UpdateProfilePhoto)
			profile.POST("/coverphoto", userHandler.UpdateCoverPhoto)
			profile.PATCH("/profilephoto", userHandler.UpdateProfilePhoto)
			profile.PATCH("/coverphoto", userHandler.UpdateCoverPhoto)

			profile.DELETE("/photoprofile", userHandler.DeletePhotFromUserProfile)

			profile.POST("/status", userHandler.UpdateProfileStatus)
			profile.POST("/description", userHandler.UpdateProfileDescription)
			profile.PATCH("/status", userHandler.UpdateProfileStatus)
			profile.PATCH("/description", userHandler.UpdateProfileDescription)
		}

		users := engin.Group("users")
		{
			users.GET("", userHandler.SerchUsers)
		}
	}
}
