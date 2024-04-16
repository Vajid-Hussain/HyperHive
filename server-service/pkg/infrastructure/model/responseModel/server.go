package responsemodel_server_service

import (
	"time"
)

type Server struct {
	ServerID    string `gorm:"column:id"`
	Name        string
	Description string
	Icon        string
	CoverPhoto  string
}

type ChannelCategory struct {
	CategoryID string
	ServerID   string
	Name       string
}

type ServerAdmin struct {
	ID       string
	UserID   string
	ServerID string
	Role     string
}

type Channel struct {
	ChannelID  string
	CategoryID string
	Name       string
	Type       string
}

type FullServerChannel struct {
	CategoryID string
	Name       string
	Channel    []*Channel `gorm:"-"`
}

type UserServerList struct {
	ServerID string
}

type ServerMessage struct {
	ID        string    `bson:"_id,omitempty"`
	UserID    int       `bson:"UserID"`
	ChannelID int       `bson:"ChannelID"`
	ServerID  int       `bson:"ServerID"`
	Content   string    `bson:"Content"`
	TimeStamp time.Time `bson:"TimeStamp"`
	Type      string    `bson:"Type"`
}
