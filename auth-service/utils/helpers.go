package utils_auth_server

import (
	"bytes"
	"errors"
	"fmt"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		Bucket: aws.String("hiper-hive-data"),
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
