package requestmodel_auth_svc

type UserSignup struct {
	UserName        string `json:"UserName"  validate:"required"`
	Name            string `json:"Name" validate:"min=1"`
	Email           string `json:"Email" validate:"required,email"`
	Password        string `json:"Password" validate:"min=5"`
	ConfirmPassword string `json:"ConfirmPassword" validate:"eqfield=Password,required"`
	// ProfilePhoto    *multipart.FileHeader
}

type UserLogin struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=5"`
}

type UserProfileStatus struct {
	UserID   string  `json:"-"`
	Status   string  `json:"Status" validate:"required"`
	Duration float32 `json:"Duration" validate:"required"`
}

type UserProfileDescription struct {
	UserID      string `json:"-"`
	Description string `json:"Description" validate:"required"`
}

type UserIDReq struct {
	UserID string `json:"UserID"`
}

type EmailReq struct {
	Email string `json:"Email" validate:"email,required"`
}

type ForgotPassword struct {
	Password        string `json:"Password" validate:"min=5"`
	ConfirmPassword string `json:"ConfirmPassword" validate:"eqfield=Password,required"`
	Otp             string `json:"Otp" validate:"len=4"`
}
