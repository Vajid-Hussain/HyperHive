package requestmodel_auth_server

type UserSignup struct {
	UserName        string
	Name            string
	Email           string
	Password        string
	ConfirmPassword string
	ProfilePhoto    []byte
	CoverPhoto []byte
	ProfilePhotoUrl string
	CoverPhotoUrl string
}
