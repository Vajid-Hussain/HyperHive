package utils_server_service

import (
	"bytes"
	"fmt"
	"strconv"

	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func CreateSession(cfg config_server_service.S3Bucket) *session.Session {
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

func UploadImageToS3(file []byte, sess *session.Session) (string, error) {

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("hyper-hive-data"),
		Key:    aws.String("server images/" + fileName),
		Body:   aws.ReadSeekCloser(bytes.NewReader(file)),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println("err from s3 upload", err)
		return "", err
	}
	// fmt.Println("----", upload.Location)
	return upload.Location, nil

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
		return "", responsemodel_server_service.ErrPaginationWrongValue
	}

	return strconv.Itoa((offsetInt * limitInt) - limitInt), nil
}
