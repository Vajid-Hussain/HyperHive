package handler_friend_svc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	requestmodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/requestModel"
	responsemodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/pb"
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type FriendSvc struct {
	clind  pb.FriendServiceClient
	helper *Helper
}

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var User = make(map[string]*websocket.Conn)

func NewFriendSvc(clind pb.FriendServiceClient, helper *Helper) *FriendSvc {
	return &FriendSvc{
		clind:  clind,
		helper: helper,
	}
}

func (h *FriendSvc) FriendRequest(ctx echo.Context) error {
	var req requestmodel_friend_svc.FriendRequest
	ctx.Bind(&req)

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := h.clind.FriendRequest(context, &pb.FriendRequestRequest{
		UserID:   ctx.Get("userID").(string),
		FriendID: req.FriendID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "Friend request send succesfully", result, nil))
}

func (h *FriendSvc) GetFriends(ctx echo.Context) error {

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := h.clind.FriendList(context, &pb.FriendListRequest{Friend: &pb.GetPendingListRequestModel{
		UserID: ctx.Get("userID").(string),
		OffSet: ctx.QueryParam("page"),
		Limit:  ctx.QueryParam("limit"),
	}})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) GetReceivedFriendRequest(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetReceivedFriendRequest(context, &pb.GetReceivedFriendRequestRequest{
		Received: &pb.GetPendingListRequestModel{
			UserID: ctx.Get("userID").(string),
			OffSet: ctx.QueryParam("page"),
			Limit:  ctx.QueryParam("limit"),
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) GetSendFriendRequest(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetSendFriendRequest(context, &pb.GetSendFriendRequestRequest{
		Send: &pb.GetPendingListRequestModel{
			UserID: ctx.Get("userID").(string),
			OffSet: ctx.QueryParam("page"),
			Limit:  ctx.QueryParam("limit"),
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) GetBlockFriendRequest(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetBlockFriendRequest(context, &pb.GetBlockFriendRequestRequest{
		Block: &pb.GetPendingListRequestModel{
			UserID: ctx.Get("userID").(string),
			OffSet: ctx.QueryParam("page"),
			Limit:  ctx.QueryParam("limit"),
		},
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

func (h *FriendSvc) UpdateFriendshipStatus(ctx echo.Context) error {
	var req requestmodel_friend_svc.FrendShipStatusUpdate
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := h.clind.UpdateFriendshipStatus(context, &pb.UpdateFriendshipStatusRequest{

		FriendShipID: req.FrendShipID,
		Status:       ctx.QueryParam("action"),
		UserID: ctx.Get("userID").(string),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "succesfully updated as "+ctx.QueryParam("action"), "", nil))
}

func (h *FriendSvc) FriendMessage(ctx echo.Context) error {
	fmt.Println("message called")

	conn, err := upgrade.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	defer delete(User, ctx.Get("userID").(string))
	defer conn.Close()

	User[ctx.Get("userID").(string)] = conn

	for {
		fmt.Println("loop starts", ctx.Get("userID"), User)

		_, msg, err := conn.ReadMessage()
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
		}

		h.helper.SendMessageToUser(User, msg, ctx.Get("userID").(string))
	}
}

func (h *FriendSvc) GetChat(ctx echo.Context) error {
	var chatRequet requestmodel_friend_svc.ChatRequest
	ctx.Bind(&chatRequet)
	validateError := helper_api_gateway.Validator(chatRequet)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := h.clind.GetFriendChat(context, &pb.GetFriendChatRequest{
		UserID:   ctx.Get("userID").(string),
		FriendID: chatRequet.FriendID,
		OffSet:   chatRequet.Offset,
		Limit:    chatRequet.Limit,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}

//nginx config

// 	#root /var/www/your_domain/html;
//    # index index.html index.htm index.nginx-debian.html;

// map $http_upgrade $connection_upgrade {
// 	default upgrade;
// 	 ''      close;
// }
// server {
// 	server_name hyperhive.vajid.tech www.hyperhive.vajid.tech;

// 	location / {
// 			proxy_pass http://localhost:9000;
// 	proxy_set_header Upgrade $http_upgrade;
// 	proxy_set_header Connection $connection_upgrade;
// 			#try_files $uri $uri/ =404;
// 	}
// }
