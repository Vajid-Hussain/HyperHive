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
	ErrExistMemberJoin              = errors.New("you are alrady a member")
	ErrNotAnAdmin                   = errors.New("you are not a admin")
	ErrNotSuperAdmin                = errors.New("you are not a super admin")
	ErrRemoveMember                 = errors.New("you can't remove the member from server because of the role")
	ErrSuperAdminLeft               = errors.New("super admin can't left")
)

// usecase
var (
	ErrChannelTypeIsNotMatch   = errors.New("channel type is not mathing valid are text, forem, and voice")
	ErrServerPhotoTypeNotMatch = errors.New("photo type is not matching")
	ErrServerDescriptionLength = errors.New("description under under 20 letters")
	ErrNotExpectedRole         = errors.New("undefine role, confirm role")
)
