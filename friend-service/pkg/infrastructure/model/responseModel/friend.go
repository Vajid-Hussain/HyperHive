package responsemodel_friend_server

import "time"

type FriendRequest struct {
	FriendShipID string `gorm:"column:friend_ship_id"`
	User      string `gorm:"column:users"`
	Friend    string 
	UpdateAt  time.Time
	Status    string
}

type FriendList struct {
	UniqueFriendID string `gorm:"column:friend_ship_id"`
	UserID string `gorm:"column:users"`
	FriendID       string `gorm:"column:friend"`
	UpdateAt       time.Time
	LastMessage    interface{}         `gorm:"-"`
	UserProfile    AbstractUserProfile `gorm:"-"`
}

type AbstractUserProfile struct {
	UserID       string
	UserName     string
	Name         string
	ProfilePhoto string
}
