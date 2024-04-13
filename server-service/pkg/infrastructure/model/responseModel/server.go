package responsemodel_server_service

type Server struct {
	ServerID    string `gorm:"column:id"`
	Name        string
	Description string
	Icon        string
	CoverPhoto  string
}

type ChannelCategory struct {
	CategoryID string
	ServerID   string
	Name       string
}

type ServerAdmin struct {
	ID       string
	UserID   string
	ServerID string
	Role     string
}

type Channel struct {
	ChannelID  string
	CategoryID string
	Name       string
	Type       string
}

type FullServerChannel struct {
	CategoryID string
	Name       string
	Channel    []*Channel `gorm:"-"`
}

type UserServerList struct {
	ServerID string	
}
