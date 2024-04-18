package usecase_server_svc

import (
	"context"
	"encoding/base64"
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
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
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

const friendChatNameSpace = "/friend/"

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

	ok := socket.JoinRoom("/", friendChatNameSpace+userID, conn)
	fmt.Println("user chat room ", ok, userID)

	return nil
}

func (s *serverServiceUseCase) BroadcastMessage(msg []byte, socker *socketio.Server, conn socketio.Conn) {
	message, err := s.JsonMarshelServerMessage(msg)
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	validateError := helper_api_gateway.Validator(message)
	if len(validateError) > 0 {
		conn.Emit("error", validateError)
		return
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userProfile, err := s.authClind.UserProfile(context, &auth.UserProfileRequest{UserID: strconv.Itoa(message.UserID)})
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	message.UserProfilePhoto = userProfile.ProfilePhoto
	message.UserName = userProfile.UserName

	if message.Type != "text" {
		message.Content, err = s.uploadMediaToS3(message.Content)
		if err != nil {
			conn.Emit("error", err.Error())
			return
		}
	}

	message.TimeStamp = time.Now().In(s.Location)
	socker.BroadcastToRoom("/", strconv.Itoa(message.ServerID), "broadcast server chat", message)

	message.TimeStamp = message.TimeStamp.UTC()

	err = s.addServerMessageIntoKafa(message)
	if err != nil {
		conn.Emit("error", "error on kafka producer "+err.Error())
	}
}

func (s *serverServiceUseCase) SendFriendChat(msg []byte, socket *socketio.Server, conn socketio.Conn) {
	message, err := s.jsonUnmarshelFriendlyMessage(msg)
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	validateError := helper_api_gateway.Validator(message)
	if len(validateError) > 0 {
		conn.Emit("error", validateError)
		return
	}

	if message.Type != "text" {
		message.Content, err = s.uploadMediaToS3(message.Content)
		if err != nil {
			conn.Emit("error", err.Error())
			return
		}
	}

	fmt.Println("==-", message)

	message.Timestamp = time.Now().In(s.Location)
	message.Status = "send"
	if socket.RoomLen("/", friendChatNameSpace+message.RecipientID) == 0 {
		message.Status = "pending"
	} else {
		socket.BroadcastToRoom("/", friendChatNameSpace+message.RecipientID, "receive friendly chat", message)
	}

	fmt.Println("----------user count in room", socket.RoomLen("/", friendChatNameSpace+message.RecipientID))
	message.Timestamp = message.Timestamp.UTC()

	err = s.addFriendMessageIntoKafka(message)
	if err != nil {
		conn.Emit("error", err.Error())
	}
}

func (s *serverServiceUseCase) JsonMarshelServerMessage(data []byte) (requestmodel_server_svc.ServerMessage, error) {
	var message requestmodel_server_svc.ServerMessage
	err := json.Unmarshal(data, &message)
	return message, err
}

func (s *serverServiceUseCase) jsonUnmarshelFriendlyMessage(data []byte) (requestmodel_server_svc.FriendlyMessage, error) {
	var msg requestmodel_server_svc.FriendlyMessage
	err := json.Unmarshal(data, &msg)
	return msg, err
}

func (s *serverServiceUseCase) addServerMessageIntoKafa(msg requestmodel_server_svc.ServerMessage) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	msg.UserProfilePhoto = ""
	msg.UserName = ""
	message := s.marshelStruct(msg)

	fmt.Println("from kafka server message ", msg)

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

func (s *serverServiceUseCase) addFriendMessageIntoKafka(message requestmodel_server_svc.FriendlyMessage) error {
	fmt.Println("from kafka friend message ", message)

	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{s.config.KafkaPort}, configs)
	if err != nil {
		return err
	}

	result := s.marshelStruct(message)

	msg := &sarama.ProducerMessage{Topic: s.config.KafkaTopic, Key: sarama.StringEncoder("Amigo Chat"), Value: sarama.StringEncoder(result)}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return err
	}
	// log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return nil
}

func (s *serverServiceUseCase) marshelStruct(msg interface{}) []byte {
	message, _ := json.Marshal(msg)
	return message
}

func (s *serverServiceUseCase) uploadMediaToS3(media string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(media)
	if err != nil {
		fmt.Println("=", err)
	}

	s3Session := helper_api_gateway.CreateSession(s.config.S3)
	url, err := helper_api_gateway.UploadImageToS3(data, s3Session)
	if err != nil {
		return "", err
	}

	return url, nil

	// reader := bytes.NewReader(data)
	// img, format, err := image.Decode(reader)
	// if err != nil {
	// 	fmt.Println("==", err)
	// 	return "", err
	// }
	// fmt.Println("Image format:", format)

	// out, err := os.Create("decoded_image.jpg")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return "", err
	// }
	// defer out.Close()

	// err = jpeg.Encode(out, img, nil)
	// if err != nil {
	// 	fmt.Println("Error saving image:", err)
	// 	return "", err
	// }

	// fmt.Println("Image saved as decoded_image.jpg")
}

// func (s *serverServiceUseCase) storeConnInRedis(conn socketio.Conn, userID string) error {
// 	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	byteConn, err := json.Marshal(conn)
// 	if err != nil {
// 		return err
// 	}

// 	result := s.redisDB.Set(context, userID, byteConn, 4*time.Hour)
// 	fmt.Println("==", result.Val(), result.Err(), conn)

// 	fmt.Println("##", s.redisDB.Get(context, userID))

// 	return nil
// }
