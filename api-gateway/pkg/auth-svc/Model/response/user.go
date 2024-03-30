package response_auth_svc


type UserProfile struct {
	UserID       string `json:"UserID"`
	UserName     string `json:"UserName"  `
	Name         string `json:"Name" `
	Email        string `json:"Email"`
	ProfilePhoto string `json:"ProfilePhoto"`
	CoverPhoto   string `json:"CoverPhoto"`
	Description  string `json:"Description"`
	Status       string `json:"Status"`
}
