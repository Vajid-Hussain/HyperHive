package responsemodel_friend_server

import "errors"

var (
	ErrFriendRequestExist = errors.New("you are alredy initiated a connection")
)

var (
	ErrInternalServer       = errors.New("an unexpected error occurred. Please try again later")
	ErrDBNoRowAffected      = errors.New("no row affected")
	ErrPaginationWrongValue = errors.New("pagination value must be posiive")
	ErrEmptyResponse        = errors.New("No records available.")
	ErrStatusNotMatching    = errors.New("status is not not match")
)
