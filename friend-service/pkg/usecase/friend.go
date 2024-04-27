package usecase_friend_server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	_ "time/tzdata"

	"github.com/IBM/sarama"
	config_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/config"
	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	responsemodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HyperHive/friend-service/pkg/pb"
	interface_repository_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/repository/interface"
	interface_usecase_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/usecase/interface"
	utils_friend_service "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/utils"
)

type FriendUseCase struct {
	friendRepo interface_repository_friend_server.IFriendRepository
	authClind  pb.AuthServiceClient
	Location   *time.Location
	Kafka      config_friend_server.Kafka
}

func NewFriendUseCase(repo interface_repository_friend_server.IFriendRepository, authClind pb.AuthServiceClient, Kafka config_friend_server.Kafka) interface_usecase_friend_server.IFriendUseCase {
	locationInd, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("error at exct place", err)
	}
	return &FriendUseCase{friendRepo: repo,
		authClind: authClind,
		Location:  locationInd,
		Kafka:     Kafka}
}

func (r *FriendUseCase) FriendRequest(req *requestmodel_friend_server.FriendRequest) (*responsemodel_friend_server.FriendRequest, error) {
	if req.User == req.Friend {
		return nil, responsemodel_friend_server.ErrFriendRequestUserAndFriendIsSame
	}

	req.UpdateAt = time.Now()
	response, err := r.friendRepo.CreateFriend(req)
	if err != nil {
		return nil, err
	}
	response.UpdateAt = response.UpdateAt.In(r.Location)
	return response, nil
}

func (r *FriendUseCase) GetFriends(req *requestmodel_friend_server.GetFriendRequest) (res []*responsemodel_friend_server.FriendList, err error) {
	req.OffSet, err = utils_friend_service.Pagination(req.Limit, req.OffSet)
	if err != nil {
		return nil, err
	}

	friendRequest, err := r.friendRepo.GetFriends(req)
	if err != nil {
		return nil, err
	}

	return r.FriendListReponse(friendRequest, req.UserID), nil
}

func (r *FriendUseCase) GetReceivedFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (res []*responsemodel_friend_server.FriendList, err error) {

	req.OffSet, err = utils_friend_service.Pagination(req.Limit, req.OffSet)
	if err != nil {
		return nil, err
	}

	receivedRequest, err := r.friendRepo.GetReceivedFriendRequest(req)
	if err != nil {
		return nil, err
	}

	return r.ReceivedFriendRequestResponse(receivedRequest), nil
}

func (r *FriendUseCase) GetSendFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (res []*responsemodel_friend_server.FriendList, err error) {

	req.OffSet, err = utils_friend_service.Pagination(req.Limit, req.OffSet)
	if err != nil {
		return nil, err
	}

	sendRequest, err := r.friendRepo.GetSendFriendRequest(req)
	if err != nil {
		return nil, err
	}

	return r.CreateFriendListResponse(sendRequest), nil
}

func (r *FriendUseCase) GetBlockFriendRequest(req *requestmodel_friend_server.GetFriendRequest) (res []*responsemodel_friend_server.FriendList, err error) {

	req.OffSet, err = utils_friend_service.Pagination(req.Limit, req.OffSet)
	if err != nil {
		return nil, err
	}

	sendRequest, err := r.friendRepo.GetBlockFriendRequest(req)
	if err != nil {
		return nil, err
	}

	return r.CreateFriendListResponse(sendRequest), nil
}

//-------------- Fetch Friend details from Auth server

func (r *FriendUseCase) CreateFriendListResponse(friendList []*responsemodel_friend_server.FriendList) []*responsemodel_friend_server.FriendList {
	var ch = make(chan *responsemodel_friend_server.AbstractUserProfile)
	var mp = make(map[string]*responsemodel_friend_server.AbstractUserProfile)

	for _, val := range friendList {
		go r.UserProfile(val.FriendID, ch)
	}

	for i := 1; i <= len(friendList); i++ {
		userProfile := <-ch
		if userProfile != nil {
			mp[userProfile.UserID] = userProfile
		}
	}

	for i, val := range friendList {
		profile := mp[val.FriendID]
		if profile == nil {
			friendList[i] = nil
		} else {
			friendList[i].UserProfile = *profile
			friendList[i].UpdateAt = friendList[i].UpdateAt.In(r.Location)
		}
	}

	return friendList
}

func (r *FriendUseCase) ReceivedFriendRequestResponse(friendList []*responsemodel_friend_server.FriendList) []*responsemodel_friend_server.FriendList {
	var ch = make(chan *responsemodel_friend_server.AbstractUserProfile)
	var mp = make(map[string]*responsemodel_friend_server.AbstractUserProfile)

	for _, val := range friendList {
		go r.UserProfile(val.UserID, ch)
	}

	for i := 1; i <= len(friendList); i++ {
		userProfile := <-ch
		if userProfile != nil {
			mp[userProfile.UserID] = userProfile
		}
	}

	for i, val := range friendList {
		profile := mp[val.UserID]
		if profile == nil {
			friendList[i] = nil
		} else {
			friendList[i].UserProfile = *profile
			friendList[i].UpdateAt = friendList[i].UpdateAt.In(r.Location)
		}
	}

	return friendList
}

func (r *FriendUseCase) FriendListReponse(friendList []*responsemodel_friend_server.FriendList, userID string) []*responsemodel_friend_server.FriendList {
	var ch = make(chan *responsemodel_friend_server.AbstractUserProfile)
	var mp = make(map[string]*responsemodel_friend_server.AbstractUserProfile)

	for _, val := range friendList {
		go r.UserProfile(val.FriendID, ch)
	}

	for i := 1; i <= len(friendList); i++ {
		userProfile := <-ch
		if userProfile != nil {
			mp[userProfile.UserID] = userProfile
		}
	}

	for i, val := range friendList {
		profile := mp[val.FriendID]
		if profile == nil {
			friendList[i] = nil
		} else {
			friendList[i].UserProfile = *profile
			res, err := r.friendRepo.GetLastMessage(userID, val.FriendID)

			if err != nil {
				// fmt.Println("==, ", err)
			}

			if res == nil {
				friendList[i].UpdateAt = friendList[i].UpdateAt.In(r.Location)
			} else {
				friendList[i].UpdateAt = res.Timestamp.In(r.Location)
				friendList[i].LastMessage = res.Content
				friendList[i].LastMessageSenderID = res.SenderID
				if msgCount, err := r.friendRepo.GetMessageCount(userID, val.FriendID); err == nil {
					friendList[i].UnreadMessage = msgCount
				}
			}
		}
	}

	return friendList
}

func (r *FriendUseCase) UserProfile(userID string, ch chan *responsemodel_friend_server.AbstractUserProfile) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	result, err := r.authClind.UserProfile(context, &pb.UserProfileRequest{UserID: userID})
	if err != nil {
		ch <- nil
	} else {
		ch <- &responsemodel_friend_server.AbstractUserProfile{UserID: result.UserID, UserName: result.UserName, Name: result.Name, ProfilePhoto: result.ProfilePhoto}
	}
}

func (r *FriendUseCase) FriendShipStatusUpdate(req requestmodel_friend_server.FriendShipStatus) error {
	if req.Status == "block" || req.Status == "unblock" || req.Status == "accept" || req.Status == "reject" || req.Status == "revoke" {
		if req.Status == "accept" || req.Status == "unblock" {
			req.Status = "active"
		}

		err := r.friendRepo.FriendShipStatusUpdate(req)
		if err != nil {
			return err
		}
	} else {
		return responsemodel_friend_server.ErrStatusNotMatching
	}
	return nil
}

func (u *FriendUseCase) MessageConsumer() {
	fmt.Println("Kafka started ")
	configs := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{u.Kafka.KafkaPort}, configs)
	if err != nil {
		fmt.Println("err: ", err)
	}
	defer consumer.Close()

	consumerPartishion, err := consumer.ConsumePartition(u.Kafka.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("err :", err)
	}
	defer consumerPartishion.Close()

	for {
		message := <-consumerPartishion.Messages()
		// fmt.Println("--message: ", string(message.Value))
		msg, _ := u.UnmarshelChatMessage(message.Value)
		fmt.Println("===", msg)
		u.friendRepo.StoreFriendsChat(*msg)
	}
}

func (u *FriendUseCase) UnmarshelChatMessage(data []byte) (*requestmodel_friend_server.Message, error) {
	var message requestmodel_friend_server.Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	// fmt.Println("unmarshel ", message)
	message.Timestamp = time.Now()
	return &message, nil
}

func (u *FriendUseCase) GetFriendChat(userID, friendID string, pagination requestmodel_friend_server.Pagination) ([]responsemodel_friend_server.Message, error) {
	var err error
	pagination.OffSet, err = utils_friend_service.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	_ = u.friendRepo.UpdateReadAsMessage(userID, friendID)
	return u.friendRepo.GetFriendChat(userID, friendID, pagination)
}
