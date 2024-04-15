package responsemodel_websocket_svc

import "errors"

var (
	ErrUndefinedMessageCategory = errors.New("undefinded message category")
	ErrUnmarshelMessageCategory = errors.New("face errro while finding message category")
	ErrWhileUnmarshelChatMessage     = errors.New("some problem for unmarshel chat message")
)
