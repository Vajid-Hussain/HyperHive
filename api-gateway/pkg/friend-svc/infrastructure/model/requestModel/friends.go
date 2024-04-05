package requestmodel_friend_svc

import "time"

type FriendRequest struct {
	FriendID string `json:"FriendID"`
}

type FrendShipStatusUpdate struct {
	FrendShipID string `json:"FrendShipID"`
}

type Message struct {
	SenderID    string    `json:"sender_id"`
	RecipientID string    `json:"recipient_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}

type WebSocketInfo struct {
	RemoteAddr string `json:"remote_addr"`
}
