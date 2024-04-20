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

type Message struct {
	SenderID    string    `json:"SenderID"`
	RecipientID string    `json:"RecipientID"`
	Content     string    `json:"Content"`
	Timestamp   time.Time `json:"TimeStamp"`
	Type        string    `json:"Type"`
	Tag         string    `json:"Tag"`
	Status      string    `json:"Status" default:"send"`
}

type Pagination struct {
	Limit  string
	OffSet string
}

type FriendShipStatus struct{
	UserId string
	FriendShipID string
	Status string
}