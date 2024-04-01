package handler_friend_svc

import (
	"context"
	"net/http"
	"time"

	requestmodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/requestModel"
	responsemodel_friend_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/pb"
	"github.com/labstack/echo/v4"
)

type FriendSvc struct {
	clind pb.FreindsServiceClient
}

func NewFriendSvc(clind pb.FreindsServiceClient) *FriendSvc {
	return &FriendSvc{clind: clind}
}

func (h *FriendSvc) FriendRequest(ctx echo.Context) error {
	var req requestmodel_friend_svc.FriendRequest
	ctx.Bind(&req)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := h.clind.FriendsRequest(context, &pb.FriendsRequestRequest{
		UserOne: req.FriendOne,
		UserTwo: req.FriendTwo,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responsemodel_friend_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, responsemodel_friend_svc.Responses(http.StatusOK, "Friend request send succesfully", "", nil))
}
