package usecasel_auth_server

import (
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
	// var chanProfilePhoto = make(chan string)

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
	// 	userDetails.CoverPhotoUrl = <-chanProfilePhoto
	// }

	userRes, err := d.userRepo.Signup(userDetails)
	if err != nil {
		return nil, err
	}

	verificationToken, err := utils_auth_server.TemperveryTokenForUserAuthenticaiton(d.tokenSecret.TemperveryKey, userRes.Email)
	if err != nil {
		return nil, err
	}

	utils_auth_server.SendVerificationEmail(userRes.Email, verificationToken, d.mailConstrains)

	return userRes, nil
}
