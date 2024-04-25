package utils_auth_server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	responsemodel_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/infrastructure/model/responseModel"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "problem when hashing ")
	}
	return string(HashedPassword)
}

func CompairPassword(hashedPassword string, plainPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("passwords does not match")
	}

	return nil
}

func CreateSession(cfg configl_auth_server.S3Bucket) *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
			Endpoint: aws.String(""),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}

func UploadImageToS3(file []byte, sess *session.Session, ch chan string) {

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("hyper-hive-data"),
		Key:    aws.String("profile images/" + fileName),
		Body:   aws.ReadSeekCloser(bytes.NewReader(file)),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println("err from s3 upload", err)
		return
	}
	// fmt.Println("----", upload.Location)
	ch <- upload.Location
	close(ch)
}

func TemperveryTokenForUserAuthenticaiton(securityKey string, email string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp":   time.Now().Unix() + 3600,
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

func GenerateAcessToken(securityKey string, id string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 36000000,
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

func GenerateRefreshToken(securityKey, userID string) (string, error) {
	key := []byte(securityKey)
	clamis := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Unix() + 36000000,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("making refresh token lead to error")
	}

	return signedToken, nil
}

func VerifyAcessToken(token string, secretkey string) (string, error) {

	key := []byte(secretkey)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if parsedToken==nil{
		return "", errors.New("invalid access token")
	}

	if err != nil {
		return "", err
	}

	if len(parsedToken.Header) == 0 {
		return "", errors.New("token tamberd include header")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(string)

	if !ok {
		return "", errors.New("id is not in accessToken. access denied")
	}
	return id, nil
}

func VerifyRefreshToken(token string, securityKey string) error {
	key := []byte(securityKey)

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return errors.New(" token tamperd or expired")
	}

	return nil
}

func GettingIDClimeFromToken(token, secret string) (string, error) {
	key := []byte(secret)
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if parsedToken==nil{
		return "", errors.New("invalid access token")
	}

	if err != nil {
		return "", err
	}

	if len(parsedToken.Header) == 0 {
		return "", errors.New("token tamberd include header")
	}


	claims := parsedToken.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(string)

	if !ok {
		return "", errors.New("id is not in accessToken. access denied")
	}
	return id, nil
}

func FetchUserIDFromToken(tokenString string, secretkey string) (string, error) {

	secret := []byte(secretkey)

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		return "", errors.New("wrong token or expired")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse claims")
	}

	emailClaim, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email claim not found or not a string")
	}

	return emailClaim, nil
}

func FetchUserIDFromTokenNoWorryOnExpire(tokenString string, secretkey string) (string, error) {
	if tokenString==""{
		return "", errors.New("attach token string with request")
	}

	secret := []byte(secretkey)

	parsedToken, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// if err != nil || !parsedToken.Valid {
	// 	return "", errors.New("wrong token or expired")
	// }
	fmt.Println(parsedToken)
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse claims")
	}

	emailClaim, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("email claim not found or not a string")
	}

	return emailClaim, nil
}

type Claims struct {
	Email string `json:"email"`
}

func ExtractEmailFromToken(tokenString string) (string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid JWT token format")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("error decoding token payload: %v", err)
	}

	var claims Claims
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return "", fmt.Errorf("error decoding claims: %v", err)
	}

	return claims.Email, nil
}

func Pagination(limit, offset string) (string, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return "", err
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "", err
	}

	if limitInt < 1 || offsetInt < 1 {
		return "", responsemodel_auth_server.ErrPaginationWrongValue
	}

	return strconv.Itoa((offsetInt * limitInt) - limitInt), nil
}
