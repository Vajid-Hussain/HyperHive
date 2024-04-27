package di_auth_svc

import (
	clind_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/clind"
	handler_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/handler"
	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	router_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/router"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

var clind pb.AuthServiceClient
var envConfig *config.Config

func InitAuthClind(engin *echo.Echo, config *config.Config) (*middlewire_auth_svc.Middlewire, error) {
	var err error
	envConfig = config
	clind, err = clind_auth_svc.InitClind(config.Auth_service_port)
	if err != nil {
		return nil, err
	}

	engin.Use(middleware.Logger())

	middleWire := middlewire_auth_svc.NewAuthMiddlewire(clind)

	authCacing := AuthCache()

	UserHandler := handler_auth_svc.NewAuthHandler(clind, authCacing)
	AdminHandler := handler_auth_svc.NewAdminAuthHandler(clind)

	router_auth_svc.UserRoutes(engin.Group(""), UserHandler, middleWire)
	router_auth_svc.AdminRoutes(engin.Group("admin"), AdminHandler, middleWire)

	return middleWire, nil
}

func AuthCache() *helper_api_gateway.RedisCaching {
	return helper_api_gateway.NewRedisCaching(InitRedisDB(&envConfig.RedisDB), clind)
}

func InitRedisDB(config *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
		// DB:       0,
		// Username: "default",
	})
	return client
}
