package requestmodel_websocket_svc

import "time"

type MessageType struct {
	Category string `json:"Category"`
}

type Message struct {
	SenderID    string    `json:"SenderID" validate:"required"`
	RecipientID string    `json:"RecipientID" validate:"required"`
	Content     string    `json:"Content" validate:"required"`
	Timestamp   time.Time `json:"_"`
	Type        string    `json:"Type" validate:"required"`
	Tag         string    `json:"Tag"`
	Status      string    `json:"Status"`
}
