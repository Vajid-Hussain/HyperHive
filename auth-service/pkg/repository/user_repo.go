package repository_auth_server

import (
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

	query := "INSERT INTO users (name, user_name, email, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING *"
	result := d.DB.Raw(query, userReq.Name, userReq.UserName, userReq.Email, userReq.Password, time.Now()).Scan(&userRes)
	if result.Error != nil {
		d.DB.Rollback()
		return nil, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		d.DB.Rollback()
		return nil, responsemodel_auth_server.ErrDBNoRowAffected
	}

	d.DB.Commit()
	return userRes, nil
}

func (d *UserRepository) UserNameIsExist(userName string) (count int, err error) {
	query := "SELECT count(*) FROM users WHERE user_name = $1"
	result := d.DB.Raw(query, userName).Scan(&count)
	if result.Error != nil {
		return 0, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return 0, responsemodel_auth_server.ErrDBNoRowAffected
	}
	return count, nil
}

func (d *UserRepository) EmailIsExist(email string) (count int, err error) {
	query := "SELECT count(*) FROM users WHERE email = $1"
	result := d.DB.Raw(query, email).Scan(&count)
	if result.Error != nil {
		return 0, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return 0, responsemodel_auth_server.ErrDBNoRowAffected
	}
	return count, nil
}

func (d *UserRepository) VerifyUserSignup(userID, email string) error {
	query := "UPDATE users SET status = 'active' WHERE id = $1 AND email = $2"
	result := d.DB.Exec(query, userID, email)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

func (d *UserRepository) ConfirmSignup(userID string) (count int, err error) {
	query := "SELECT count(*) FROM users WHERE id= $1"
	result := d.DB.Raw(query, userID).Scan(&count)
	if result.Error != nil {
		return 0, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return 0, responsemodel_auth_server.ErrNotFound
	}
	return count, nil
}
