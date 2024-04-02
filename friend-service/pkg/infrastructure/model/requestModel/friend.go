package requestmodel_friend_server

import "time"

type FriendRequest struct {
	User     string
	Friend   string
	UpdateAt time.Time
}
