package repository_auth_server

import (
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
	interface_repo_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interface_repo_auth_server.IAdminRepository {
	return &AdminRepository{DB: db}
}

func (d *AdminRepository) FetchPaswordUsingEmail(email string) (password string, err error) {
	query := "SELECT password FROM admins WHERE email = $1"
	result := d.DB.Raw(query, email).Scan(&password)
	if result.Error != nil {
		return "", responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", responsemodel_auth_server.ErrNotFound
	}

	return password, nil
}

func (d *AdminRepository) BlockUserAccount(userID string) error {
	query := "UPDATE users SET status = 'block' WHERE status !='delete' "
	result := d.DB.Raw(query, userID)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

func (d *AdminRepository) UnBlockUserAccount(userID string) error {
	query := "UPDATE users SET status = 'active' WHERE status = 'block' "
	result := d.DB.Raw(query, userID)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}
