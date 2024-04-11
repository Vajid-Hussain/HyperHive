package responsemodel_server_service

type Server struct{
	ServerID string `gorm:"column=id"`
	Name string
	Description string
	Icon string
	CoverPhoto string
}
