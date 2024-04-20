package interface_repository_friend_server

import (
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
)

type IFriendRepository interface {
	CreateFriend(*requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error)
	GetFriends(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	GetReceivedFriendRequest(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	GetSendFriendRequest(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)
	FriendShipStatusUpdate(requestmodel_friend_server.FriendShipStatus) error
	GetBlockFriendRequest(*requestmodel_friend_server.GetFriendRequest) ([]*responsemodel_friend_server.FriendList, error)

	//------- mongo
	StoreFriendsChat(requestmodel_friend_server.Message) error
	GetLastMessage(string, string) (*responsemodel_friend_server.Message, error)
	GetMessageCount(string,  string) (int, error)
	GetFriendChat(string, string, requestmodel_friend_server.Pagination) ([]responsemodel_friend_server.Message, error)
	UpdateReadAsMessage(string,  string) error
}
