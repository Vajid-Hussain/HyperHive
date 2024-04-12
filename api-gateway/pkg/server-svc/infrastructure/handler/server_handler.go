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
