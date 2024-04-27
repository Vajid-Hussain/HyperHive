package di_server_svc

import (
	"fmt"

	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	clind_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/clind"
	handler_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/handler"
	router_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/router"
	usecase_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/useCase"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func InitServerClind(engin *echo.Echo, config *config.Config, middleWire *middlewire_auth_svc.Middlewire) error {
	serverClind, err := clind_server_svc.InitClind(config.Server_service_port)
	if err != nil {
		return err
	}

	authClind, err := clind_server_svc.InitAuthClind(config.Auth_service_port)
	if err != nil {
		return err
	}

	redisDB := InitRedisDB(&config.RedisDB)

	soketioServer := CreateSoketIOserver(config.RedisDB)
	serverUseCase := usecase_server_svc.NewServerServiceUseCase(serverClind.Clind, authClind.Clind, config, redisDB)
	serverHandler := handler_server_svc.NewServerService(serverClind.Clind, soketioServer, serverUseCase)

	router_server_svc.ServerRouter(engin.Group("/server"), serverHandler, middleWire, soketioServer)

	return nil
}

func CreateSoketIOserver(redisDB config.Redis) *socketio.Server {
	rediAdapter := socketio.RedisAdapterOptions{
		Port:     redisDB.RedisPort,
		Password: redisDB.RedisPassword,
		Host:     redisDB.RedisURL,
		// DB: 0,
		Addr: redisDB.RedisURL,
	}
	server := socketio.NewServer(nil)

	ok, err := server.Adapter(&rediAdapter)
	if err != nil {
		fmt.Println("redis adapter connection err ", err)
	}
	fmt.Println("redis adapter connection ", ok)
	go server.Serve()

	return server
}

func InitRedisDB(config *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
		DB: 0,
	})
	return client
}
