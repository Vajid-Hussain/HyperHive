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
	CreatedAt        time.Time `gorm:"column:created_at"`
	Status          string
	TemperveryToken string
}

type AuthenticationResponse struct{
	AccesToken string 
	RefreshToken string
}