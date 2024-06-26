package router_friend_svc

import (
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	handler_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/handler"
	"github.com/labstack/echo/v4"
)

func FriendRoute(engin *echo.Group, friend *handler_friend_svc.FriendSvc, middlewire *middlewire_auth_svc.Middlewire) {
	engin.Use(middlewire.UserAuthMiddlewire)

	engin.POST("/request", friend.FriendRequest)
	engin.GET("/friends", friend.GetFriends)
	engin.GET("/send", friend.GetSendFriendRequest)
	engin.GET("/received", friend.GetReceivedFriendRequest)
	engin.GET("/block", friend.GetBlockFriendRequest)

	engin.PATCH("/restrict", friend.UpdateFriendshipStatus)
	engin.PATCH("/friendrequest", friend.UpdateFriendshipStatus)

	friendMessage := engin.Group("/chat")
	{
		friendMessage.GET("", friend.FriendMessage)
		friendMessage.GET("/message", friend.GetChat)
	}
}
