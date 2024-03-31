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

	query := "INSERT INTO users (name, user_name, email, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING *"

	result := d.DB.Raw(query, userReq.Name, userReq.UserName, userReq.Email, userReq.Password, userReq.CreatedAt).Scan(&userRes)
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
	query := "SELECT count(*) FROM users WHERE user_name = $1  AND status != 'delete'"
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
	query := "SELECT count(*) FROM users WHERE email = $1 AND status != 'delete'"
	result := d.DB.Raw(query, email).Scan(&count)
	if result.Error != nil {
		return 0, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return 0, responsemodel_auth_server.ErrDBNoRowAffected
	}
	return count, nil
}

func (d *UserRepository) FetchMailUsingUserID(userID string) (email string, err error) {
	query := "SELECT email FROM users where id = $1"
	result := d.DB.Raw(query, userID).Scan(&email)
	if result.Error != nil {
		return "", responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", responsemodel_auth_server.ErrUserBlockedOrNoUser
	}
	return
}

func (d *UserRepository) IsUserIDExist(userID string) (count int, err error) {
	query := "SELECT COUNT(*) FROM users WHERE id= $1"
	result := d.DB.Raw(query, userID).Scan(&count)
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

	query := " SELECT count(*) FROM users WHERE id=? AND status ='active' "
	result := d.DB.Raw(query, userID).Scan(&count)
	if result.Error != nil {
		return 0, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return 0, responsemodel_auth_server.ErrNotFound
	}

	return count, nil
}

// OTP

func (d *UserRepository) CreateOtp(otp, email string, expire time.Time) error {
	fmt.Println("--", otp, email)
	qyery := "INSERT INTO otps (emails, otp, expire) VALUES ($1, $2, $3) ON CONFLICT (emails) DO UPDATE SET otp = $2, expire = $3"

	result := d.DB.Exec(qyery, email, otp, expire)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}

	return nil
}

func (d *UserRepository) FetchOtp(email string, now time.Time) (otp string, err error) {
	query := "SELECT otp FROM otps WHERE emails = $1 AND expire >= $2"
	result := d.DB.Raw(query, email, now).Scan(&otp)
	if result.Error != nil {
		return "", responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", responsemodel_auth_server.ErrOtpIsExpire
	}

	return otp, nil
}

func (d *UserRepository) ForgotPassword(email, password string) error {
	query := "UPDATE users SET password = $1 WHERE email= $2"
	result := d.DB.Exec(query, password, email)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

func (d *UserRepository) GetUserPasswordUsingEmail(email string) (password string, err error) {
	query := "SELECT password FROM users WHERE email =$1 AND status= 'active'"
	result := d.DB.Raw(query, email).Scan(&password)
	if result.Error != nil {
		return "", responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", responsemodel_auth_server.ErrNotFound
	}

	return password, nil
}

func (d *UserRepository) FetchUserIDUsingMail(email string) (userID string, err error) {
	query := "SELECT id FROM users WHERE email = $1"
	result := d.DB.Raw(query, email).Scan(&userID)

	if result.Error != nil {
		return "", responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return "", responsemodel_auth_server.ErrNotFound
	}

	return userID, nil
}

func (d *UserRepository) UpdateUserProfilePhoto(userID, photoUrl string) error {
	query := "UPDATE  users SET profile_photo_url= $1 WHERE id = $2"
	result := d.DB.Exec(query, photoUrl, userID)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

func (d *UserRepository) UpdateCoverPhoto(userID, photoUrl string) error {
	query := "UPDATE  users SET cover_photo_url= $1 WHERE id = $2"
	result := d.DB.Exec(query, photoUrl, userID)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

// user Profile

func (d *UserRepository) UpdateOrCreateUserStatus(status requestmodel_auth_server.UserProfileStatus) error {

	query := "INSERT INTO user_profile_statuses (status_id, status, status_till) VALUES ($1, $2, $3) ON CONFLICT (status_id) DO UPDATE SET status = $2, status_till = $3"
	result := d.DB.Exec(query, status.UserID, status.Status, status.Expire)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

func (d *UserRepository) UpdateOrCreateUserDescription(userID, description string) error {
	fmt.Println("--", userID, description)
	query := "INSERT INTO user_profile_statuses (status_id, description) VALUES ($1, $2) ON CONFLICT (status_id) DO UPDATE SET description = $2 "
	result := d.DB.Exec(query, userID, description)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}

func (d *UserRepository) GetUserProfile(userID string) (userProfile *responsemodel_auth_server.UserProfile, err error) {
	query := "SELECT * FROM users LEFT JOIN user_profile_statuses ON users.id = user_profile_statuses.status_id WHERE users.status= 'active' AND users.id= $1"
	result := d.DB.Raw(query, userID).Scan(&userProfile)
	if result.Error != nil {
		return nil, responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_auth_server.ErrUserBlockedOrNoUser
	}
	return
}

func (d *UserRepository) DeleteUserAcoount(userID string) error {
	query := "UPDATE users SET status = 'delete' WHERE id= $1"
	result := d.DB.Exec(query, userID)
	if result.Error != nil {
		return responsemodel_auth_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_auth_server.ErrNotFound
	}
	return nil
}
