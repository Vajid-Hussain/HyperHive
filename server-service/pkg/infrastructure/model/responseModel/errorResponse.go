package responsemodel_server_service

import "errors"

// Repository
var (
	ErrInternalServer               = errors.New("an unexpected error occurred. Please try again later")
	ErrDBNoRowAffected              = errors.New("no row affected")
	ErrPaginationWrongValue         = errors.New("pagination value must be posiive")
	ErrEmptyResponse                = errors.New("no records available")
	ErrStatusNotMatching            = errors.New("status is not not match")
	ErrNotUniqueServerName          = errors.New("server name alrady taken")
	ErrChannelNameAlradyExist       = errors.New("channel name alrady exist in server")
	ErrChannelExistOrNotSuperAdmin  = errors.New("only superAdmin can create non exist channel")
	ErrcategoryExistOrNotSuperAdmin = errors.New("only superAdmin can create non exist category")
	ErrExistMemberJoin              = errors.New("the user is already a member of this server")
)

// usecase
var (
	ErrChannelTypeIsNotMatch = errors.New("channel type is not mathing valid are text, forem, and voice")
)
