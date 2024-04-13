package handler_server_svc

import (
	"context"
	"net/http"
	"time"

	requestmodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/requestModel"
	resonsemodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/resonseModel"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	"github.com/labstack/echo/v4"
)

type ServerService struct {
	Clind pb.ServerClient
}

func NewServerService(clind pb.ServerClient) *ServerService {
	return &ServerService{Clind: clind}
}

func (c *ServerService) CreateServer(ctx echo.Context) error {
	var serverReq requestmodel_server_svc.Server

	err := ctx.Bind(&serverReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.CreateServer(context, &pb.CreateServerRequest{ServerName: serverReq.Name, UserID: ctx.Get("userID").(string)})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "server create succesfully", result, nil))
}

func (c *ServerService) CreateCategory(ctx echo.Context) error {
	var req requestmodel_server_svc.CreateCategory
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = c.Clind.CreateCategory(context, &pb.CreateCategoryRequest{
		UserID:       ctx.Get("userID").(string),
		ServerID:     req.ServerID,
		CategoryName: req.CategoryName,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "category succesfully created", "", nil))
}

func (c *ServerService) CreateChannel(ctx echo.Context) error {
	var req requestmodel_server_svc.CreateChannel
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = c.Clind.CreateChannel(context, &pb.CreateChannelRequest{
		UserID:      ctx.Get("userID").(string),
		ServerID:    req.ServerID,
		CategoryID:  req.CategoryID,
		ChannelName: req.ChannelName,
		ChannelType: req.Type,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "channel succesfully created", "", nil))
}

func (c *ServerService) JoinToServer(ctx echo.Context) error {
	var req requestmodel_server_svc.JoinToServer
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = c.Clind.JoinToServer(context, &pb.JoinToServerRequest{
		UserID:   ctx.Get("userID").(string),
		ServerID: req.ServerID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "Now you are a member", "", nil))
}

func (c *ServerService) GetCategoryOfServer(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerReq
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.GetCategoryOfServer(context, &pb.GetCategoryOfServerRequest{ServerID: req.ServerID})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) GetChannelsOfServer(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerReq
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.GetChannelsOfServer(context, &pb.GetChannelsOfServerRequest{ServerID: req.ServerID})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}
