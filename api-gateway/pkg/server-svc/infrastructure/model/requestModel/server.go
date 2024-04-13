package requestmodel_server_svc

type Server struct {
	Name string `json:"Name"`
}

type CreateCategory struct {
	UserID       string `json:"-"`
	ServerID     string `json:"serverID"`
	CategoryName string `json:"categoryName"`
}

type CreateChannel struct {
	ChannelName string `json:"channelName"`
	UserID      string `json:"-"`
	ServerID    string `json:"serverID"`
	CategoryID  string `json:"categoryID"`
	Type        string `json:"type"`
}

type JoinToServer struct {
	UserID   string `json:"-"`
	ServerID string `json:"ServerID"`
}

type ServerReq struct {
	ServerID string `json:"ServerID"`
}
