package domainl_auth_server

type Admin struct{
	ID int `gorm:"primarykey"`
	Emain string 
	Password string
}