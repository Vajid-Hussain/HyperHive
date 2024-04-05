package handler_friend_svc

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	requestmodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/requestModel"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type Helper struct {
	RedisDb *redis.Client
	config  *config.Config
}

func NewHelper(connection *redis.Client, config *config.Config) *Helper {
	return &Helper{RedisDb: connection,
		config: config}
}

func (r *Helper) MessageProducer(message requestmodel_friend_svc.Message) error {
	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{r.config.KafkaPort}, configs)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{Topic: r.config.KafkaTopic, Key: sarama.StringEncoder("pearTopear"), Value: sarama.StringEncoder(message)}
}

func (r *Helper) HelperUserConnection(userID string, conn *websocket.Conn) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := r.RedisDb.Set(ctx, userID, conn.RemoteAddr().String(), 0).Result()
	if err != nil {
		fmt.Println("----err", err)
		return err
	}
	fmt.Println("--", result)

	fmt.Println("key---", r.RedisDb.Get(ctx, userID))
	return nil
}

func (r *Helper) SendMessage(message requestmodel_friend_svc.Message) error {
	fmt.Println(message)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := r.RedisDb.Get(ctx, message.SenderID).Result()
	if err != nil {
		fmt.Println("--", err)
		return err
	}

	// var dialer *websocket.Dialer
	// newConn, _, err := dialer.Dial("ws://"+result, nil)
	// if err != nil {
	// 	fmt.Println("err1", err)
	// }
	// defer newConn.Close()

	fmt.Println("ip", result)
	newConn, _, err := websocket.DefaultDialer.DialContext(ctx, "ws://"+result, nil)
	fmt.Println("err6", err)

	err = newConn.WriteMessage(websocket.TextMessage, []byte(message.Content))
	if err != nil {
		fmt.Println("err2", err)
	}

	fmt.Println("Message sent to user A:", message)

	return nil
}

// type websocketCon struct {
// 	Connection websocket.Conn `json:"webcocket"`
// }

// func (r *Helper) SaveCoonection(userID string, conn *websocket.Conn) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()

// 	// connJson := fmt.Sprintf("%v", conn)
// 	connJson, err := json.Marshal(conn)
// 	if err != nil {
// 		fmt.Println("Error marshaling websocket.Conn:", err)
// 		return err
// 	}
// 	fmt.Println(connJson)

// 	result, err := r.RedisDb.Set(ctx, userID, connJson, time.Hour).Result()
// 	if err != nil {
// 		fmt.Println("===", err)
// 		return err
// 	}
// 	fmt.Println("==", result, err)

// 	connData, err := r.RedisDb.Get(ctx, userID).Result()
// 	if err != nil {
// 		fmt.Println("**", err)
// 	}

// 	var connDataMap map[string]interface{}
// 	err = json.Unmarshal([]byte(connData), &connDataMap)
// 	if err != nil {
// 		fmt.Println("err1", err)
// 		return err
// 	}

// 	connBytes, err := json.Marshal(connDataMap["conn"])
// 	if err != nil {
// 		fmt.Println("err2", err)
// 		return err
// 	}

// 	var conns *websocket.Conn
// 	err = json.Unmarshal(connBytes, &conns)
// 	if err != nil {
// 		fmt.Println("err3", err)
// 		return err
// 	}

// 	fmt.Println("connfinal",conns)

// 	return nil
// }

// type connection struct{
// 	conn *websocket.Conn
// 	userID string
// }

// func (r *Helper) StoreConnection( userID string, WebsocketConn *websocket.Conn)error{
// 	NewConnection:= connection{conn: WebsocketConn,userID: userID}

// 	err:= r.RedisDb.Set(context.Background(), userID, NewConnection, 0).Err()
// 	if err!=nil{
// 		fmt.Println("err1", err)
// 	}

// 	value, err:= r.RedisDb.Get(context.Background(), userID).Result()
// 	if err!=nil{
// 		fmt.Println("Err2", err)
// 	}

// 	var getConn *connection
// 	err=json.Unmarshal([]byte(value), getConn)
// 	if err!=nil{
// 		fmt.Println("err3",err)
// 	}

// 	fmt.Println(getConn)
// 	return nil
// }
