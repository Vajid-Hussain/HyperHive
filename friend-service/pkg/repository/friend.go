package repository_friend_server

import (
	"fmt"

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

func (d *AdminRepository) GetFriends(req *requestmodel_friend_server.GetFriendRequest) (friends []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT friend, update_at, friendship_id FROM friends WHERE users= $1 AND status = 'active' UNION SELECT users, update_at,friendship_id FROM friends WHERE friend = $1 AND status= 'active' "
	result := d.DB.Raw(query, req.UserID).Scan(&friends)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return friends, nil
}

func (d *AdminRepository) GetReceivedFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (request []*responsemodel_friend_server.FriendList, err error) {
	fmt.Println("--", req)
	query := "SELECT * FROM friends WHERE friend= $1 AND status= 'pending' ORDER BY update_at DESC LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, req.UserID, req.Limit, req.OffSet).Scan(&request)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return request, nil
}

func (d *AdminRepository) GetSendFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (request []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT * FROM friends WHERE users= $1 AND status= 'pending' ORDER BY update_at DESC LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, req.UserID, req.Limit, req.OffSet).Scan(&request)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return request, nil
}

func (d *AdminRepository) GetBlockFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (request []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT * FROM friends WHERE users= $1 AND status= 'block' ORDER BY update_at DESC LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, req.UserID, req.Limit, req.OffSet).Scan(&request)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return request, nil
}

func (d *AdminRepository) FriendShipStatusUpdate(friendShipID, status string) error {
	query := "UPDATE friends SET status= $1 WHERE friends_id =$2"
	result := d.DB.Exec(query, status, friendShipID)
	if result.Error != nil {
		return responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_friend_server.ErrEmptyResponse
	}

	return nil
}
