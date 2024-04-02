package response_auth_svc

import (
	"errors"
)

var (
	ErrImageOverSize      = errors.New("upload image less than 1 MB")
	ErrUnsupportImageType = errors.New("image type not supported, try JPG, PNG, and GIF formats")
	ErrNoImageInRequest   = errors.New("kindly attach your cover photo")
)

var (
	EmailSendSuccessfully = "Email sent successfully"
	DeleteProfiesPhotos   = "succesfully deleted"
)
