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

type UserLogin struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=5"`
}
