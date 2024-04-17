package resonsemodel_server_svc

import "errors"

var (
	ErrImageOverSize      = errors.New("upload image less than 1 MB")
	ErrUnsupportImageType = errors.New("image type not supported, try JPG, PNG, and GIF formats")
	ErrNoImageInRequest   = errors.New("kindly attach your cover photo")
)

var (
	ServerImageUpdateSuccesFully = " sussefully updated "
	ServerDescriptionUpdate      = "succesfully description updated"
)
