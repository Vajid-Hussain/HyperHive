package domain_server_service

// Role of the admin (e.g., 'Owner', 'Moderator', 'Member', etc.)

type Server struct {
	ID          int    `gorm:"primarykey; autoIncrement"`
	Name        string `gorm:"unique; not null"`
	Icon        string
	CoverPhoto  string
	Description string
	Status      string `gorm:"default:active"`
}

type ChannelCategory struct {
	CategoryID int `gorm:"primaryKey;autoIncrement"`
	ServerID   string
	Fkey       Server `gorm:"foreignkey:ServerID;referances:ID"`
	Name       string
	Status     string `gorm:"default:active"`
}

type Channels struct {
	ChannelID    int `gorm:"primaryKey;autoIncrement"`
	ServerID     string
	FKey         Server `gorm:"foreignkey:ServerID;referances:ID"`
	Categoryid   string
	FkeyCategory ChannelCategory `gorm:"foreignkey:Categoryid;references:CategoryID"`
	Name         string
	Type         string
	Status       string `gorm:"default:active"`
}

type ServerMembers struct {
	ID       int `gorm:"primarykey;autoIncrement"`
	ServerID string
	Fkey     Server `gorm:"foreignkey:ServerID;referances:ID"`
	UserID   string
	Role     string
	Status   string `gorm:"default:active"`
}
