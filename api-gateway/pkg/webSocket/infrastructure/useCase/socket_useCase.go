package usecase_websocket_svc

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	requestmodel_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/model/requestModel"
	responsemodel_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/model/responseModel"
	"github.com/gorilla/websocket"
)

type WebSocketUseCase struct {
	config   config.Config
	Location *time.Location
}

func NewWebSocketUseCase(config config.Config) *WebSocketUseCase {
	locationInd, _ := time.LoadLocation("Asia/Kolkata")
	return &WebSocketUseCase{
		config:   config,
		Location: locationInd,
	}
}

var user = make(map[string]*websocket.Conn)

func (w *WebSocketUseCase) PreprocessOfSendingMessage(msg []byte, userCollection map[string]*websocket.Conn, userID string, err error) {
	user = userCollection // set user map as global veriable
	if err != nil {
		w.sendErrorMessage(userID, err)
		return
	}

	messageCategory := w.findMessageCategory(msg, userID)

	switch messageCategory {
	case "friend chat":
		w.sendFrindlyChat(msg, userID)
	case "group chat":
		// groupfunc()
	default:
		w.sendErrorMessage(userID, responsemodel_websocket_svc.ErrUndefinedMessageCategory)
	}
}

func (w *WebSocketUseCase) sendFrindlyChat(msg []byte, userID string) {
	var message requestmodel_websocket_svc.Message
	err := json.Unmarshal(msg, &message)
	if err != nil {
		w.sendErrorMessage(userID, responsemodel_websocket_svc.ErrWhileUnmarshelChatMessage)
	}

	message.Status = "pending"
	message.Timestamp = time.Now()
	message.SenderID = userID

	conn, ok := user[message.RecipientID]
	if ok {
		message.Timestamp = message.Timestamp.In(w.Location)
		err := conn.WriteMessage(websocket.TextMessage, w.marshelStruct(message))
		if err != nil {
			delete(user, message.RecipientID)
		} else {
			message.Status = "send"
		}
	}
	w.kafkaProducer(message)
}

func (w *WebSocketUseCase) kafkaProducer(message requestmodel_websocket_svc.Message) error {
	fmt.Println("from kafka ", message)

	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{w.config.KafkaPort}, configs)
	if err != nil {
		return err
	}

	result := w.marshelStruct(message)

	msg := &sarama.ProducerMessage{Topic: w.config.KafkaTopic, Key: sarama.StringEncoder("Amigo Chat"), Value: sarama.StringEncoder(result)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("err send message in kafka ", err)
	}
	log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}

func (w *WebSocketUseCase) findMessageCategory(msg []byte, userID string) string {
	var messageType requestmodel_websocket_svc.MessageType
	err := json.Unmarshal(msg, &messageType)
	if err != nil {
		w.sendErrorMessage(userID, responsemodel_websocket_svc.ErrUndefinedMessageCategory)
	}
	return messageType.Category
}

func (w *WebSocketUseCase) sendErrorMessage(userID string, err error) {
	conn, ok := user[userID]
	if ok {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	}
}

func (w *WebSocketUseCase) marshelStruct(msg interface{}) []byte {
	message, _ := json.Marshal(msg)
	return message
}
