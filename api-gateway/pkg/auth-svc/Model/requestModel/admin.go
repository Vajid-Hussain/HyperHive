package requestmodel_auth_svc

type AdminLogin struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=5"`
}
