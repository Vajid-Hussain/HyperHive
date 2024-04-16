package usecase_server_service

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	interface_Repository_Server_Service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/repository/interface"
	interface_useCase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase/interface"
	utils_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/utils"
	"github.com/Vajid-Hussain/HyperHive/server-service/pkg/pb"
)

type ServerUsecase struct {
	repository interface_Repository_Server_Service.IRepositoryServer
	s3         config_server_service.S3Bucket
	kafka      config_server_service.Kafka
	authClind pb.AuthServiceClient
}

func NewServerUseCase(repo interface_Repository_Server_Service.IRepositoryServer, kafka config_server_service.Kafka,authClind pb.AuthServiceClient) interface_useCase_server_service.IServerUseCase {
	return &ServerUsecase{
		repository: repo,
		kafka:      kafka,
		authClind: authClind,
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
	if req.Type != "voice" && req.Type != "text" && req.Type != "forum" {
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

// func (d *ServerUsecase) UpdateCoverPhoto(req requestmodel_server_service.ServerImages) (err error) {
// 	s3Session := utils_server_service.CreateSession(d.s3)

// 	url, err = utils_server_service.UploadImageToS3(image, s3Session)
// 	if err != nil {
// 		return err
// 	}

// 	err = d.repository.(userID, url)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *ServerUsecase) KafkaConsumerServerMessage() error {
	var messageModel requestmodel_server_service.ServerMessage
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
		json.Unmarshal(message.Value, &messageModel)
		err := r.repository.KeepMessageInDB(messageModel)
		if err != nil {
			fmt.Println("err on adding message in db ", err)
		}
	}
}

func (r *ServerUsecase) GetChannelMessages(channelID string, pagination requestmodel_server_service.Pagination) ([]responsemodel_server_service.ServerMessage, error) {
	var err error
	pagination.OffSet, err = utils_server_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	return r.repository.GetChannelMessages(channelID, pagination)
}
