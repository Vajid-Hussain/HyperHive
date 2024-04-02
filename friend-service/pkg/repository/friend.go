package repository_friend_server

import (
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
	interface_repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interface_repository_friend_server.IFriendRepository {
	return &AdminRepository{DB: db}
}

func (d *AdminRepository) CreateFriend(FriendReq *requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error) {

	var friendRequest responsemodel_friend_server.FriendRequest
	query := "INSERT INTO friends (users, friend, update_at) SELECT $1, $2, $3 WHERE NOT EXISTS ( SELECT 1 FROM friends WHERE (users=$1 AND friend=$2) OR (users=$2 AND friend=$1) ) RETURNING *"
	result := d.DB.Raw(query, FriendReq.User, FriendReq.Friend, FriendReq.UpdateAt).Scan(&friendRequest)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrFriendRequestExist
	}

	return &friendRequest, nil
}

// singnal0

func (d *AdminRepository) GetFriends(userID string) (friends []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT friend, update_at, friends_id FROM friends WHERE users= $1 AND status = 'pending' UNION SELECT users, update_at,friends_id FROM friends WHERE friend = $1 AND status= 'pending' "
	result := d.DB.Raw(query, userID).Scan(&friends)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrDBNoRowAffected
	}

	return friends, nil
}
