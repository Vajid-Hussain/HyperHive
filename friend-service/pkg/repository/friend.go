package repository_friend_server

import (
	"context"
	"fmt"
	"strconv"

	db_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/db"
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
	interface_repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type FriendRepository struct {
	DB              *gorm.DB
	mongoCollection *db_friend_server.MongoDBCollections
}

func NewFriendRepository(db *gorm.DB, mongoCollection *db_friend_server.MongoDBCollections) interface_repository_friend_server.IFriendRepository {
	return &FriendRepository{DB: db,
		mongoCollection: mongoCollection}
}

func (d *FriendRepository) CreateFriend(FriendReq *requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error) {

	var friendRequest responsemodel_friend_server.FriendRequest
	query := "INSERT INTO friends (users, friend, update_at) SELECT $1, $2, $3 WHERE NOT EXISTS ( SELECT 1 FROM friends WHERE (users=$1 AND friend=$2 AND status!='revoke') OR (users=$2 AND friend=$1 AND status!='revoke') ) RETURNING *"
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

func (d *FriendRepository) GetFriends(req *requestmodel_friend_server.GetFriendRequest) (friends []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT friend, update_at, friend_ship_id FROM friends WHERE users= $1 AND status = 'active' UNION SELECT users, update_at,friend_ship_id FROM friends WHERE friend = $1 AND status= 'active' "
	result := d.DB.Raw(query, req.UserID).Scan(&friends)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return friends, nil
}

func (d *FriendRepository) GetReceivedFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (response []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT * FROM friends WHERE friend= $1 AND status= 'pending' ORDER BY update_at DESC LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, req.UserID, req.Limit, req.OffSet).Scan(&response)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return response, nil
}

func (d *FriendRepository) GetSendFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (response []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT * FROM friends WHERE users= $1 AND status= 'pending' ORDER BY update_at DESC LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, req.UserID, req.Limit, req.OffSet).Scan(&response)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return response, nil
}

func (d *FriendRepository) GetBlockFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (response []*responsemodel_friend_server.FriendList, err error) {
	query := "SELECT * FROM friends WHERE users= $1 AND status= 'block' ORDER BY update_at DESC LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, req.UserID, req.Limit, req.OffSet).Scan(&response)
	if result.Error != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_friend_server.ErrEmptyResponse
	}

	return response, nil
}

func (d *FriendRepository) FriendShipStatusUpdate(friendShipID, status string) error {
	query := "UPDATE friends SET status= $1 WHERE friend_ship_id =$2"
	result := d.DB.Exec(query, status, friendShipID)
	if result.Error != nil {
		return responsemodel_friend_server.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_friend_server.ErrEmptyResponse
	}

	return nil
}

// -------------- MongoQuery

func (d *FriendRepository) StoreFriendsChat(message requestmodel_friend_server.Message) error {
	// fmt.Println("==repo", message)
	_, err := d.mongoCollection.FriendChatCollection.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	return nil
}

func (d *FriendRepository) GetLastMessage(userID, friendID string) (*responsemodel_friend_server.Message, error) {
	var res = responsemodel_friend_server.Message{}
	filter := bson.M{"senderid": bson.M{"$in": bson.A{userID, friendID}}, "recipientid": bson.M{"$in": bson.A{friendID, userID}}}
	option := options.FindOne().SetSort(bson.D{{"timestamp", -1}})

	err := d.mongoCollection.FriendChatCollection.FindOne(context.TODO(), filter, option).Decode(&res)
	if err != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}
	return &res, nil
}

func (d *FriendRepository) GetMessageCount(userID, friendID string) (int, error) {

	filter := bson.M{"senderid": friendID, "recipientid": userID, "status": "pending"}
	count, err := d.mongoCollection.FriendChatCollection.CountDocuments(context.TODO(), filter)
	if err != nil {
		fmt.Println("-", err)
		return 0, responsemodel_friend_server.ErrInternalServer
	}
	fmt.Println("count", count)
	return int(count), nil
}

func (d *FriendRepository) GetFriendChat(userID, friendID string, pagination requestmodel_friend_server.Pagination) ([]responsemodel_friend_server.Message, error) {

	var messages []responsemodel_friend_server.Message
	filter := bson.M{"senderid": bson.M{"$in": bson.A{userID, friendID}}, "recipientid": bson.M{"$in": bson.A{friendID, userID}}}
	limit, _ := strconv.Atoi(pagination.Limit)
	offset, _ := strconv.Atoi(pagination.OffSet)

	option := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := d.mongoCollection.FriendChatCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{"timestamp", -1}}), option)
	if err != nil {
		return nil, responsemodel_friend_server.ErrInternalServer
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var message responsemodel_friend_server.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, responsemodel_friend_server.ErrInternalServer
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (d *FriendRepository) UpdateReadAsMessage(userID, friendID string) error{

	_,err:= d.mongoCollection.FriendChatCollection.UpdateMany(context.TODO(), bson.M{"senderid": bson.M{"$in": bson.A{ friendID}} , "recipientid": bson.M{"$in":bson.A{userID}}}, bson.D{{ "$set", bson.D{{ "status", "send"}}  } } )
	if err!=nil{
		return responsemodel_friend_server.ErrInternalServer
	}
	return nil
}