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

type ServerMembers struct {
	UserID      string `gorm:"column:user_id"`
	UserProfile string
	UserName    string
	Name        string
	Role        string `gorm:"column:role"`
}

type ForumPost struct {
	ID              string `bson:"_id"`
	UserProfile     string
	UserName        string
	UserID          int       `json:"UserID" bson:"UserID"`
	ChannelID       int       `json:"ChannelID" bson:"ChannelID"`
	ServerID        int       `json:"ServerID" bson:"ServerID"`
	Content         string    `json:"Content" bson:"Content"`
	MainContentType string    `json:"MainContentType" bson:"MainContentType"`
	SubContent      string    `json:"SubContent" bson:"SubContent"`
	TimeStamp       time.Time `json:"TimeStamp" bson:"TimeStamp"`
	Type            string    `json:"Type" bson:"Type"`
	CommandContent  string    `bson:"content"`
}

type ForumCommand struct {
	ID          string `bson:"_id"`
	UserProfile string
	UserName    string
	UserID      int       `json:"UserID" validate:"required"`
	ChannelID   int       `json:"ChannelID" validate:"required"`
	ServerID    int       `json:"ServerID" validate:"required"`
	ParentID    string    `json:"parentID" validate:"required"`
	Content     string    `json:"Content" validate:"required"`
	TimeStamp   time.Time `bson:"timestamp"`
	Type        string    `json:"Type"`
	Thread      []*ForumCommand
}

type PostCommands struct {
	Commands []*ForumCommand
}
