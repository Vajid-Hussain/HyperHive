package usecase_server_svc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	_ "time/tzdata"

	"github.com/IBM/sarama"
	di_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/di"
	auth "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	requestmodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/requestModel"
	resonsemodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/resonseModel"
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
	authCache *helper_api_gateway.RedisCaching
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
		authCache: di_auth_svc.AuthCache(),
	}
}

func (s *serverServiceUseCase) JoinToServerRoom(userID string, socket *socketio.Server, conn socketio.Conn) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userServerList, err := s.clind.GetUserServer(context, &server.GetUserServerRequest{UserID: userID})
	if err != nil {
		return err
	}

	for _, data := range userServerList.UserServerList {
		ok := socket.JoinRoom("/", data.ServerId, conn)
		fmt.Println("join room ", ok, data.ServerId)
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

	// context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()

	// userProfile, err := s.authClind.UserProfile(context, &auth.UserProfileRequest{UserID: strconv.Itoa(message.UserID)})
	userProfile, err := s.authCache.GetUserProfile(strconv.Itoa(message.UserID))
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	message.UserProfilePhoto = userProfile.ProfilePhoto
	message.UserName = userProfile.UserName

	if message.Type == "file" {
		message.Content, err = s.uploadMediaToS3(message.Content)
		if err != nil {
			conn.Emit("error", err.Error())
			return
		}
	} else if message.Type != "text" && message.Type != "file" {
		conn.Emit("error", resonsemodel_server_svc.ErrServerMessageType.Error())
		return
	}

	message.TimeStamp = time.Now().In(s.Location)
	socker.BroadcastToRoom("/", strconv.Itoa(message.ServerID), "broadcast server chat", message)

	message.TimeStamp = message.TimeStamp.UTC()
	message.UserProfilePhoto = ""
	message.UserName = ""

	err = s.addServerMessageIntoKafa(message)
	if err != nil {
		conn.Emit("error", "error on kafka producer "+err.Error())
	}
}

func (s *serverServiceUseCase) SendFriendChat(msg []byte, socket *socketio.Server, conn socketio.Conn) {
	fmt.Println("===",msg)
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

	if message.Type == "file" {
		message.Content, err = s.uploadMediaToS3(message.Content)
		if err != nil {
			conn.Emit("error", err.Error())
			return
		}
	} else if message.Type != "text" {
		conn.Emit("error", resonsemodel_server_svc.ErrUserMessageSupportType)
		return
	}

	fmt.Println("==-", message)

	message.Timestamp = time.Now().In(s.Location)
	message.Status = "send"

	// if socket.RoomLen("/", friendChatNameSpace+message.RecipientID) == 0 {
	// 	message.Status = "pending"
	// } else {
	// 	socket.BroadcastToRoom("/", friendChatNameSpace+message.RecipientID, "receive friendly chat", message)
	// }

	socket.BroadcastToRoom("/", friendChatNameSpace+message.RecipientID, "receive friendly chat", message)

	
	// fmt.Println("----------user count in room", socket.RoomLen("/", friendChatNameSpace+message.RecipientID))
	message.Timestamp = message.Timestamp.UTC()

	err = s.addFriendMessageIntoKafka(message)
	if err != nil {
		conn.Emit("error", err.Error())
	}
	fmt.Println("finish message send succesfully ",err)
}

func (s *serverServiceUseCase) BroadcastForum(msg []byte, soket socketio.Server, conn socketio.Conn) {
	types, err := s.jsonUnmarshelFormType(msg)
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	if types.Type == "post" {
		s.broadCastForumPost(msg, soket, conn)
	} else if types.Type == "command" {
		s.broadcastForumCommands(msg, soket, conn)
	} else {
		conn.Emit("error", resonsemodel_server_svc.ErrForumUnexpectedType.Error())
	}
}

func (s *serverServiceUseCase) broadCastForumPost(msg []byte, soket socketio.Server, conn socketio.Conn) {
	post, err := s.jsonUnmarshelForumPost(msg)
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	validateError := helper_api_gateway.Validator(post)
	if len(validateError) > 0 {
		conn.Emit("error", validateError)
		return
	}

	// userProfile, err := s.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: strconv.Itoa(post.UserID)})
	userProfile, err := s.authCache.GetUserProfile(strconv.Itoa(post.UserID))
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	post.UserName = userProfile.UserName
	post.UserProfilePhoto = userProfile.ProfilePhoto
	post.TimeStamp = time.Now().In(s.Location)

	if post.MainContentType == "image" {
		post.Content, err = s.uploadMediaToS3(post.Content)
		if err != nil {
			conn.Emit("error", err.Error())
			return
		}
	} else if post.MainContentType != "text" {
		conn.Emit("error", resonsemodel_server_svc.ErrForumPostUnexpectedContent.Error())
		return
	}

	soket.BroadcastToRoom("/", strconv.Itoa(post.ServerID), "broadcast forum", post)

	post.TimeStamp = post.TimeStamp.UTC()
	post.UserProfilePhoto = ""
	post.UserName = ""

	err = s.addServerMessageIntoKafa(post)
	if err != nil {
		conn.Emit("error", err.Error())
	}
}

func (s *serverServiceUseCase) broadcastForumCommands(data []byte, soket socketio.Server, conn socketio.Conn) {
	command, err := s.jsonUnmarshelFormCommands(data)
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	validateError := helper_api_gateway.Validator(command)
	if len(validateError) > 0 {
		conn.Emit("error", validateError)
		return
	}

	// userProfile, err := s.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: strconv.Itoa(command.UserID)})
	userProfile, err := s.authCache.GetUserProfile(strconv.Itoa(command.UserID))
	if err != nil {
		conn.Emit("error", err.Error())
		return
	}

	command.UserProfilePhoto = userProfile.ProfilePhoto
	command.UserName = userProfile.Name
	command.TimeStamp = time.Now().In(s.Location)

	soket.BroadcastToRoom("/", strconv.Itoa(command.ServerID), "broadcast forum", command)

	command.UserProfilePhoto = ""
	command.UserName = ""
	command.TimeStamp = command.TimeStamp.UTC()

	err = s.addServerMessageIntoKafa(command)
	if err != nil {
		conn.Emit("error", err.Error())
	}
}

func (s *serverServiceUseCase) JsonMarshelServerMessage(data []byte) (requestmodel_server_svc.ServerMessage, error) {
	var message requestmodel_server_svc.ServerMessage
	return message, json.Unmarshal(data, &message)
}

func (s *serverServiceUseCase) jsonUnmarshelFriendlyMessage(data []byte) (requestmodel_server_svc.FriendlyMessage, error) {
	var msg requestmodel_server_svc.FriendlyMessage
	return msg, json.Unmarshal(data, &msg)
}

func (s *serverServiceUseCase) jsonUnmarshelFormType(data []byte) (requestmodel_server_svc.FormType, error) {
	var types requestmodel_server_svc.FormType
	return types, json.Unmarshal(data, &types)
}

func (s *serverServiceUseCase) jsonUnmarshelForumPost(data []byte) (requestmodel_server_svc.ForumPost, error) {
	var msg requestmodel_server_svc.ForumPost
	return msg, json.Unmarshal(data, &msg)
}

func (s *serverServiceUseCase) jsonUnmarshelFormCommands(data []byte) (requestmodel_server_svc.FormCommand, error) {
	var command requestmodel_server_svc.FormCommand
	return command, json.Unmarshal(data, &command)
}

func (s *serverServiceUseCase) addServerMessageIntoKafa(data interface{}) error {
	var kafkaKey string
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	switch data.(type) {
	case requestmodel_server_svc.ServerMessage:
		kafkaKey = "server message"
	case requestmodel_server_svc.FormCommand:
		kafkaKey = "forum command"
	case requestmodel_server_svc.ForumPost:
		kafkaKey = "forum post"
	}

	message := s.marshelStruct(data)

	fmt.Println("from kafka server message ", data)

	producer, err := sarama.NewSyncProducer([]string{s.config.KafkaPort}, config)
	if err != nil {
		return err
	}

	finalMessage := sarama.ProducerMessage{Topic: s.config.KafkaServerTopic, Key: sarama.StringEncoder(kafkaKey), Value: sarama.StringEncoder(message)}
	_, _, err = producer.SendMessage(&finalMessage)
	return err
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
	// log.Printf("[producer] partition id: %d; offset:%d, value: %v\n", partition, offset, msg)
	return err
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
}

func (s *serverServiceUseCase) SetDataInReddis() {
	// result := s.redisDB.Set(context.TODO(), "bismilla", "muhammad", time.Hour)
	// fmt.Println("===", result.Val())

	value := s.redisDB.Get(context.TODO(), "bismilla")
	fmt.Println("0000", value.Val())

}
