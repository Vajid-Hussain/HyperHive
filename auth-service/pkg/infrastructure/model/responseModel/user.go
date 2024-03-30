package responsemodel_auth_server

import (
	"fmt"
	"time"
)

var (
	ErrRegexNotMatch    = fmt.Errorf("password is not satify criteria")
	ErrDBNoRowAffected  = fmt.Errorf("no row affected")
	ErrDBQueryExecution = fmt.Errorf("request have some missaderstanding polisht your request")
)

type UserSignup struct {
	ID              string
	UserName        string
	Name            string
	Email           string
	Password        string
	ProfilePhotoUrl string
	CoverPhotoUrl   string
	CreatedAt       time.Time `gorm:"column:created_at"`
	Status          string
	TemperveryToken string
}

type AuthenticationResponse struct {
	AccesToken   string
	RefreshToken string
}

type UserProfile struct {
	UserID           string    `gorm:"column:id"`
	UserName         string    `gorm:"column:user_name"`
	Name             string    `gorm:"column:name"`
	Email            string    `gorm:"column:email"`
	ProfilePhoto     string    `gorm:"column:profile_photo_url"`
	CoverPhoto       string    `gorm:"column:cover_photo_url"`
	Description      string    `gorm:"column:description"`
	Status           string    `gorm:"column:status"`
	StatusExpireTime time.Time `gorm:"column:status_till"`
	UserSince        time.Time `gorm:"column:created_at"`
}

type AbstractUserDetails struct{
	UserID string 
	UserName string
	Name string
}
