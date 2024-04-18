package helper_api_gateway

import (
	"bytes"
	"fmt"

	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func Validator(request any) []string {
	errResponse := []string{}

	var validate = validator.New()
	errs := validate.Struct(request)

	if errs == nil {
		return nil
	}

	for _, err := range errs.(validator.ValidationErrors) {
		errResponse = append(errResponse, fmt.Sprintf("[%s] : '%v' | become '%s' %s", err.Field(), err.Value(), err.Tag(), err.Param()))
	}

	return errResponse
}

func CreateSession(cfg config.S3Bucket) *session.Session {
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
		Key:    aws.String("chat media/" + fileName),
		Body:   aws.ReadSeekCloser(bytes.NewReader(file)),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println("err from s3 upload", err)
		return "", err
	}
	return upload.Location, nil
}
