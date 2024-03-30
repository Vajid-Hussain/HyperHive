package usecasel_auth_server

import (
	"errors"
	"fmt"
	"regexp"

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

	// var chanCoverPhoto = make(chan string)
	// fmt.Println(userDetails.ConfirmPassword, userDetails.Name, userDetails.UserName)

	// r, _ := regexp.Compile("p([a-zA-Z0-9!@#$%^&*]+)ch")
	// r := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{5,})`)
	// if !r.MatchString(userDetails.Password) {
	// 	return nil, responsemodel_auth_server.ErrRegexNotMatch
	// }

	if !hasLowercase.MatchString(userDetails.Password) || !hasUppercase.MatchString(userDetails.Password) || !hasDigit.MatchString(userDetails.Password) || !hasMinimumLength.MatchString(userDetails.Password) || !hasSymbol.MatchString(userDetails.Password) {
		return nil, responsemodel_auth_server.ErrRegexNotMatch
	}

	userDetails.Password = utils_auth_server.HashPassword(userDetails.Password)

	// s3Session := utils_auth_server.CreateSession(d.s3)

	// if len(userDetails.ProfilePhoto) > 1 {
	// 	fmt.Println("profile proto is sending to s3")
	// 	go utils_auth_server.UploadImageToS3(userDetails.ProfilePhoto, s3Session, chanProfilePhoto)
	// }

	// if len(userDetails.CoverPhoto) > 1 {
	// 	fmt.Println("CoverPhoto is sending to s3")
	// go	utils_auth_server.UploadImageToS3(userDetails.CoverPhoto, s3Session, chanCoverPhoto)
	// }

	// if len(userDetails.ProfilePhoto) > 1 {
	// 	fmt.Println("profile  url fetch by chan to s3")
	// 	userDetails.ProfilePhotoUrl = <-chanProfilePhoto
	// }

	// if len(userDetails.CoverPhoto) > 1 {
	// 	fmt.Println("CoverPhoto url fetch by chan to s3")
	// 	userDetails.CoverPhotoUrl = <-chanCoverPhoto
	// }

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
		fmt.Println("error at accestoken", err)
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
