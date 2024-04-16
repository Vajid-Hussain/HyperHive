package requestmodel_server_service

import (
	"time"
)

type Server struct {
	UserID string
	Name   string
}

type ServerAdmin struct {
	UserID   string
	ServerID string
	Role     string
}

type CreateCategory struct {
	UserID       string
	ServerID     string
	CategoryName string
}

type CreateChannel struct {
	ChannelName string
	UserID      string
	ServerID    string
	CategoryID  string
	Type        string
}

type JoinToServer struct {
	UserID   string
	ServerID string
	Role     string
}

type MemberStatusUpdate struct {
	UserID       string
	ServerID     string
	TargetUserID string
}

type ServerImages struct {
	ServerID string
	Image    []byte
	Type     string
}

type ServerMessage struct {
	// ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    int       `bson:"UserID" json:"UserID"`
	ChannelID int       `bson:"ChannelID" json:"ChannelID"`
	ServerID  int       `bson:"ServerID" json:"ServerID"`
	Content   string    `bson:"Content" json:"Content"`
	TimeStamp time.Time `bson:"TimeStamp" json:"TimeStamp"`
	Type      string    `bson:"Type" json:"Type"`
}

type Pagination struct {
    Limit  string
    OffSet string
}
