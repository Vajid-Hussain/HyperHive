package response_auth_svc

import "errors"

var (
	ErrImageOverSize      = errors.New("image size more than one 1MB, keep try with less than a MB")
	ErrUnsupportImageType = errors.New("image type not supported, only JPG, PNG, and GIF formats are allowed")
	ErrNoImageInRequest =errors.New("kindly attach your cover photo")
)
