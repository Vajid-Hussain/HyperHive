package responsemodel_friend_server

import "time"

type FriendRequest struct {
	FriendsID string
	User      string `gorm:"column:users"`
	Friend    string
	UpdateAt  time.Time
	Status    string
}

type FriendList struct {
	UniqueFriendID string `gorm:"column:friends_id"`
	FriendID       string `gorm:"column:friend"`
	UpdateAt       time.Time
	LastMessage    interface{} `gorm:"-"`
	UserProfile    AbstractUserProfile `gorm:"-"`
}

type AbstractUserProfile struct {
	UserID       string
	UserName     string
	Name         string
	ProfilePhoto string
}
