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

type AuthHanlder struct {
	clind pb.AuthServiceClient
}

func NewAuthHandler(clind pb.AuthServiceClient) *AuthHanlder {
	return &AuthHanlder{clind: clind}
}

func (c AuthHanlder) Signup(ctx echo.Context) error {
	var (
		UserDetails requestmodel_auth_svc.UserSignup
		// validImageExtention = map[string]struct{}{}
	)
	// validImageExtention["image/jpb"] = struct{}{}
	// validImageExtention["image/png"] = struct{}{}
	// validImageExtention["image/gif"] = struct{}{}

	err := ctx.Bind(&UserDetails)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(UserDetails)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	// file, err := ctx.FormFile("ProfilePhoto")
	// if err != nil {
	// 	// return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "kindly attach your profile photo", "", err.Error()))
	// }

	// if file.Size/(1024) > 1024 {
	// 	return ctx.JSON(http.StatusRequestEntityTooLarge, response_auth_svc.Responses(http.StatusRequestEntityTooLarge, "", "", "image size more than one 1MB, keep try with less than a MB"))
	// }

	// if _, ok := validImageExtention[file.Header.Get("Content-Type")]; !ok {
	// 	return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", "Image type not supported, only JPG, PNG, and GIF formats are allowed."))
	// }

	// image, err := file.Open()
	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	// }

	// buffer := make([]byte, file.Size)
	// image.Read(buffer)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.clind.Signup(context, &pb.SignupRequest{
		UserName:        UserDetails.UserName,
		Name:            UserDetails.Name,
		Email:           UserDetails.Email,
		Password:        UserDetails.Password,
		ConfirmPassword: UserDetails.ConfirmPassword,
	})

	fmt.Println("--", result)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, "", result, nil))
}

func (c *AuthHanlder) MailVerificationCallback(ctx echo.Context) error {
	mail := ctx.QueryParam("email")
	token := ctx.QueryParam("token")

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.clind.VerifyUser(context, &pb.UserVerifyRequest{
		Token: token,
		Email: mail,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()+" no user match "))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, "You are verified, now you can confirm your signup", "", nil))
}

func (c *AuthHanlder) ConfirmSignup(ctx echo.Context) error {
	token := ctx.Request().Header.Get("ConfirmToken")

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := c.clind.ConfirmSignup(context, &pb.ConfirmSignupRequest{
		TemperveryToken: token,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, "", result, nil))
}

