package router_server_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	handler_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/handler"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

func ServerRouter(engin *echo.Group, handler *handler_server_svc.ServerService, authMiddleWire *middlewire_auth_svc.Middlewire, soketioServer *socketio.Server) {
	engin.Use(authMiddleWire.UserAuthMiddlewire)
	{
		engin.POST("", handler.CreateServer)
		engin.GET("/:id", handler.GetServer)
		engin.POST("/join", handler.JoinToServer)
		engin.GET("/userserver", handler.GetUserServer)

		engin.GET("/", func(ctx echo.Context) error {
			// fmt.Println("===", ctx.Get("userID").(string))
			// ctx.Set("soketIOserver", soketioServer)
			handler.InitSoketio(ctx)
			return echo.WrapHandler(soketioServer)(ctx)
		})

		categoryManagement := engin.Group("/category")
		{
			categoryManagement.POST("", handler.CreateCategory)
			categoryManagement.GET("", handler.GetCategoryOfServer)
		}

		channelManagement := engin.Group("/channel")
		{
			channelManagement.POST("", handler.CreateChannel)
			channelManagement.GET("", handler.GetChannelsOfServer)
		}
	}
}
