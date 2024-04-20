package responsemodel_friend_server

import "errors"

// usecase
var (
	ErrFriendRequestExist               = errors.New("you are alredy initiated a connection")
	ErrFriendRequestUserAndFriendIsSame = errors.New("both user and friend are same")
)

// repository
var (
	ErrInternalServer       = errors.New("an unexpected error occurred. Please try again later")
	ErrDBNoRowAffected      = errors.New("no row affected")
	ErrPaginationWrongValue = errors.New("pagination value must be posiive")
	ErrEmptyResponse        = errors.New("no records available")
	ErrStatusNotMatching    = errors.New("status is not not match")
)
