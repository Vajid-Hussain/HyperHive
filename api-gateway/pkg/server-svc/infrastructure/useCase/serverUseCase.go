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
	resonsemodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/resonseModel"
	interface_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/useCase/interface"
	server "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	socketio "github.com/googollee/go-socket.io"
)

type serverServiceUseCase struct {
	clind     server.ServerClient
	Location  *time.Location
	authClind auth.AuthServiceClient
	config    *config.Config
}

func NewServerServiceUseCase(clind server.ServerClient, authClind auth.AuthServiceClient, config *config.Config) interface_server_svc.IserverServiceUseCase {
	locationInd, _ := time.LoadLocation("Asia/Kolkata")
	return &serverServiceUseCase{
		clind:     clind,
		Location:  locationInd,
		authClind: authClind,
		config:    config,
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
	return nil
}

func (s *serverServiceUseCase) BroadcastMessage(msg []byte, socker *socketio.Server) {
	message := s.JsonMarshelServerMessage(msg)
	fmt.Println("message ", message)
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userProfile, _ := s.authClind.UserProfile(context, &auth.UserProfileRequest{UserID: strconv.Itoa(message.UserID)})
	message.UserProfilePhoto = userProfile.ProfilePhoto
	message.UserName = userProfile.UserName

	message.TimeStamp = time.Now().In(s.Location)
	socker.BroadcastToRoom("/", strconv.Itoa(message.ServerID), "broadcast server chat", message)

	err := s.addMessageIntoKafa(message)
	if err != nil {
		fmt.Println("error on kafka producer ", err)
	}
}

func (s *serverServiceUseCase) EmitErrorMessage(conn socketio.Conn, err string) {
	errMessage := s.JsonMarshelErrorMessage([]byte(err))
	conn.Emit("error", errMessage)
}

func (s *serverServiceUseCase) JsonMarshelErrorMessage(data []byte) resonsemodel_server_svc.ErrorMessage {
	var errMessage resonsemodel_server_svc.ErrorMessage
	json.Unmarshal(data, &errMessage)
	return errMessage
}

func (s *serverServiceUseCase) JsonMarshelServerMessage(data []byte) requestmodel_server_svc.ServerMessage {
	var message requestmodel_server_svc.ServerMessage
	json.Unmarshal(data, &message)
	return message
}

func (s *serverServiceUseCase) addMessageIntoKafa(msg requestmodel_server_svc.ServerMessage) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	message := s.marshelStruct(msg)

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
