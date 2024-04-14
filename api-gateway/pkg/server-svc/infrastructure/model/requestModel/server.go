package requestmodel_server_svc

type Server struct {
	Name string `json:"Name" validate:"required"`
}

type CreateCategory struct {
	UserID       string `json:"-"`
	ServerID     string `json:"serverID" validate:"required"`
	CategoryName string `json:"categoryName" validate:"required"`
}

type CreateChannel struct {
	ChannelName string `json:"channelName" validate:"required"`
	UserID      string `json:"-"`
	ServerID    string `json:"serverID" validate:"required"`
	CategoryID  string `json:"categoryID" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

type JoinToServer struct {
	UserID   string `json:"-"`
	ServerID string `json:"ServerID" validate:"required"`
}

type ServerReq struct {
	ServerID string `json:"ServerID" param:"id" validate:"required"`
}
