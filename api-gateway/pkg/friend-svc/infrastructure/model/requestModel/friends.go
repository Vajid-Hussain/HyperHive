package requestmodel_friend_svc

import "time"

type FriendRequest struct {
	FriendID string `json:"FriendID" validate:"required"`
}

type FrendShipStatusUpdate struct {
	FrendShipID string `json:"FrendShipID"`
}

type Message struct {
	SenderID    string    `json:"SenderID" validate:"required"`
	RecipientID string    `json:"RecipientID" validate:"required"`
	Content     string    `json:"Content" validate:"required"`
	Timestamp   time.Time `json:"TimeStamp" validate:"required"`
	Type        string    `json:"Type" validate:"required"`
	Tag         string    `json:"Tag"`
	Status      string    `json:"Status"`
}

type WebSocketInfo struct {
	RemoteAddr string `json:"remote_addr"`
}

type ChatRequest struct {
	FriendID string `query:"FriendID" validate:"required"`
	Offset   string `query:"Offset" validate:"required"`
	Limit    string `query:"Limit" validate:"required"`
}

type Sample struct {
	SenderID string `json:"SenderID" validate:"required"`
}

