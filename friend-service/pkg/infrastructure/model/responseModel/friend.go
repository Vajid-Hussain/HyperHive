package responsemodel_friend_server

import "time"

type FriendRequest struct {
	FriendsID string
	User      string
	Friend    string
	UpdateAt  time.Time
	Status string
}
