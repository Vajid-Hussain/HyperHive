package requestmodel_friend_svc

import "time"

type FriendRequest struct {
	FriendID string `json:"FriendID"`
}

type FrendShipStatusUpdate struct {
	FrendShipID string `json:"FrendShipID"`
}

type Message struct {
	SenderID    string    `json:"SenderID"`
	RecipientID string    `json:"RecipientID"`
	Content     string    `json:"Content"`
	Timestamp   time.Time `json:"TimeStamp"`
	Type        string    `json:"Type"`
	Tag         string    `json:"Tag"`
	Status      string    `json:"Status" default:"send"`
}

type WebSocketInfo struct {
	RemoteAddr string `json:"remote_addr"`
}
