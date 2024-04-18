package responsemodel_auth_server

import "errors"

// Repository
var (
	ErrUnauthorizedAccess = errors.New("unauthorized access. Please log in to continue")
	ErrNotFound           = errors.New("the requested resource was not found")
	ErrOtpIsExpire        = errors.New("no otp available or expired")
	ErrOtpNotMatch        = errors.New("otp not match")
	ErrInternalServer     = errors.New("an unexpected error occurred. Please try again later")
	ErrDBNoRowAffected    = errors.New("no row affected")
	ErrSerchUsers         = errors.New("no user exist")

	ErrInvalidInput        = errors.New("invalid input. Please check your data and try again")
	ErrUserBlockedOrNoUser = errors.New("no resourse or user is blocked")

	ErrDatabaseQuery = errors.New("an error occurred while processing your request")
	ErrTimeout       = errors.New("ahe request timed out. Please try again later")
)

// UseCase
var (
	ErrLoginNoActiveUser            = errors.New("no active user associated with email id")
	ErrUsernameTaken                = errors.New("username already taken, please choose another")
	ErrEmailExists                  = errors.New("email associated with an accound")
	ErrEnterWithoutEmailConfimation = errors.New("please confirm your email before proceeding")
	ErrNoUserExist                  = errors.New("no user exist")
	ErrEmailNotVerified             = errors.New("confirm mail to proceesed")
	ErrDeletePhotoProfile           = errors.New("mention which photo should delete")
	// User Profile
	ErrStatuTimeLongExpireTime = errors.New("duration must under 24 hours")
	ErrRegexNotMatch           = errors.New("password is not satify criteria")
	ErrDBQueryExecution        = errors.New("request have some missaderstanding polisht your request")
	ErrPaginationWrongValue    = errors.New("pagination value must be posiive")
)
