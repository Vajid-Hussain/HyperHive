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
	)

	err := ctx.Bind(&UserDetails)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	fmt.Println("--", UserDetails)

	validateError := helper_api_gateway.Validator(UserDetails)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

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

func (c *AuthHanlder) ReSendVerificationEmail(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := c.clind.ReSendVerificationEmail(context, &pb.ReSendVerificationEmailRequest{
		Token: ctx.Request().Header.Get("ConfirmToken"),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, response_auth_svc.EmailSendSuccessfully, token, nil))
}

func (c *AuthHanlder) SendOtp(ctx echo.Context) error {
	var req requestmodel_auth_svc.EmailReq

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	token, err := c.clind.SendOtp(context, &pb.SendOtpRequest{
		Emain: req.Email,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, response_auth_svc.EmailSendSuccessfully, token, nil))
}

func (c *AuthHanlder) ForgotPassword(ctx echo.Context) error {
	var req requestmodel_auth_svc.ForgotPassword
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = c.clind.ForgotPassword(context, &pb.ForgotPasswordRequest{
		Otp:      req.Otp,
		Password: req.Password,
		Token:    ctx.Request().Header.Get("ConfirmToken"),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, "password succesfully changed", "", nil))
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

func (c *AuthHanlder) UserLogin(ctx echo.Context) error {
	var loginReq requestmodel_auth_svc.UserLogin

	ctx.Bind(&loginReq)
	validateError := helper_api_gateway.Validator(loginReq)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := c.clind.UserLogin(context, &pb.UserLoginRequest{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *AuthHanlder) UpdateProfilePhoto(ctx echo.Context) error {
	var validImageExtention = map[string]struct{}{}

	validImageExtention["image/jpb"] = struct{}{}
	validImageExtention["image/png"] = struct{}{}
	validImageExtention["image/gif"] = struct{}{}

	file, err := ctx.FormFile("ProfilePhoto")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, response_auth_svc.ErrNoImageInRequest.Error(), "", err.Error()))
	}

	if file.Size/(1024) > 1024 {
		return ctx.JSON(http.StatusRequestEntityTooLarge, response_auth_svc.Responses(http.StatusRequestEntityTooLarge, "", "", response_auth_svc.ErrImageOverSize.Error()))
	}

	if _, ok := validImageExtention[file.Header.Get("Content-Type")]; !ok {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", response_auth_svc.ErrUnsupportImageType.Error()))
	}

	image, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	buffer := make([]byte, file.Size)
	image.Read(buffer)

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := c.clind.UpdateProfilePhoto(context, &pb.UpdateprofilePhotoRequest{
		Image:  buffer,
		UserID: ctx.Get("userID").(string),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *AuthHanlder) UpdateCoverPhoto(ctx echo.Context) error {
	var validImageExtention = map[string]struct{}{}

	validImageExtention["image/jpb"] = struct{}{}
	validImageExtention["image/png"] = struct{}{}
	validImageExtention["image/gif"] = struct{}{}

	file, err := ctx.FormFile("CoverPhoto")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, response_auth_svc.ErrNoImageInRequest.Error(), "", err.Error()))
	}

	if file.Size/(1024) > 1024 {
		return ctx.JSON(http.StatusRequestEntityTooLarge, response_auth_svc.Responses(http.StatusRequestEntityTooLarge, "", "", response_auth_svc.ErrImageOverSize.Error()))
	}

	if _, ok := validImageExtention[file.Header.Get("Content-Type")]; !ok {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", response_auth_svc.ErrUnsupportImageType.Error()))
	}

	image, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	buffer := make([]byte, file.Size)
	image.Read(buffer)

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := c.clind.UpdateCoverPhoto(context, &pb.UpdateCoverPhotoRequest{
		CoverPhoto: buffer,
		UserID:     ctx.Get("userID").(string),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *AuthHanlder) UpdateProfileStatus(ctx echo.Context) error {
	var statusReq requestmodel_auth_svc.UserProfileStatus
	ctx.Bind(&statusReq)
	statusReq.UserID = ctx.Get("userID").(string)

	fmt.Println("--", statusReq)
	validateError := helper_api_gateway.Validator(statusReq)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.clind.UpdateUserProfileStatus(context, &pb.UpdateUserProfileStatusRequest{
		UserID:   statusReq.UserID,
		Status:   statusReq.Status,
		Duration: statusReq.Duration,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", "status succesfully updated", nil))
}

func (c *AuthHanlder) UpdateProfileDescription(ctx echo.Context) error {
	var descriptionReq requestmodel_auth_svc.UserProfileDescription
	ctx.Bind(&descriptionReq)
	descriptionReq.UserID = ctx.Get("userID").(string)

	fmt.Println("--", descriptionReq)
	validateError := helper_api_gateway.Validator(descriptionReq)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := c.clind.UpdateUserProfileDescription(context, &pb.UpdateUserProfileDescriptionRequest{
		UserID:      descriptionReq.UserID,
		Description: descriptionReq.Description,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", "description succesfully updated", nil))
}

func (c *AuthHanlder) GetUserProfile(ctx echo.Context) error {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.clind.UserProfile(context, &pb.UserProfileRequest{
		UserID: ctx.Get("userID").(string),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *AuthHanlder) DeleteUserAcoount(ctx echo.Context) error {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.clind.DeleteAccount(context, &pb.DeleteAccountRequest{
		UserID: ctx.Get("userID").(string),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "Succesfully Deleted", result, nil))
}
