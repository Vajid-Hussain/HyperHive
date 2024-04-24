package helper_api_gateway

import (
	"context"
	"encoding/json"
	"time"

	response_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/Model/response"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	"github.com/redis/go-redis/v9"
)

type RedisCaching struct {
	redis     *redis.Client
	authClind pb.AuthServiceClient
}

func NewRedisCaching(redis *redis.Client, authClind pb.AuthServiceClient) *RedisCaching {
	return &RedisCaching{
		redis:     redis,
		authClind: authClind,
	}
}

func (r *RedisCaching) GetUserProfile(userID string) (*response_auth_svc.UserProfile, error) {
	var UserProfieFinalResult = &response_auth_svc.UserProfile{}
	var err error

	userProfile := r.redis.Get(context.Background(), "user-"+userID)
	if userProfile.Val() == "" {
		UserProfieFinalResult, err = r.SetUserProfile(userID)
		if err != nil {
			return nil, err
		}
	} else {
		err = r.jsonUnmarshel(UserProfieFinalResult, []byte(userProfile.Val()))
		if err != nil {
			return nil, err
		}
	}

	return UserProfieFinalResult, nil
}

func (r *RedisCaching) SetUserProfile(userID string) (*response_auth_svc.UserProfile, error) {

	contextTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userProfileFromService, err := r.authClind.UserProfile(contextTimeout, &pb.UserProfileRequest{UserID: userID})
	if err != nil {
		return nil, err
	}

	profileByte, err := r.structMarshel(userProfileFromService)
	if err != nil {
		return nil, err
	}

	result := r.redis.Set(context.Background(), "user-"+userID, profileByte, time.Hour)
	if result.Err() != nil {
		return nil, err
	}

	return &response_auth_svc.UserProfile{
		UserID:       userProfileFromService.UserID,
		Name:         userProfileFromService.Name,
		UserName:     userProfileFromService.UserName,
		Email:        userProfileFromService.Email,
		ProfilePhoto: userProfileFromService.ProfilePhoto,
		CoverPhoto:   userProfileFromService.CoverPhoto,
		Description:  userProfileFromService.Description,
		Status:       userProfileFromService.Status,
	}, nil
}

func (r *RedisCaching) structMarshel(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (r *RedisCaching) jsonUnmarshel(model interface{}, data []byte) error {
	return json.Unmarshal(data, model)
}
