package handler_auth_svc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	requestmodel_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/Model/requestModel"
	response_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/Model/response"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
	"github.com/labstack/echo/v4"
)

type AdminAuthHanlder struct {
	clind pb.AuthServiceClient
}

func NewAdminAuthHandler(clind pb.AuthServiceClient) *AdminAuthHanlder {
	return &AdminAuthHanlder{clind: clind}
}

func (c *AdminAuthHanlder) AdminLogin(ctx echo.Context) error {
	var loginReq requestmodel_auth_svc.AdminLogin
	ctx.Bind(&loginReq)
	validateError := helper_api_gateway.Validator(loginReq)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	fmt.Println("--", loginReq)

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := c.clind.AdminLogin(context, &pb.AdminLoginRequest{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *AdminAuthHanlder) BlockUserAccount(ctx echo.Context) error {
	var blockRequest requestmodel_auth_svc.UserIDReq
	ctx.Bind(&blockRequest)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.clind.BlockUse(context, &pb.BlockUseRequest{
		UserID:  blockRequest.UserID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "Succesfully blocked", result, nil))
}

func (c *AdminAuthHanlder) UnBlockUserAccount(ctx echo.Context) error {
	var unblockRequest requestmodel_auth_svc.UserIDReq
	ctx.Bind(&unblockRequest)
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.clind.UnBlockUser(context, &pb.UnBlockUserRequest{
		UserID: unblockRequest.UserID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "Succesfully UnBlocked", result, nil))
}
