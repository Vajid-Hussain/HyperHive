package responsemodel_auth_server

import "errors"

// Repository
var (
	ErrUnauthorizedAccess  = errors.New("unauthorized access. Please log in to continue")
	ErrNotFound            = errors.New("the requested resource was not found")
	ErrInternalServer      = errors.New("an unexpected error occurred. Please try again later")
	ErrInvalidInput        = errors.New("invalid input. Please check your data and try again")
	ErrUserBlockedOrNoUser = errors.New("no resourse or user is blocked")

	ErrDatabaseQuery = errors.New("an error occurred while processing your request")
	ErrTimeout       = errors.New("ahe request timed out. Please try again later")
)

// UseCase
var (
	ErrLoginNoActiveUser            = errors.New("no active user associated with email id")
	ErrUsernameTaken                = errors.New("username already taken, please choose another")
	ErrEmailExists                  = errors.New("email already exists, please use a different one or log in")
	ErrEnterWithoutEmailConfimation = errors.New("please confirm your email before proceeding")

	// User Profile
	ErrStatuTimeLongExpireTime = errors.New("duration must under 6 hours")
)
