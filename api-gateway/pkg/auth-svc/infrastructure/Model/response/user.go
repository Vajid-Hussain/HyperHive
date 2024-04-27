package response_auth_svc

type UserProfile struct {
	UserID       string `json:"UserID,omitempty"`
	UserName     string `json:"UserName,omitempty"`
	Name         string `json:"Name,omitempty"`
	Email        string `json:"Email,omitempty"`
	ProfilePhoto string `json:"ProfilePhoto,omitempty"`
	CoverPhoto   string `json:"CoverPhoto,omitempty"`
	Description  string `json:"Description,omitempty"`
	Status       string `json:"Status,omitempty"`
	UserSince    string `json:"UserSince"`
}
