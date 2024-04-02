package interface_repository_friend_server

import (
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
)

type IFriendRepository interface {
	CreateFriend(*requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error)
	GetFriends(string) ([]*responsemodel_friend_server.FriendList, error)
}
