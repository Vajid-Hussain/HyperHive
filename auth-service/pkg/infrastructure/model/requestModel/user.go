package requestmodel_auth_server

import "time"

type UserSignup struct {
	UserName        string
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	ProfilePhoto    []byte
	CoverPhoto      []byte
	ProfilePhotoUrl string
	CoverPhotoUrl   string
	CreatedAt       time.Time
}

type VerifyUser struct {
	Token string
	Email string
}

type UserProfileStatus struct {
	UserID string
	Status string
	Expire string
}

type ForgotPassword struct {
	Password string
	Token    string
	Otp      string
}
