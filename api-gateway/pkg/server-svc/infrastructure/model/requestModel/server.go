package requestmodel_server_svc

import "time"

type Server struct {
	Name string `json:"Name" validate:"required"`
}

type CreateCategory struct {
	UserID       string `json:"-"`
	ServerID     string `json:"serverID" validate:"required"`
	CategoryName string `json:"categoryName" validate:"required"`
}

type CreateChannel struct {
	ChannelName string `json:"channelName" validate:"required"`
	UserID      string `json:"-"`
	ServerID    string `json:"serverID" validate:"required"`
	CategoryID  string `json:"categoryID" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

type JoinToServer struct {
	UserID   string `json:"-"`
	ServerID string `json:"ServerID" validate:"required"`
}

type ServerReq struct {
	ServerID string `json:"ServerID" param:"id" validate:"required"`
}

type ServerMessage struct {
	UserProfilePhoto string `json:"UserProfilePhoto,omitempty"`
	UserName         string `json:"UserName"`
	UserID           int    `json:"UserID" validate:"required"`
	ChannelID        int    `json:"ChannelID" validate:"required"`
	ServerID         int    `json:"ServerID" validate:"required"`
	Content          string `json:"Content" validate:"required"`
	TimeStamp        time.Time
	Type             string `json:"Type"`
}

type KafkaServerMessage struct {
	UserID    int    `json:"UserID"`
	ChannelID int    `json:"ChannelID"`
	ServerID  int    `json:"ServerID"`
	Content   string `json:"content"`
	TimeStamp time.Time
	Type      string `json:"type"`
}

type ChatRequest struct {
	ChannelID string `query:"ChannelID" validate:"required"`
	Offset    string `query:"Page" validate:"required"`
	Limit     string `query:"Limit" validate:"required"`
}

type ServerDescription struct {
	ServerID    string `json:"ServerID" validate:"required"`
	Description string `json:"Description" validate:"max=20"`
}

type ServerMember struct {
	ServerID string `query:"ServerID" validate:"required"`
	Offset   string `query:"Page" validate:"required"`
	Limit    string `query:"Limit" validate:"required"`
}

type RemoveUser struct {
	RemoveUserID string `json:"RemoveUserID" validate:"required"`
	ServerID     string `json:"ServerID" validate:"required"`
}

type UpdateMemberRole struct {
	TargetUserID string `json:"TargetUserID" validate:"required"`
	TargetRole   string `json:"TargetRole" validate:"required"`
	ServerID     string `json:"ServerID" validate:"required"`
}

type MessageType struct {
	Category string `json:"Category"`
}

type FriendlyMessage struct {
	SenderID    string `json:"SenderID" validate:"required"`
	RecipientID string `json:"RecipientID" validate:"required"`
	Content     string `json:"Content" validate:"required"`
	Timestamp   time.Time
	Type        string `json:"Type" validate:"required"`
	Tag         string `json:"Tag"`
	Status      string `json:"Status"`
}
