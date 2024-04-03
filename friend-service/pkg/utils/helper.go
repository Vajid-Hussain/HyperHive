package utils_friend_service

import (
	"strconv"

	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
)

func Pagination(limit, offset string) (string, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return "", err
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return "", err
	}

	if limitInt < 1 || offsetInt < 1 {
		return "", responsemodel_friend_server.ErrPaginationWrongValue
	}

	return strconv.Itoa((offsetInt * limitInt) - limitInt), nil
}
