package domain_friend_server

import "time"

type Friends struct {
	FriendsID string `gorm:"primaryKey;autoincrement;unique; type:integer"`
	Users     string
	Friend    string
	UpdateAt  time.Time
	Status    string `gorm:"default:pending"`
}
