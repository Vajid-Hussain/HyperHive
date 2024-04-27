package requestmodel_friend_server

import "time"

type FriendRequest struct {
	User     string
	Friend   string
	UpdateAt time.Time
}

type GetFriendRequest struct {
	UserID string
	Limit  string
	OffSet string
}

//	type Message struct {
//		SenderID    string    `json:"SenderID"`
//		RecipientID string    `json:"RecipientID"`
//		Content     string    `json:"Content"`
//		Timestamp   time.Time `json:"TimeStamp"`
//		Type        string    `json:"Type"`
//		Tag         string    `json:"Tag"`
//		Status      string    `json:"Status" default:"send"`
//	}
type Message struct {
	SenderID    string    `json:"sender_id" validate:"required"`
	RecipientID string    `json:"recipient_id" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Timestamp   time.Time `json:"timestamp"`
	Type        string    `json:"type" validate:"required"`
	Tag         string    `json:"Tag"`
	Status      string    `json:"status"`
}

type Pagination struct {
	Limit  string
	OffSet string
}

type FriendShipStatus struct {
	UserId       string
	FriendShipID string
	Status       string
}
