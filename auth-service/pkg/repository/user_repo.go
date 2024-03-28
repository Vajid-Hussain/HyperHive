package repository_auth_server

import (
	"fmt"
	"time"

	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
	interfacel_repo_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfacel_repo_auth_server.IUserRepository {
	return &UserRepository{DB: db}
}

func (d *UserRepository) Signup(userReq requestmodel_auth_server.UserSignup) (userRes *responsemodel_auth_server.UserSignup, err error) {
	d.DB.Begin()
	defer d.DB.Rollback()

	query := "INSERT INTO users (name, user_name, email, password, profile_photo_url, cover_photo_url, created_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	result := d.DB.Raw(query, userReq.Name, userReq.UserName, userReq.Email, userReq.Password, userReq.ProfilePhotoUrl, userReq.CoverPhotoUrl, time.Now()).Scan(&userRes)
	if result.Error != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, err
	}

	fmt.Println("===", userRes)
	d.DB.Commit()
	return userRes, nil
}
