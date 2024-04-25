package handler_auth_svc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	requestmodel_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/Model/requestModel"
	response_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/Model/response"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
	"github.com/labstack/echo/v4"
)

type AuthHanlder struct {
	clind      pb.AuthServiceClient
	authCacing *helper_api_gateway.RedisCaching
}

func NewAuthHandler(clind pb.AuthServiceClient, authCacing *helper_api_gateway.RedisCaching) *AuthHanlder {
	return &AuthHanlder{
		clind:      clind,
		authCacing: authCacing,
	}
}

// @Summary User Signup
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body requestmodel_auth_svc.UserSignup true "User details for signup"
// @Success 201 {object} response_auth_svc.Response "User signup successful"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /signup [post]
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

// @Summary Resend verification email
// @Description Resend verification email to the user
// @Tags Authentication
// @Accept json
// @Produce json
// @Security UserConfirmToken
// @Param ConfirmToken header string true "Confirmation token"
// @Success 201 {object} response_auth_svc.Response "Email send successful"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /auth/verifyemailresend [post]
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

// Handler function for /sendotp endpoint.
// @Summary Send OTP
// @Description Sends OTP to the user's email.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body requestmodel_auth_svc.EmailReq true "Request body"
// @Success 201 {object} response_auth_svc.Response "Email send successful"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /sendotp [post]
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

// Handler function for /forgotpassword endpoint.
// @Summary Forgot Password
// @Description Sends a password reset email to the user's email.
// @Tags Authentication
// @Accept json
// @Produce json
// @Security UserConfirmToken
// @Param body body requestmodel_auth_svc.ForgotPassword true "Request body"
// @Success 201 {object} response_auth_svc.Response "Password reset email sent successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /forgotpassword [post]
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
	fmt.Println("==verrify email no email or token outside")

	if mail == "" || token == "" {
		fmt.Println("==verrify email no email or token ")
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", "no email or token"))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.clind.VerifyUser(context, &pb.UserVerifyRequest{
		Token: token,
		Email: mail,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusCreated, response_auth_svc.Responses(http.StatusCreated, "You are verified, now you can confirm your signup", "", nil))
}

// Handler function for /confirm signup endpoint.
// @Summary Confirm Signup
// @Description Confirm user signup with the provided token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Security UserConfirmToken
// @Success 200 {object} response_auth_svc.Response "Signup confirmed successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /confirm [post]
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

// Handler function for /login endpoint.
// @Summary User Login
// @Description Authenticate user and generate access token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body requestmodel_auth_svc.UserLogin true "User login details"
// @Success 200 {object} response_auth_svc.Response "Login successful"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /login [post]
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

// Handler function for /profile/profilephoto endpoint.
// @Summary Update Profile Photo
// @Description Update user's profile photo.
// @Tags Profile
// @Accept multipart/form-data
// @Produce json
// @Security UserAuthorization
// @Param ProfilePhoto formData file true "User's profile photo"
// @Success 200 {object} response_auth_svc.Response "Profile photo updated successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /profile/profilephoto [post]
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

	// if _, ok := validImageExtention[file.Header.Get("Content-Type")]; !ok {
	// 	return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", response_auth_svc.ErrUnsupportImageType.Error()))
	// }

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

// Handler function for /profile/coverphoto endpoint.
// @Summary Update Cover Photo
// @Description Update user's cover photo.
// @Tags Profile
// @Accept multipart/form-data
// @Produce json
// @Security UserAuthorization
// @Param CoverPhoto formData file true "User's cover photo"
// @Success 200 {object} response_auth_svc.Response "Cover photo updated successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /profile/coverphoto [post]
func (c *AuthHanlder) UpdateCoverPhoto(ctx echo.Context) error {
	var validImageExtention = map[string]struct{}{}

	validImageExtention["image/jpb"] = struct{}{}
	validImageExtention["image/png"] = struct{}{}
	validImageExtention["image/gif"] = struct{}{}

	file, err := ctx.FormFile("CoverPhoto")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, response_auth_svc.ErrNoImageInRequest.Error(), "", err.Error()))
	}

	// fmt.Println("==", file.Header, file.Size)

	if file.Size/(1024) > 1024 {
		return ctx.JSON(http.StatusRequestEntityTooLarge, response_auth_svc.Responses(http.StatusRequestEntityTooLarge, "", "", response_auth_svc.ErrImageOverSize.Error()))
	}

	// if _, ok := validImageExtention[file.Header.Get("Content-Type")]; !ok {
	// 	return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", response_auth_svc.ErrUnsupportImageType.Error()))
	// }

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

// Handler function for /profile/photoprofile endpoint.
// @Summary Delete Photo from User Profile
// @Description Delete a photo from the user's profile.
// @Tags Profile
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Param body body requestmodel_auth_svc.DeleteUserProfilePhotoType true "Request body for deleting user profile photo"
// @Success 200 {object} response_auth_svc.Response "Photo deleted successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /profile/photoprofile [delete]
func (c *AuthHanlder) DeletePhotFromUserProfile(ctx echo.Context) error {
	var req requestmodel_auth_svc.DeleteUserProfilePhotoType

	ctx.Bind(&req)
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := c.clind.DeletePhotoInProfile(context, &pb.DeletePhotoInProfileRequest{
		UserID: ctx.Get("userID").(string),
		Types:  req.Types,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", response_auth_svc.DeleteProfiesPhotos, nil))
}

// Handler function for /profile/status endpoint.
// @Summary Update Profile Status
// @Description Update user's profile status.
// @Tags Profile
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Param body body requestmodel_auth_svc.UserProfileStatus true "Request body for updating user profile status"
// @Success 200 {object} response_auth_svc.Response "Profile status updated successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /profile/status [post]
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
		fmt.Println(err, "--error in status update in profile with nill return from grpc")
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", "status succesfully updated", nil))
}

// Handler function for /profile/description endpoint.
// @Summary Update Profile Description
// @Description Update user's profile description.
// @Tags Profile
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Param body body requestmodel_auth_svc.UserProfileDescription true "Request body for updating user profile description"
// @Success 200 {object} response_auth_svc.Response "Profile description updated successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /profile/description [post]
func (c *AuthHanlder) UpdateProfileDescription(ctx echo.Context) error {
	var descriptionReq requestmodel_auth_svc.UserProfileDescription
	err := ctx.Bind(&descriptionReq)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, response_auth_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}
	descriptionReq.UserID = ctx.Get("userID").(string)

	// json_map := make(map[string]interface{})
	// json.NewDecoder(ctx.Request().Body).Decode(&json_map)
	// fmt.Println("****", json_map)

	validateError := helper_api_gateway.Validator(descriptionReq)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = c.clind.UpdateUserProfileDescription(context, &pb.UpdateUserProfileDescriptionRequest{
		UserID:      descriptionReq.UserID,
		Description: descriptionReq.Description,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	fmt.Println("===", err)
	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", "description succesfully updated", nil))
}

// Handler function for /profile endpoint to get user profile.
// @Summary Get User Profile
// @Description Retrieve user's profile information.
// @Tags Profile
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Success 200 {object} response_auth_svc.UserProfile "User profile information"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /profile/ [get]
func (c *AuthHanlder) GetUserProfile(ctx echo.Context) error {
	var userID string
	userID = ctx.Get("userID").(string)

	if ctx.Param("userID") != "" {
		userID = ctx.Param("userID")
	}

	// context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// result, err := c.clind.UserProfile(context, &pb.UserProfileRequest{
	// 	UserID: userID,
	// })

	result, err := c.authCacing.GetUserProfile(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

// Handler function for deleting a user account.
// @Summary Delete User Account
// @Description Delete user's account permanently.
// @Tags Profile
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Success 200 {object} response_auth_svc.Response "User account deleted successfully"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /account [delete]
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

// Handler function for searching users.
// @Summary Search Users
// @Description Search users based on specified criteria.
// @Tags Profile
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Param username query string false "Username for search"
// @Param limit query string false "Limit number of results"
// @Param page query string false "Offset for paginated results"
// @Success 200 {object} response_auth_svc.Response "User search response"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /users [get]
func (c *AuthHanlder) SerchUsers(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("==", ctx.QueryParam("username"), ctx.QueryParam("page"), ctx.QueryParam("limit"))

	result, err := c.clind.SerchUsers(context, &pb.SerchUsersRequest{
		UserName: ctx.QueryParam("username"),
		Limit:    ctx.QueryParam("limit"),
		Offset:   ctx.QueryParam("page"),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

// Handler function for creating access token by validating refresh token.
// @Summary Create Access Token by Validating Refresh Token
// @Description Create access token by validating refresh token.
// @Tags Token
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Param body body requestmodel_auth_svc.RefreshToken true "Request body for validating refresh token and creating access token"
// @Success 200 {object} response_auth_svc.Response "Token response"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /token/accesstoken [post]
func (c *AuthHanlder) CreateAcceesTokenByValidatingRefreshToken(ctx echo.Context) error {
	var req requestmodel_auth_svc.RefreshToken
	ctx.Bind(&req)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.clind.CreateAcceesTokenByValidatingRefreshToken(context, &pb.CreateAcceesTokenByValidatingRefreshTokenRequest{RefreshToken: req.RefreshToken})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}

// Handler function for separating user ID from access token.
// @Summary Separate User ID from Access Token
// @Description Separate user ID from access token.
// @Tags Token
// @Accept json
// @Produce json
// @Security UserAuthorization
// @Success 200 {object} response_auth_svc.Response "Token separation response"
// @Failure 400 {object} response_auth_svc.Response "Bad request"
// @Router /token [get]
func (c *AuthHanlder) SeperateUserIDFromAccessToken(ctx echo.Context) error {
	// var req requestmodel_auth_svc.AccessToken
	// ctx.Bind(&req)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.clind.SeperateUserIDFromAccessToken(context, &pb.SeperateUserIDFromAccessTokenRequest{AccessToken: ctx.Request().Header.Get("AccessToken")})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response_auth_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, response_auth_svc.Responses(http.StatusOK, "", result, nil))
}
