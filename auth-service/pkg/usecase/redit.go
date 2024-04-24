package usecasel_auth_server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	interface_repo_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository/interface"
	interface_usecase_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase/interface"
	"github.com/redis/go-redis/v9"
)

type authCache struct {
	userRepo interface_repo_auth_server.IUserRepository
	redisDb  *redis.Client
}

func NewAuthCache(userRepo interface_repo_auth_server.IUserRepository, redisDb *redis.Client) interface_usecase_auth_server.IAuthCache {
	return &authCache{userRepo: userRepo, redisDb: redisDb}
}

func (c *authCache) UpdateUserProfile(userID string) error {
	fmt.Println("==", userID)
	userProfile, err := c.userRepo.GetUserProfile(userID)
	if err != nil {
		return err
	}

	byteProfile, err := c.jsonMarshel(userProfile)
	if err != nil {
		return err
	}

	result := c.redisDb.Set(context.Background(), "user-"+userID, byteProfile, time.Hour)
	if result.Err() != nil {
		return err
	}
	return nil
}

func (c *authCache) jsonMarshel(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
