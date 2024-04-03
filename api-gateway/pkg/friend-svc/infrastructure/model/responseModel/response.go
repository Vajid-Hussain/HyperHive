package responsemodel_friend_svc

import (
	"strings"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"result,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func Responses(statusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      trimPrefixOfRpcError(err),
	}
}

func trimPrefixOfRpcError(err interface{}) interface{} {
	errMessage, ok := err.(string)
	if ok {
		return strings.TrimPrefix(errMessage, "rpc error: code = Unknown desc = ")
	}
	return err
}
