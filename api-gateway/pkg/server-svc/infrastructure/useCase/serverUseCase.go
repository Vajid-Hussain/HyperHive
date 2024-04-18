package usecase_server_svc

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	auth "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	requestmodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/requestModel"
	interface_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/useCase/interface"
	server "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	socketio "github.com/googollee/go-socket.io"
	"github.com/redis/go-redis/v9"
)

type serverServiceUseCase struct {
	clind     server.ServerClient
	Location  *time.Location
	authClind auth.AuthServiceClient
	config    *config.Config
	redisDB   *redis.Client
}

func NewServerServiceUseCase(clind server.ServerClient, authClind auth.AuthServiceClient, config *config.Config, redisDB *redis.Client) interface_server_svc.IserverServiceUseCase {
	locationInd, _ := time.LoadLocation("Asia/Kolkata")
	return &serverServiceUseCase{
		clind:     clind,
		Location:  locationInd,
		authClind: authClind,
		config:    config,
		redisDB:   redisDB,
	}
}

func (s *serverServiceUseCase) JoinToServerRoom(userID string, socket *socketio.Server, conn socketio.Conn) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userServerList, err := s.clind.GetUserServer(context, &server.GetUserServerRequest{UserID: userID})
	if err != nil {
		return err
	}

	for _, serverID := range userServerList.ServerId {
		ok := socket.JoinRoom("/", serverID, conn)
		fmt.Println("join room ", ok, serverID)
	}

	ok := socket.JoinRoom("/", "/friend/"+userID, conn)
	fmt.Println("user chat room ", ok, userID)

	return nil
}

func (s *serverServiceUseCase) BroadcastMessage(userID string, msg []byte, socker *socketio.Server, conn socketio.Conn) {
	message := s.JsonMarshelServerMessage(msg)
	fmt.Println("message ", message)
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userProfile, err := s.authClind.UserProfile(context, &auth.UserProfileRequest{UserID: strconv.Itoa(message.UserID)})
	if err != nil {
		fmt.Println("====---", err)
		conn.Emit("error", err.Error())
	}

	message.UserProfilePhoto = userProfile.ProfilePhoto
	message.UserName = userProfile.UserName
	message.UserID, _ = strconv.Atoi(userID)

	message.TimeStamp = time.Now().In(s.Location)
	socker.BroadcastToRoom("/", strconv.Itoa(message.ServerID), "broadcast server chat", message)

	fmt.Println("==", message)
	message.TimeStamp = message.TimeStamp.UTC()
	err = s.addMessageIntoKafa(message)
	if err != nil {
		conn.Emit("error", "error on kafka producer "+err.Error())
	}
}

func (s *serverServiceUseCase) SendFriendChat(userID string, msg []byte, socket *socketio.Server, conn socketio.Conn) {
	message, err := s.jsonUnmarshelFriendlyMessage(msg)
	if err != nil {
		conn.Emit("error", err.Error())
	}

	fmt.Println("==-", message)
	message.Timestamp = time.Now().In(s.Location)
	message.SenderID = userID
	ok := socket.BroadcastToRoom("/", "/friend/"+message.RecipientID, "receive friendly chat", message)
	fmt.Println("=== friendly message", message, ok)
}

func (s *serverServiceUseCase) JsonMarshelServerMessage(data []byte) requestmodel_server_svc.ServerMessage {
	var message requestmodel_server_svc.ServerMessage
	json.Unmarshal(data, &message)
	return message
}

func (s *serverServiceUseCase) jsonUnmarshelFriendlyMessage(data []byte) (requestmodel_server_svc.FriendlyMessage, error) {
	var msg requestmodel_server_svc.FriendlyMessage
	err := json.Unmarshal(data, &msg)
	return msg, err
}

func (s *serverServiceUseCase) jsonMarshelFriendlyMessage(data requestmodel_server_svc.FriendlyMessage) ([]byte, error) {
	return json.Marshal(data)
}

func (s *serverServiceUseCase) addMessageIntoKafa(msg requestmodel_server_svc.ServerMessage) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	msg.UserProfilePhoto = ""
	msg.UserName = ""
	message := s.marshelStruct(msg)

	fmt.Println("kafka message ", msg)

	producer, err := sarama.NewSyncProducer([]string{s.config.KafkaPort}, config)
	if err != nil {
		return err
	}

	finalMessage := sarama.ProducerMessage{Topic: s.config.KafkaServerTopic, Key: sarama.StringEncoder("server message"), Value: sarama.StringEncoder(message)}
	_, _, err = producer.SendMessage(&finalMessage)
	if err != nil {
		return err
	}
	return nil
}

func (s *serverServiceUseCase) marshelStruct(msg interface{}) []byte {
	message, _ := json.Marshal(msg)
	return message
}

func (s *serverServiceUseCase) storeConnInRedis(conn socketio.Conn, userID string) error {
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	byteConn, err := json.Marshal(conn)
	if err != nil {
		return err
	}

	result := s.redisDB.Set(context, userID, byteConn, 4*time.Hour)
	fmt.Println("==", result.Val(), result.Err(), conn)

	fmt.Println("##", s.redisDB.Get(context, userID))

	return nil
}
