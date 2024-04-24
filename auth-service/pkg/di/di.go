package di_auth_server

import (
	// "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	cron_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/cronJob"
	db_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/db"
	server_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/server"
	repositoryl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository"
	usecase_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase"
	usecasel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase"
	"github.com/redis/go-redis/v9"
)

func InitAuthServer(config *configl_auth_server.Config) (*server_auth_server.AuthServer, error) {
	DB, err := db_auth_server.InitDB(&config.DB)
	if err != nil {
		return nil, err
	}

	redisDB := InitRedisDB(&config.RedisDB)

	userRepository := repositoryl_auth_server.NewUserRepository(DB)
	redisCache := usecasel_auth_server.NewAuthCache(userRepository, redisDB)
	userUseCase := usecase_auth_server.NewUserUseCase(userRepository, config.S3, config.Mail, config.Token, redisCache)

	adminRepository := repositoryl_auth_server.NewAdminRepository(DB)
	adminUseCase := usecase_auth_server.NewAdminUseCase(adminRepository, config.Token)

	crons := cron_auth_server.NewCronJob(userRepository)
	crons.StartCronInAuthService()

	return server_auth_server.NewAuthServer(userUseCase, adminUseCase), nil
}

func InitRedisDB(config *configl_auth_server.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURL,
		Password: config.RedisPassword,
	})
	return client
}
