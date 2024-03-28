package requestmodel_auth_svc

import "mime/multipart"

type UserSignup struct {
	UserName        string `form:"UserName"  validate:"required"`
	Name            string `form:"Name" validate:"min=1"`
	Email           string `form:"Email" validate:"required,email"`
	Password        string `form:"Password" validate:"min=5"`
	ConfirmPassword string `form:"ConfirmPassword" validate:"eqfield=Password,required"`
	ProfilePhoto    *multipart.FileHeader
}
