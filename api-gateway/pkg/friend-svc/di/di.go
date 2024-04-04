package di_friend_svc

import (
	"fmt"

	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	clind_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/clind"
	handler_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/handler"
	router_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/router"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func InitFriendClind(engin *echo.Echo, config *config.Config, middlewire *middlewire_auth_svc.Middlewire) error {
	clind, err := clind_friend_svc.InitClind(config.Friend_service_Port)
	if err != nil {
		return err
	}

	redisDB := InitRedisDB(&config.RedisDB)
	fmt.Println("db connected ==", redisDB)

	redis:= handler_friend_svc.NewCaching(redisDB)

	handler := handler_friend_svc.NewFriendSvc(clind.Clind, redis)

	router_friend_svc.FriendRoute(engin.Group("friend"), handler, middlewire)

	return nil
}

func InitRedisDB(config *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
	})
	return client
}
