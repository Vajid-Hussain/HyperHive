package resonsemodel_server_svc

import "errors"

var (
	ErrImageOverSize              = errors.New("upload image less than 1 MB")
	ErrUnsupportImageType         = errors.New("image type not supported, try JPG, PNG, and GIF formats")
	ErrNoImageInRequest           = errors.New("kindly attach your cover photo")
	ErrForumUnexpectedType        = errors.New("unexpected forum request type")
	ErrForumPostUnexpectedContent = errors.New("unexpected forum post content type")
	ErrServerMessageType          = errors.New("message type must be file or test")
	ErrNoPostIDINQueryParams      = errors.New("no postid in qeury param")
	ErrUserMessageSupportType =errors.New("user message only support file and text")
)

var (
	ServerImageUpdateSuccesFully = " sussefully updated "
	ServerDescriptionUpdate      = "succesfully description updated"
)
