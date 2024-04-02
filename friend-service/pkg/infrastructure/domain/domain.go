package domain_friend_server

import "time"

type Friends struct {
	FriendsID string `gorm:"primary key"`
	User      string
	Friend    string
	UpdateAt  time.Time
	Status    string `gorm:"default:pending"`
}
