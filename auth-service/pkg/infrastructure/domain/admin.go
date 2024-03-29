package domainl_auth_server

type Admins struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Email    string
	Password string
}
