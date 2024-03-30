package usecasel_auth_server

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	requestmodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/requestModel"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
	interface_repo_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/repository/interface"
	interface_usecase_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/usecase/interface"
	utils_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/utils"
)

type UserUseCase struct {
	userRepo       interface_repo_auth_server.IUserRepository
	s3             configl_auth_server.S3Bucket
	mailConstrains configl_auth_server.Mail
	tokenSecret    configl_auth_server.Token
}

func NewUserUseCase(repo interface_repo_auth_server.IUserRepository, s3 configl_auth_server.S3Bucket, mailConstrains configl_auth_server.Mail, tokenSecret configl_auth_server.Token) interface_usecase_auth_server.IUserUseCase {
	return &UserUseCase{userRepo: repo,
		s3:             s3,
		mailConstrains: mailConstrains,
		tokenSecret:    tokenSecret}
}

func (d *UserUseCase) Signup(userDetails requestmodel_auth_server.UserSignup) (*responsemodel_auth_server.UserSignup, error) {
	hasLowercase := regexp.MustCompile(`[a-z]`)
	hasUppercase := regexp.MustCompile(`[A-Z]`)
	hasDigit := regexp.MustCompile(`[0-9]`)
	hasMinimumLength := regexp.MustCompile(`.{5,}`)
	hasSymbol := regexp.MustCompile(`[!@#$%^&*]`)

	if !hasLowercase.MatchString(userDetails.Password) || !hasUppercase.MatchString(userDetails.Password) || !hasDigit.MatchString(userDetails.Password) || !hasMinimumLength.MatchString(userDetails.Password) || !hasSymbol.MatchString(userDetails.Password) {
		return nil, responsemodel_auth_server.ErrRegexNotMatch
	}

	userDetails.Password = utils_auth_server.HashPassword(userDetails.Password)

	count, err := d.userRepo.EmailIsExist(userDetails.Email)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, responsemodel_auth_server.ErrEmailExists
	}

	count, err = d.userRepo.UserNameIsExist(userDetails.UserName)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, responsemodel_auth_server.ErrUsernameTaken
	}

	userRes, err := d.userRepo.Signup(userDetails)
	if err != nil {
		return nil, err
	}

	verificationToken, err := utils_auth_server.TemperveryTokenForUserAuthenticaiton(d.tokenSecret.TemperveryKey, userRes.ID)
	if err != nil {
		return nil, err
	}

	utils_auth_server.SendVerificationEmail(userRes.Email, verificationToken, d.mailConstrains)

	userRes.TemperveryToken = verificationToken
	return userRes, nil
}

func (d *UserUseCase) VerifyUserSignup(email, token string) error {
	// fmt.Println("--", email, token)

	userID, err := utils_auth_server.FetchUserIDFromToken(token, d.tokenSecret.TemperveryKey)
	if err != nil {
		return err
	}

	err = d.userRepo.VerifyUserSignup(userID, email)
	if err != nil {
		return err
	}

	return nil
}

func (d *UserUseCase) ReSendVerificationMail(token string) (string, error) {
	var verificationToken string
	userID, err := utils_auth_server.FetchUserIDFromTokenNoWorryOnExpire(token, d.tokenSecret.TemperveryKey)
	if err != nil {
		return "", err
	}

	count, err := d.userRepo.IsUserIDExist(userID)
	if err != nil {
		return "", err
	}

	if count == 1 {
		verificationToken, err = utils_auth_server.TemperveryTokenForUserAuthenticaiton(d.tokenSecret.TemperveryKey, userID)
		if err != nil {
			return "", err
		}

		email, err := d.userRepo.FetchMailUsingUserID(userID)
		if err != nil {
			return "", err
		}

		utils_auth_server.SendVerificationEmail(email, verificationToken, d.mailConstrains)
	}

	return verificationToken, nil
}

func (d *UserUseCase) ConfirmSignup(token string) (*responsemodel_auth_server.AuthenticationResponse, error) {
	var verifyRes responsemodel_auth_server.AuthenticationResponse

	userID, err := utils_auth_server.FetchUserIDFromToken(token, d.tokenSecret.TemperveryKey)
	if err != nil {
		return nil, err
	}

	exist, err := d.userRepo.ConfirmSignup(userID)
	if err != nil {
		return nil, err
	}

	if exist == 0 {
		return nil, errors.New("confirm email first then next")
	}

	verifyRes.AccesToken, err = utils_auth_server.GenerateAcessToken(d.tokenSecret.UserSecurityKey, userID)
	if err != nil {
		return nil, err
	}

	verifyRes.RefreshToken, err = utils_auth_server.GenerateRefreshToken(d.tokenSecret.UserSecurityKey)
	if err != nil {
		return nil, err
	}

	return &verifyRes, nil
}

func (d *UserUseCase) UserLogin(email, password string) (*responsemodel_auth_server.AuthenticationResponse, error) {
	var loginRes responsemodel_auth_server.AuthenticationResponse

	storedPassword, err := d.userRepo.GetUserPasswordUsingEmail(email)
	if err != nil {
		return nil, err
	}

	if storedPassword == "" {
		return nil, responsemodel_auth_server.ErrLoginNoActiveUser
	}

	err = utils_auth_server.CompairPassword(storedPassword, password)
	if err != nil {
		return nil, err
	}

	userID, err := d.userRepo.FetchUserIDUsingMail(email)
	if err != nil {
		return nil, err
	}

	loginRes.AccesToken, err = utils_auth_server.GenerateAcessToken(d.tokenSecret.UserSecurityKey, userID)
	if err != nil {
		return nil, err
	}

	loginRes.RefreshToken, err = utils_auth_server.GenerateRefreshToken(d.tokenSecret.UserSecurityKey)
	if err != nil {
		return nil, err
	}

	return &loginRes, nil
}

func (u *UserUseCase) VerifyUserToken(accessToken, refreshToken string) (string, error) {
	fmt.Println("user middlewiere")
	id, err := utils_auth_server.VerifyAcessToken(accessToken, u.tokenSecret.UserSecurityKey)
	if err != nil {
		return "", err
	}

	err = utils_auth_server.VerifyRefreshToken(refreshToken, u.tokenSecret.UserSecurityKey)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (u *UserUseCase) UpdateProfilePhoto(userID string, image []byte) (url string, err error) {
	var chanProfilePhoto = make(chan string)
	s3Session := utils_auth_server.CreateSession(u.s3)

	if len(image) > 1 {
		fmt.Println("profile proto is sending to s3")
		go utils_auth_server.UploadImageToS3(image, s3Session, chanProfilePhoto)
	}

	if len(image) > 1 {
		fmt.Println("profile  url fetch by chan to s3")
		url = <-chanProfilePhoto
	}

	err = u.userRepo.UpdateUserProfilePhoto(userID, url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (u *UserUseCase) UpdateCoverPhoto(userID string, image []byte) (url string, err error) {
	var chanCoverPhoto = make(chan string)
	s3Session := utils_auth_server.CreateSession(u.s3)

	if len(image) > 1 {
		go utils_auth_server.UploadImageToS3(image, s3Session, chanCoverPhoto)
	}

	if len(image) > 1 {
		url = <-chanCoverPhoto
	}

	err = u.userRepo.UpdateCoverPhoto(userID, url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (u *UserUseCase) UpdateStatusOfUser(status requestmodel_auth_server.UserProfileStatus, expire float32) error {
	status.Expire = time.Now().Add(time.Hour * time.Duration(expire)).Format("2006-01-02 15:04:05")
	fmt.Println("-----", time.Now(), status.Expire)

	if expire > 6 {
		return responsemodel_auth_server.ErrStatuTimeLongExpireTime
	}

	err := u.userRepo.UpdateOrCreateUserStatus(status)
	if err != nil {
		return nil
	}

	return nil
}

func (u *UserUseCase) UpdateDescriptionOfUser(userID, description string) error {
	err := u.userRepo.UpdateOrCreateUserDescription(userID, description)
	if err != nil {
		return nil
	}

	return nil
}

func (u *UserUseCase) GetUserProfile(userID string) (*responsemodel_auth_server.UserProfile, error) {
	profile, err := u.userRepo.GetUserProfile(userID)
	if err != nil {
		return nil, err
	}

	if profile.StatusExpireTime.Before(time.Now()) {
		profile.Status = ""
	}

	return profile, nil
}

func (u *UserUseCase) DeleteAccount(userID string) error {
	return u.userRepo.DeleteUserAcoount(userID)
}
