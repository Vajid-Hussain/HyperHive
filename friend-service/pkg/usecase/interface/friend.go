package interface_usecase_friend_server

import (
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
)

type IFriendUseCase interface{
	FriendRequest( *requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error)
	GetFriendRequest( string) ([]*responsemodel_friend_server.FriendList, error)
}
