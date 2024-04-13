package requestmodel_server_service

type Server struct {
	UserID string
	Name   string
}

type ServerAdmin struct {
	UserID   string
	ServerID string
	Role     string
}

type CreateCategory struct {
	UserID       string
	ServerID     string
	CategoryName string
}

type CreateChannel struct {
	ChannelName string
	UserID      string
	ServerID    string
	CategoryID  string
	Type        string
}

type JoinToServer struct {
	UserID   string
	ServerID string
	Role     string
}

type MemberStatusUpdate struct{
	UserID string
	ServerID string
	TargetUserID string
}
