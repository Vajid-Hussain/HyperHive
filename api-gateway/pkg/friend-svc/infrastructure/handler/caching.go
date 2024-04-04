package handler_friend_svc

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type Caching struct {
	RedisDb *redis.Client
}

func NewCaching(connection *redis.Client) *Caching {
	return &Caching{RedisDb: connection}
}

func (r *Caching) CachingUserConnection(userID string, conn *websocket.Conn) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// result, err := r.RedisDb.Ping(ctx).Result()
	// fmt.Println("===", result, err)

	connection, err := json.Marshal(WebSocketInfo{RemoteAddr: conn.RemoteAddr().String()})
	if err != nil {
		return err
	}

	result, err := r.RedisDb.Set(ctx, userID, connection, time.Hour*3).Result()
	if err != nil {
		fmt.Println("----err", err)
		return err
	}
	fmt.Println("--", result)

	fmt.Println("key---", r.RedisDb.Get(ctx, userID))
	return nil
}

func (r *Caching) SendMessage(message Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := r.RedisDb.Get(ctx, strconv.Itoa(message.RecipientID)).Result()
	if err != nil {
		fmt.Println("--", err)
		return err
	}

	var websocketInfo WebSocketInfo
	if err := json.Unmarshal([]byte(result), &websocketInfo); err != nil {
		fmt.Println("==", err)
		return err
	}

	wsURL := fmt.Sprintf("ws://%s", websocketInfo.RemoteAddr)
	fmt.Println("-===", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		fmt.Println("88", err)
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte(message.Content)); err != nil {
		fmt.Println("^^", err)
		return err
	}

	fmt.Println("Message sent to user A:", message)

	return nil
}
