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
	UserID           int    `json:"UserID"`
	ChannelID        int    `json:"ChannelID"`
	ServerID         int    `json:"ServerID"`
	Content          string `json:"Content"`
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
