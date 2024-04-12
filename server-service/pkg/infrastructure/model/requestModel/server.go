package requestmodel_server_service

type Server struct {
	Name string
}

type ServerAdmin struct {
	UserID   string
	ServerID string
	Role     string
}
