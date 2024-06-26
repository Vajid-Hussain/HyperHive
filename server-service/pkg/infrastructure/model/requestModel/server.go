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
	UserID   string
	Image    []byte
	Type     string
	Url      string
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

type Description struct {
	UserID      string
	Description string
	ServerID    string
}

type UpdateMemberRole struct {
	UserID       string
	TargetUserID string
	TargetRole   string
	ServerID     string
}

type RemoveUser struct {
	UserID    string
	RemoverID string
	ServerID  string
}

type ForumPost struct {
	UserID          int       `json:"UserID" bson:"UserID"`
	ChannelID       int       `json:"ChannelID" bson:"ChannelID"`
	ServerID        int       `json:"ServerID" bson:"ServerID"`
	Content         string    `json:"Content" bson:"Content"`
	MainContentType string    `json:"MainContentType" bson:"MainContentType"`
	SubContent      string    `json:"SubContent" bson:"SubContent"`
	TimeStamp       time.Time `bson:"TimeStamp" json:"TimeStamp"`
	Type            string    `json:"Type" bson:"Type"`
}

type FormCommand struct {
	UserID           int    `json:"UserID" validate:"required"`
	ChannelID        int    `json:"ChannelID" validate:"required"`
	ServerID         int    `json:"ServerID" validate:"required"`
	ParentID         string `json:"parentID" validate:"required"`
	Content          string `json:"Content" validate:"required"`
	TimeStamp        time.Time
	Type             string `json:"Type"`
}

