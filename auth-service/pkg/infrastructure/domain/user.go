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
	UserName        string `gorm:"not null, unique"`
	Name            string
	Email           string `gorm:"not null, unique"`
	ProfilePhotoUrl string
	CoverPhotoUrl   string
	Password        string
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	Status          Status    `gorm:"default:active"`
}
