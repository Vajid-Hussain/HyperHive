package domain_server_service

// Role of the admin (e.g., 'Owner', 'Moderator', 'Member', etc.)

type Server struct {
	ID          int    `gorm:"primarykey; autoIncrement"`
	Name        string `gorm:"unique; not null"`
	Icon        string
	CoverPhoto  string
	Description string
}

type ChannelCategory struct {
	CategoryID int `gorm:"primaryKey;autoIncrement"`
	ServerID   string
	Fkey       Server `gorm:"foreignkey:ServerID;referances:ID"`
	Name       string
}

type Channels struct {
	ChannelID string `gorm:"primaryKey;autoIncrement"`
	ServerID  string
	FKey      Server `gorm:"foreignkey:ServerID;referances:ID"`
	Name      string
}

type ServerAdmins struct {
	ID       int `gorm:"primarykey;autoIncrement"`
	ServerID string
	Fkey     Server `gorm:"foreignkey:ServerID;referances:ID"`
	UserID   string
	Role     string
}
