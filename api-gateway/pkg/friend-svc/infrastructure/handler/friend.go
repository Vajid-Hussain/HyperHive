package handler_friend_svc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	requestmodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/requestModel"
	responsemodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/pb"
	"github.com/labstack/echo/v4"
)

type FriendSvc struct {
	clind pb.FriendServiceClient
}

func NewFriendSvc(clind pb.FriendServiceClient) *FriendSvc {
	return &FriendSvc{clind: clind}
}

func (h *FriendSvc) FriendRequest(ctx echo.Context) error {
	var req requestmodel_friend_svc.FriendRequest
	ctx.Bind(&req)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("====", req, ctx.Get("userID").(string))

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
	
	result, err := h.clind.FriendList(context, &pb.FriendListRequest{
		UserID: ctx.Get("userID").(string),
		OffSet: ctx.QueryParam("page"),
		Limit:  ctx.QueryParam("limit"),
		Status: ctx.QueryParam("status"),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "", result, nil))
}
