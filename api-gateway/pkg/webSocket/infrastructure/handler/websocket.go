package handler_werbsocket_svc

import (
	"fmt"
	"net/http"
	"time"

	usecase_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/infrastructure/useCase"
	utils_websocket_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/webSocket/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WebsocketHandler struct {
	WebSocketUseCase *usecase_websocket_svc.WebSocketUseCase
}

func NewWebSocketHandler(usecase *usecase_websocket_svc.WebSocketUseCase) *WebsocketHandler {
	return &WebsocketHandler{WebSocketUseCase: usecase}
}

var User = make(map[string]*websocket.Conn)

var upgarde = websocket.Upgrader{
	HandshakeTimeout: 10 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

func (w WebsocketHandler) WebSocketConnection(ctx echo.Context) error {
	conn, err := upgarde.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils_websocket_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	defer delete(User, ctx.Get("userID").(string))
	defer conn.Close()
	User[ctx.Get("userID").(string)] = conn

	for {
		fmt.Println("==loop start ", User)
		_, msg, err := conn.ReadMessage()
		w.WebSocketUseCase.PreprocessOfSendingMessage(msg, User, ctx.Get("userID").(string), err)
	}
}
