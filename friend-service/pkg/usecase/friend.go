package usecase_friend_server

import (
	"time"

	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
	interface_repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository/interface"
	interface_usecase_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/usecase/interface"
)

type FriendUseCase struct {
	friendRepo interface_repository_friend_server.IFriendRepository
}

func NewFriendUseCase(repo interface_repository_friend_server.IFriendRepository) interface_usecase_friend_server.IFriendUseCase {
	return &FriendUseCase{friendRepo: repo}
}

func (r *FriendUseCase) FriendRequest(req *requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error) {
	req.UpdateAt = time.Now()
	response, err := r.friendRepo.CreateFriend(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *FriendUseCase)GetFriendRequest(userID string) ([]*responsemodel_friend_server.FriendList, error){
	return r.friendRepo.GetFriends(userID)
}