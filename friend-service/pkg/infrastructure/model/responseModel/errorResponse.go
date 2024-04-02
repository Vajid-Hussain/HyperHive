package responsemodel_friend_server

import "errors"

var (
	ErrFriendRequestExist = errors.New("you are alredy initiated a connection")
)

var (
	ErrInternalServer  = errors.New("an unexpected error occurred. Please try again later")
	ErrDBNoRowAffected = errors.New("no row affected")
)
