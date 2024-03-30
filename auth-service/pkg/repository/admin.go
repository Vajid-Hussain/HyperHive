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
	query := "SELECT password FROM admins WHERE email = $1 RETURNING *"
	result := d.DB.Raw(query, email).Scan(&password)
	if result.Error != nil {
		return "", responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", responsemodel_auth_server.ErrNotFound
	}

	return password, nil
}

func (d *AdminRepository) BlockUserAccount(userID string) (*responsemodel_auth_server.AbstractUserDetails, error) {
	var user responsemodel_auth_server.AbstractUserDetails

	query := "UPDATE users SET status = 'block' WHERE status !='delete' AND id= $1 RETURNING *"
	result := d.DB.Raw(query, userID).Scan(&user)
	if result.Error != nil {
		return nil, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_auth_server.ErrNotFound
	}
	return &user, nil
}

func (d *AdminRepository) UnBlockUserAccount(userID string) (*responsemodel_auth_server.AbstractUserDetails, error) {
	var user responsemodel_auth_server.AbstractUserDetails
	query := "UPDATE users SET status = 'active' WHERE status = 'block' AND id= $1 RETURNING *"
	result := d.DB.Raw(query, userID).Scan(&user)
	if result.Error != nil {
		return nil, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_auth_server.ErrNotFound
	}
	return &user, nil
}
