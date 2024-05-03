package repository_server_service

import (
	"context"
	"fmt"
	"strconv"

	db_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/db"
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	interface_Repository_Server_Service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type ServerRepository struct {
	DB              *gorm.DB
	mongoCollection *db_server_service.MongoCollection
}

func NewServerRepository(db *gorm.DB, mongoCollection *db_server_service.MongoCollection) interface_Repository_Server_Service.IRepositoryServer {
	return &ServerRepository{
		DB:              db,
		mongoCollection: mongoCollection,
	}
}

func (d *ServerRepository) CreateServer(server *requestmodel_server_service.Server) (*responsemodel_server_service.Server, error) {
	var res responsemodel_server_service.Server

	query := "INSERT INTO servers (name) SELECT $1 WHERE NOT EXISTS (SELECT 1 FROM servers WHERE name= $1 AND status='active') RETURNING *"
	result := d.DB.Raw(query, server.Name).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrNotUniqueServerName
	}
	return &res, nil
}

func (d *ServerRepository) CreateOrUpdateChannelCategory(name string, serverID string) (res *responsemodel_server_service.ChannelCategory, err error) {
	query := "INSERT INTO channel_categories (server_id, name) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM channel_categories WHERE server_id= $1 AND name=$2  AND status='active') RETURNING *"
	result := d.DB.Raw(query, serverID, name).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return res, nil
}

func (d *ServerRepository) CreateSuperAdmin(admin requestmodel_server_service.ServerAdmin) (*responsemodel_server_service.ServerAdmin, error) {
	var res responsemodel_server_service.ServerAdmin
	query := "INSERT INTO server_members (server_id , user_id, role) VALUES ($1, $2, $3)  RETURNING *"
	result := d.DB.Raw(query, admin.ServerID, admin.UserID, admin.Role).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}

	return &res, nil
}

func (d *ServerRepository) CreateCategory(req *requestmodel_server_service.CreateCategory) error {
	query := "INSERT INTO channel_categories (server_id, name) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM channel_categories WHERE server_id=$1 AND name =$2 AND status='active') AND EXISTS(SELECT 1 FROM server_members WHERE server_id= $1 AND user_id= $3 AND role= 'SuperAdmin' AND status='active')"
	result := d.DB.Exec(query, req.ServerID, req.CategoryName, req.UserID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrcategoryExistOrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) CreateChannel(req *requestmodel_server_service.CreateChannel) error {
	query := "INSERT INTO channels (server_id, categoryid, name, type) SELECT $1, $2, $3, $4 WHERE NOT EXISTS (SELECT 1 FROM channels WHERE server_id= $1 AND name= $3  AND status='active') AND EXISTS(SELECT 1 FROM server_members WHERE server_id= $1 AND user_id= $5 AND role= 'SuperAdmin' AND status='active')"
	result := d.DB.Exec(query, req.ServerID, req.CategoryID, req.ChannelName, req.Type, req.UserID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrChannelExistOrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) JoinInServer(req *requestmodel_server_service.JoinToServer) error {
	query := "INSERT INTO server_members (server_id, user_id, role) SELECT $1, $2, $3 WHERE NOT EXISTS (SELECT 1 FROM server_members WHERE server_id= $1 AND user_id= $2  AND status='active')"
	result := d.DB.Exec(query, req.ServerID, req.UserID, req.Role)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrExistMemberJoin
	}
	return nil
}

func (d *ServerRepository) GetServerCategory(serverID string) ([]*responsemodel_server_service.FullServerChannel, error) {
	var res []*responsemodel_server_service.FullServerChannel
	query := "SELECT category_id, name FROM channel_categories WHERE server_id= $1 AND status='active'"
	result := d.DB.Raw(query, serverID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	// if result.RowsAffected == 0 {
	// 	return nil, responsemodel_server_service.ErrEmptyResponse
	// }
	return res, nil
}

func (d *ServerRepository) GetChannelUnderCategory(categoryID string) ([]*responsemodel_server_service.Channel, error) {
	var res []*responsemodel_server_service.Channel
	query := "SELECT * FROM channels WHERE categoryid= $1 AND status='active'"
	result := d.DB.Raw(query, categoryID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	// if result.RowsAffected == 0 {
	// 	return nil, responsemodel_server_service.ErrEmptyResponse
	// }
	return res, nil
}

func (d *ServerRepository) GetServer(serverID string) (*responsemodel_server_service.Server, error) {
	var res *responsemodel_server_service.Server
	query := "SELECT * FROM servers WHERE id=$1 AND status= 'active'"
	result := d.DB.Raw(query, serverID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}
	return res, nil
}

func (d *ServerRepository) GetUserServers(userID string) ([]*responsemodel_server_service.UserServerList, error) {
	fmt.Println(userID)
	var res []*responsemodel_server_service.UserServerList
	query := "SELECT * FROM servers INNER JOIN server_members ON servers.id= server_members.server_id WHERE user_id= $1 AND server_members.status='active' AND servers.status='active'"
	result := d.DB.Raw(query, userID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	// if result.RowsAffected == 0 {
	// 	return nil, responsemodel_server_service.ErrEmptyResponse
	// }
	return res, nil
}

func (d *ServerRepository) UpdateServerCoverPhoto(req *requestmodel_server_service.ServerImages) error {
	query := "UPDATE servers SET cover_photo= $1 WHERE EXISTS (SELECT 1 FROM server_members WHERE user_id =$2 AND role='SuperAdmin' AND server_id =$3)  AND id=$3"
	result := d.DB.Exec(query, req.Url, req.UserID, req.ServerID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) UpdateServerIcon(req *requestmodel_server_service.ServerImages) error {
	query := "UPDATE servers SET icon= $1 WHERE EXISTS (SELECT 1 FROM server_members WHERE user_id= $2 AND role= 'SuperAdmin' AND server_id =$3) AND id=$3"
	result := d.DB.Exec(query, req.Url, req.UserID, req.ServerID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) UpdateServerDiscription(req *requestmodel_server_service.Description) error {
	query := "UPDATE servers SET description= $1 WHERE EXISTS (SELECT 1 FROM server_members WHERE user_id= $2 AND (role= 'SuperAdmin' OR role='Admin') AND server_id =$3) AND id=$3"
	result := d.DB.Exec(query, req.Description, req.UserID, req.ServerID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrNotSuperAdmin
	}
	return nil
}

func (d *ServerRepository) GetServerMembers(serverID string, pagination requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMembers, error) {
	var res []responsemodel_server_service.ServerMembers
	query := "SELECT * FROM server_members WHERE server_id = $1 AND status='active' ORDER BY CASE role WHEN 'SuperAdmin' THEN 1 WHEN 'Admin' THEN 2 WHEN 'member' THEN 3 ELSE 4 END LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, serverID, pagination.Limit, pagination.OffSet).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}
	return res, nil
}

func (d *ServerRepository) ChangeMemberRole(req *requestmodel_server_service.UpdateMemberRole) error {
	query := "UPDATE server_members SET role=$3 WHERE user_id = $1 AND server_id=$4 AND EXISTS (SELECT 1 FROM server_members WHERE (role='SuperAdmin' OR role= 'Admin') AND user_id= $2 AND server_id= $4) AND NOT EXISTS (SELECT 1 FROM server_members WHERE user_id= $1 AND role='SuperAdmin' AND server_id= $4)"
	result := d.DB.Exec(query, req.TargetUserID, req.UserID, req.TargetRole, req.ServerID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrNotAnAdmin
	}
	return nil
}

// /----------
func (d *ServerRepository) RemoveUserFromServer(req *requestmodel_server_service.RemoveUser) error {
	query := "UPDATE server_members SET status='remove' WHERE user_id =$1 AND server_id=$3 AND EXISTS (SELECT 1 FROM server_members WHERE (role='SuperAdmin' OR role= 'Admin') AND user_id= $2 AND server_id =$3) AND NOT EXISTS (SELECT 1 FROM server_members WHERE user_id= $1 AND role='SuperAdmin' AND server_id= $3)"
	result := d.DB.Exec(query, req.RemoverID, req.UserID, req.ServerID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrRemoveMember
	}
	return nil
}

func (d *ServerRepository) LeftFromServer(userID, serveID string) error {
	query := "UPDATE server_members SET status='left' WHERE server_id= $1 AND user_id =$2 AND NOT EXISTS (SELECT 1 FROM server_members WHERE user_id=$2 AND role='SuperAdmin' AND server_id =$1)"
	result := d.DB.Exec(query, serveID, userID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrSuperAdminLeft
	}
	return nil
}

func (d *ServerRepository) DeleteServer(userID, ServerID string) error {
	query := "UPDATE servers SET status='delete' WHERE id = $1 AND EXISTS (SELECT 1 FROM server_members WHERE user_id= $2 AND role='SuperAdmin' AND server_id=$1)"
	result := d.DB.Exec(query, ServerID, userID)
	if result.Error != nil {
		return responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return responsemodel_server_service.ErrNotAnAdmin
	}
	return nil
}

func (d *ServerRepository) GetServers(serverID string, pagination requestmodel_server_service.Pagination) (res []*responsemodel_server_service.Server, err error) {
	query := "SELECT * FROM servers WHERE name ILIKE '%' || $1 || '%' AND status = 'active' ORDER BY name LIMIT $2 OFFSET $3"
	result := d.DB.Raw(query, serverID, pagination.Limit, pagination.OffSet).Scan(&res)
	if result.Error != nil {
		return nil, responsemodel_server_service.ErrInternalServer
	}

	if result.RowsAffected == 0 {
		return nil, responsemodel_server_service.ErrEmptyResponse
	}
	return res, nil
}

// Mongodb Queries

func (d *ServerRepository) KeepMessageInDB(message requestmodel_server_service.ServerMessage) error {
	fmt.Println("message before store ", message)
	_, err := d.mongoCollection.ServerChat.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}
	return nil
}

func (d *ServerRepository) GetChannelMessages(chanelID string, pagination requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMessage, error) {
	fmt.Println("==", chanelID, pagination)
	var messages []responsemodel_server_service.ServerMessage
	limit, _ := strconv.Atoi(pagination.Limit)
	offset, _ := strconv.Atoi(pagination.OffSet)

	option := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	channelIDInt, _ := strconv.Atoi(chanelID)

	cursor, err := d.mongoCollection.ServerChat.Find(context.TODO(), bson.M{"ChannelID": channelIDInt}, option, options.Find().SetSort(bson.D{{"TimeStamp", -1}}))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	cursor.All(context.TODO(), &messages)

	fmt.Println("==", messages)
	return messages, nil
}

func (d *ServerRepository) InsertForumPost(post requestmodel_server_service.ForumPost) error {
	fmt.Println("post ", post)
	_, err := d.mongoCollection.ForunPost.InsertOne(context.TODO(), post)
	return err
}

func (d *ServerRepository) InsertForumCommand(command requestmodel_server_service.FormCommand) error {

	fmt.Println("command ", command)
	_, err := d.mongoCollection.ForumCommand.InsertOne(context.TODO(), command)
	return err
}

func (d *ServerRepository) GetForumPost(channelID string, pagination requestmodel_server_service.Pagination) ([]*responsemodel_server_service.ForumPost, error) {
	fmt.Println(pagination, channelID)
	var post []*responsemodel_server_service.ForumPost
	limit, _ := strconv.Atoi(pagination.Limit)
	offset, _ := strconv.Atoi(pagination.OffSet)

	option := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	channelIDInt, _ := strconv.Atoi(channelID)

	filter := bson.M{"ChannelID": channelIDInt}

	cursor, err := d.mongoCollection.ForunPost.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &post)
	return post, err
}

func (d *ServerRepository) GetForumCommands(parentID string, pagination requestmodel_server_service.Pagination) ([]*responsemodel_server_service.ForumCommand, error) {
	var command []*responsemodel_server_service.ForumCommand

	limit, _ := strconv.Atoi(pagination.Limit)
	offset, _ := strconv.Atoi(pagination.OffSet)

	option := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)).SetSort(bson.D{{"timestamp", -1}})
	filter := bson.M{"parentid": parentID}
	cursor, err := d.mongoCollection.ForumCommand.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &command)
	return command, err
}

func (d *ServerRepository) GetFormSinglePost(PostID string) (*responsemodel_server_service.ForumPost, error) {
	var res *responsemodel_server_service.ForumPost
	objectID, err := primitive.ObjectIDFromHex(PostID)
	if err != nil {
		return nil, err
	}
	fmt.Println("==-9", objectID)

	filter := bson.M{"_id": objectID}
	err = d.mongoCollection.ForunPost.FindOne(context.TODO(), filter).Decode(&res)
	return res, err
}
