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

// Handler function for sending a friend request.
// @Summary Send Friend Request
// @Description Send a friend request.
// @Tags Friend
// @Accept json
// @Produce json
// @Param body body requestmodel_friend_svc.FriendRequest true "Request body for sending friend request"
// @Success 201 {object} responsemodel_friend_svc.Response "Friend request sent successfully"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/request [post]
// @Security UserAuthorization
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

// Handler function for getting user's friends.
// @Summary Get Friends
// @Description Retrieve user's friends.
// @Tags Friend
// @Accept json
// @Produce json
// @Param page query string false "Offset for paginated results"
// @Param limit query string false "Limit number of results"
// @Success 200 {object} responsemodel_friend_svc.Response "List of user's friends"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/friends [get]
// @Security UserAuthorization
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

// Handler function for getting received friend requests.
// @Summary Get Received Friend Requests
// @Description Retrieve received friend requests.
// @Tags Friend
// @Accept json
// @Produce json
// @Param page query string false "Offset for paginated results"
// @Param limit query string false "Limit number of results"
// @Success 200 {object} responsemodel_friend_svc.Response "List of received friend requests"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/received [get]
// @Security UserAuthorization
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

// Handler function for getting sent friend requests.
// @Summary Get Sent Friend Requests
// @Description Retrieve sent friend requests.
// @Tags Friend
// @Accept json
// @Produce json
// @Param page query string false "Offset for paginated results"
// @Param limit query string false "Limit number of results"
// @Success 200 {object} responsemodel_friend_svc.Response "List of sent friend requests"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/send [get]
// @Security UserAuthorization
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

// Handler function for getting blocked friend requests.
// @Summary Get Blocked Friend Requests
// @Description Retrieve blocked friend requests.
// @Tags Friend
// @Accept json
// @Produce json
// @Param page query string false "Offset for paginated results"
// @Param limit query string false "Limit number of results"
// @Success 200 {object} responsemodel_friend_svc.Response "List of blocked friend requests"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/block [get]
// @Security UserAuthorization
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

// Handler function for updating friendship status.
// @Summary Update Friendship Status
// @Description Update friendship status (e.g., restrict or unblock).
// @Tags Friend
// @Accept json
// @Produce json
// @Param body body requestmodel_friend_svc.FrendShipStatusUpdate true "Request body for updating friendship status"
// @Success 200 {object} responsemodel_friend_svc.Response "Friendship status updated successfully"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/restrict [patch]
// @Security UserAuthorization
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
		UserID:       ctx.Get("userID").(string),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "succesfully updated as "+ctx.QueryParam("action"), "", nil))
}

// websocket
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

// Handler function for getting chat messages.
// @Summary Get Chat Messages
// @Description Retrieve chat messages.
// @Tags Friend
// @Accept json
// @Produce json
// @Param page query string false "Offset for paginated results"
// @Param limit query string false "Limit number of results"
// @Param FriendID query string false "required friend id"
// @Success 200 {object} responsemodel_friend_svc.Response "List of chat messages"
// @Failure 400 {object} responsemodel_friend_svc.Response "Bad request"
// @Router /friend/chat/message [get]
// @Security UserAuthorization
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
// 			proxy_set_header Upgrade $http_upgrade;
// 			proxy_set_header Connection $connection_upgrade;
// 			proxy_connect_timeout 7d;
// 			proxy_send_timeout 7d;
// 			proxy_read_timeout 7d;
// 			#try_files $uri $uri/ =404;
// 	}
// }
