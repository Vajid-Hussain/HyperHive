package requestmodel_auth_svc

type UserSignup struct {
	UserName        string `josn:"UserName"  validate:"required"`
	Name            string `josn:"Name" validate:"min=1"`
	Email           string `josn:"Email" validate:"required,email"`
	Password        string `josn:"Password" validate:"min=5"`
	ConfirmPassword string `josn:"ConfirmPassword" validate:"eqfield=Password,required"`
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
