package usecase_server_service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/IBM/sarama"
	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HyperHive/server-service/pkg/pb"
	interface_Repository_Server_Service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository/interface"
	interface_useCase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase/interface"
	utils_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/utils"
)

type ServerUsecase struct {
	repository interface_Repository_Server_Service.IRepositoryServer
	s3         config_server_service.S3Bucket
	kafka      config_server_service.Kafka
	authClind  pb.AuthServiceClient
}

func NewServerUseCase(repo interface_Repository_Server_Service.IRepositoryServer, kafka config_server_service.Kafka, authClind pb.AuthServiceClient, s3 config_server_service.S3Bucket) interface_useCase_server_service.IServerUseCase {
	return &ServerUsecase{
		repository: repo,
		kafka:      kafka,
		authClind:  authClind,
		s3:         s3,
	}
}

func (r *ServerUsecase) CreateServer(server *requestmodel_server_service.Server) (*responsemodel_server_service.Server, error) {
	serverRes, err := r.repository.CreateServer(server)
	if err != nil {
		return nil, err
	}

	_, err = r.repository.CreateOrUpdateChannelCategory("General", serverRes.ServerID)
	if err != nil {
		return nil, err
	}

	_, err = r.repository.CreateSuperAdmin(requestmodel_server_service.ServerAdmin{UserID: server.UserID, ServerID: serverRes.ServerID, Role: "SuperAdmin"})
	if err != nil {
		return nil, err
	}
	return serverRes, nil
}

func (r *ServerUsecase) CreateCategory(req *requestmodel_server_service.CreateCategory) error {
	return r.repository.CreateCategory(req)
}

func (r *ServerUsecase) CreateChannel(req *requestmodel_server_service.CreateChannel) error {
	if req.Type != "voice" && req.Type != "text" && req.Type != "guild forum" {
		return responsemodel_server_service.ErrChannelTypeIsNotMatch
	}

	return r.repository.CreateChannel(req)
}

func (r *ServerUsecase) JoinToServer(req *requestmodel_server_service.JoinToServer) error {
	req.Role = "member"
	return r.repository.JoinInServer(req)
}

func (r *ServerUsecase) GetServerCategory(serverID string) ([]*responsemodel_server_service.FullServerChannel, error) {
	return r.repository.GetServerCategory(serverID)
}

func (r *ServerUsecase) GetChannels(serverID string) ([]*responsemodel_server_service.FullServerChannel, error) {
	category, err := r.repository.GetServerCategory(serverID)
	if err != nil {
		return nil, err
	}

	for i, val := range category {
		category[i].Channel, _ = r.repository.GetChannelUnderCategory(val.CategoryID)
	}
	return category, nil
}

func (r *ServerUsecase) GetUserServers(userID string) ([]*responsemodel_server_service.UserServerList, error) {
	return r.repository.GetUserServers(userID)
}

func (r *ServerUsecase) GetServer(serverID string) (*responsemodel_server_service.Server, error) {
	return r.repository.GetServer(serverID)
}

func (d *ServerUsecase) UpdateServerPhoto(req *requestmodel_server_service.ServerImages) (err error) {
	s3Session := utils_server_service.CreateSession(d.s3)

	req.Url, err = utils_server_service.UploadImageToS3(req.Image, s3Session)
	if err != nil {
		return err
	}

	if req.Type == "cover photo" {
		err := d.repository.UpdateServerCoverPhoto(req)
		if err != nil {
			return err
		}
	} else if req.Type == "icon" {
		err := d.repository.UpdateServerIcon(req)
		if err != nil {
			return err
		}
	} else {
		return responsemodel_server_service.ErrServerPhotoTypeNotMatch
	}
	return nil
}

func (d *ServerUsecase) UpdateServerDiscription(req *requestmodel_server_service.Description) error {
	if len(req.Description) > 20 {
		return responsemodel_server_service.ErrServerDescriptionLength
	}
	err := d.repository.UpdateServerDiscription(req)
	if err != nil {
		return err
	}
	return nil
}

func (d *ServerUsecase) GetServers(serverID string, pagination requestmodel_server_service.Pagination) ([]*responsemodel_server_service.Server, error) {
	var err error
	pagination.OffSet, err = utils_server_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	return d.repository.GetServers(serverID, pagination)
}

func (r *ServerUsecase) KafkaConsumerServerMessage() error {
	var messageModel requestmodel_server_service.ServerMessage
	var formPost requestmodel_server_service.ForumPost
	var formCommand requestmodel_server_service.FormCommand
	fmt.Println("Kafka started ")

	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{r.kafka.KafkaPort}, config)
	if err != nil {
		fmt.Println("error from kafka ", err)
	}
	defer consumer.Close()

	consumerPartioshion, err := consumer.ConsumePartition(r.kafka.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error from kafka ", err)
	}
	defer consumerPartioshion.Close()

	for {
		message := <-consumerPartioshion.Messages()
		switch string(message.Key) {
		case "server message":
			json.Unmarshal(message.Value, &messageModel)
			fmt.Println("==", messageModel)
			err := r.repository.KeepMessageInDB(messageModel)
			if err != nil {
				fmt.Println("err on adding message in db ", err)
			}
		case "forum post":
			json.Unmarshal(message.Value, &formPost)
			r.repository.InsertForumPost(formPost)
		case "forum command":
			json.Unmarshal(message.Value, &formCommand)
			r.repository.InsertForumCommand(formCommand)
		}
	}
}

func (r *ServerUsecase) GetChannelMessages(channelID string, pagination requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMessage, error) {
	var err error
	pagination.OffSet, err = utils_server_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	message, err := r.repository.GetChannelMessages(channelID, pagination)
	if err != nil {
		return nil, err
	}

	for i, val := range message {
		userProfile, _ := r.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: strconv.Itoa(val.UserID)})
		message[i].UserProfile = userProfile.ProfilePhoto
		message[i].UserName = userProfile.UserName
	}
	return message, nil
}

func (r *ServerUsecase) GetServerMembers(serverID string, pagination requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMembers, error) {
	var err error
	pagination.OffSet, err = utils_server_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	members, err := r.repository.GetServerMembers(serverID, pagination)
	if err != nil {
		return nil, err
	}
	for i, user := range members {
		userDetails, _ := r.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: user.UserID})
		members[i].UserProfile = userDetails.ProfilePhoto
		members[i].UserName = userDetails.UserName
		members[i].Name = userDetails.Name
	}
	return members, nil
}

func (r *ServerUsecase) UpdateMemberRole(req requestmodel_server_service.UpdateMemberRole) error {
	if req.TargetRole == "Admin" || req.TargetRole == "member" {
		return r.repository.ChangeMemberRole(&req)
	}
	return responsemodel_server_service.ErrNotExpectedRole
}

func (r *ServerUsecase) RemoveUserFromServer(req *requestmodel_server_service.RemoveUser) error {
	return r.repository.RemoveUserFromServer(req)
}

func (r *ServerUsecase) DeleteServer(userID, serverID string) error {
	return r.repository.DeleteServer(userID, serverID)
}

func (r *ServerUsecase) LeftFromServer(userID, serverID string) error {
	return r.repository.LeftFromServer(userID, serverID)
}

func (r *ServerUsecase) GetForumPost(channelID string, pagination requestmodel_server_service.Pagination) ([]*responsemodel_server_service.ForumPost, error) {
	var err error
	pagination.OffSet, err = utils_server_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	post, err := r.repository.GetForumPost(channelID, pagination)
	if err != nil {
		return nil, err
	}

	for i, singlePost := range post {
		userProfile, err := r.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: strconv.Itoa(singlePost.UserID)})
		if err != nil {
			return nil, err
		}

		post[i].UserProfile = userProfile.ProfilePhoto
		post[i].UserName = userProfile.UserName

		command, err := r.repository.GetForumCommands(singlePost.ID, requestmodel_server_service.Pagination{Limit: "1", OffSet: "0"})
		if err != nil {
			return nil, err
		}

		if len(command) == 1 {
			post[i].CommandContent = command[0].Content
		}
	}

	return post, nil
}

func (r *ServerUsecase) GetFormSinglePost(postID string) (*responsemodel_server_service.ForumPost, error) {
	post, err := r.repository.GetFormSinglePost(postID)
	if err != nil {
		return nil, err
	}
	userProfile, err := r.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: strconv.Itoa(post.UserID)})
	if err != nil {
		return nil, err
	}

	post.UserName = userProfile.UserName
	post.UserProfile = userProfile.ProfilePhoto

	return post, nil
}

func (r *ServerUsecase) GetPostCommand(postID string, pagination requestmodel_server_service.Pagination) (*responsemodel_server_service.PostCommands, error) {
	var err error
	pagination.OffSet, err = utils_server_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	var commands responsemodel_server_service.PostCommands
	commands.Commands, err = r.repository.GetForumCommands(postID, pagination)
	if err != nil {
		return nil, err
	}

	r.fetchAllCommands(commands.Commands)
	// for _, val := range commands.Commands {
	// 	fmt.Println("==",val)
	// 	fmt.Println("==", val.Thread[0])
	// 	fmt.Println("===",val.Thread[0].Thread[0])
	// }
	return &commands, nil
}

// func (r *ServerUsecase) fetchAllCommands(command []*responsemodel_server_service.ForumCommand, lenght, pos int) {
// 	if lenght-1 == pos {
// 		return
// 	}

// 	command[pos].Thread, _ = r.repository.GetForumCommands(command[pos].ID, requestmodel_server_service.Pagination{Limit: "100", OffSet: "0"})
// 	r.fetchAllCommands(command[pos].Thread, len(command[pos].Thread), pos)
// }

func (r *ServerUsecase) fetchAllCommands(command []*responsemodel_server_service.ForumCommand) {
	for i, value := range command {
		userProfile, _ := r.authClind.UserProfile(context.Background(), &pb.UserProfileRequest{UserID: strconv.Itoa(value.UserID)})
		command[i].UserProfile = userProfile.ProfilePhoto
		command[i].UserName = userProfile.UserName
		command[i].Thread, _ = r.repository.GetForumCommands(value.ID, requestmodel_server_service.Pagination{Limit: "100", OffSet: "0"})
		r.fetchAllCommands(command[i].Thread)
	}
}
