package requestmodel_friend_svc

import "time"

type FriendRequest struct {
	FriendID string `json:"FriendID"`
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
	FriendID string `query:"FriendID"`
	Offset   string `query:"Offset"`
	Limit    string `query:"Limit"`
}

type Sample struct {
	SenderID string `json:"SenderID" validate:"required"`
}
