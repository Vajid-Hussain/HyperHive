package domainl_auth_server

import "time"

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "active"
	Rejected Status = "block"
)

type Users struct {
	ID              int    `gorm:"primarykey; autoIncrement"`
	UserName        string `gorm:"not null,unique"`
	Name            string
	Email           string `gorm:"not null;unique;index"`
	ProfilePhotoUrl string
	CoverPhotoUrl   string
	Password        string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	Status          Status    `gorm:"default:pending"`
}

type UserProfileStatus struct {
	StatusID    int   `gorm:"unique"`
	Frkey       Users `gorm:"foreignkey:StatusID;referances:ID"`
	Status      string
	StatusTill  time.Time
	Description string
}

type Otp struct {
	Expire time.Time
	Otp    string
	Emails string `gorm:"primarykey"`
	// User   Users  `gorm:"foreignKey:Emails; references:Email"`
}
