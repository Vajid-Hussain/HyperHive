package interface_usecase_friend_server

import (
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
)

type IFriendUseCase interface {
	FriendRequest(*requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error)
	GetFriends(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	GetReceivedFriendRequest(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	GetSendFriendRequest(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	FriendShipStatusUpdate(string, string) error
	GetBlockFriendRequest(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	MessageConsumer()
	GetFriendChat(string, string, requestmodel_friend_server.Pagination) ([]responsemodel_friend_server.Message, error)
}
