package server_friend_server

import (
	"context"
	"fmt"

	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	"github.com/Vajid-Hussain/HyperHive/friend-service/pkg/pb"
	interface_usecase_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/usecase/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FriendServer struct {
	useCase interface_usecase_friend_server.IFriendUseCase
	pb.UnimplementedFriendServiceServer
}

func NewFriendServer(usecase interface_usecase_friend_server.IFriendUseCase) *FriendServer {
	return &FriendServer{useCase: usecase}
}

func (u *FriendServer) FriendRequest(ctx context.Context, req *pb.FriendRequestRequest) (*pb.FriendRequestResponse, error) {
	fmt.Println("friend request called")
	var friendRequest requestmodel_friend_server.FriendRequest
	friendRequest.Friend = req.FriendID
	friendRequest.User = req.UserID

	result, err := u.useCase.FriendRequest(&friendRequest)
	if err != nil {
		return nil, err
	}

	return &pb.FriendRequestResponse{
		FriendsID: result.FriendsID,
		UserID:    result.User,
		Status:    result.Status,
		UpdateAt:  result.UpdateAt.String(),
	}, nil
}

func (u *FriendServer) FriendList(ctx context.Context, req *pb.FriendListRequest) (*pb.FriendListResponse, error) {
	var constrain requestmodel_friend_server.GetFriendRequest
	constrain.Limit = req.Friend.Limit
	constrain.OffSet = req.Friend.OffSet
	constrain.UserID = req.Friend.UserID

	result, err := u.useCase.GetFriends(&constrain)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.GetPendingListResponseModel
	for _, val := range result {
		if val != nil {
			finalResult = append(finalResult, &pb.GetPendingListResponseModel{
				FriendID:        val.FriendID,
				UpdateAt:        val.UpdateAt.String(),
				UniqueFriendsID: val.UniqueFriendID,
				UserID:          val.UserProfile.UserID,
				UserName:        val.UserProfile.UserName,
				Name:            val.UserProfile.Name,
				ProfilePhoto:    val.UserProfile.ProfilePhoto,
			})
		}
	}
	return &pb.FriendListResponse{Friends: finalResult}, nil
}

func (u *FriendServer) GetReceivedFriendRequest(ctx context.Context, req *pb.GetReceivedFriendRequestRequest) (*pb.GetReceivedFriendRequestResponse, error) {

	var listRequest requestmodel_friend_server.GetFriendRequest
	listRequest.Limit = req.Received.Limit
	listRequest.OffSet = req.Received.OffSet
	listRequest.UserID = req.Received.UserID

	result, err := u.useCase.GetReceivedFriendRequest(&listRequest)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.GetPendingListResponseModel
	for _, val := range result {
		fmt.Println("--",val.UniqueFriendID)
		if val != nil {
			finalResult = append(finalResult, &pb.GetPendingListResponseModel{
				FriendID:        val.FriendID,
				UpdateAt:        val.UpdateAt.String(),
				UniqueFriendsID: val.UniqueFriendID,
				UserID:          val.UserProfile.UserID,
				UserName:        val.UserProfile.UserName,
				Name:            val.UserProfile.Name,
				ProfilePhoto:    val.UserProfile.ProfilePhoto,
			})
		}
	}
	return &pb.GetReceivedFriendRequestResponse{Received: finalResult}, nil

}

func (u *FriendServer) GetSendFriendRequest(ctx context.Context, req *pb.GetSendFriendRequestRequest) (*pb.GetSendFriendRequestResponse, error) {
	var listSend requestmodel_friend_server.GetFriendRequest
	listSend.Limit = req.Send.Limit
	listSend.OffSet = req.Send.OffSet
	listSend.UserID = req.Send.UserID

	result, err := u.useCase.GetSendFriendRequest(&listSend)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.GetPendingListResponseModel
	for _, val := range result {
		if val != nil {
			finalResult = append(finalResult, &pb.GetPendingListResponseModel{
				FriendID:        val.FriendID,
				UpdateAt:        val.UpdateAt.String(),
				UniqueFriendsID: val.UniqueFriendID,
				UserID:          val.UserProfile.UserID,
				UserName:        val.UserProfile.UserName,
				Name:            val.UserProfile.Name,
				ProfilePhoto:    val.UserProfile.ProfilePhoto,
			})
		}
	}
	return &pb.GetSendFriendRequestResponse{Send: finalResult}, nil
}

func (u *FriendServer) GetBlockFriendRequest(ctx context.Context, req *pb.GetBlockFriendRequestRequest) (*pb.GetBlockFriendRequestResponse, error) {
	var listBlock requestmodel_friend_server.GetFriendRequest
	listBlock.Limit = req.Block.Limit
	listBlock.OffSet = req.Block.OffSet
	listBlock.UserID = req.Block.UserID

	result, err := u.useCase.GetBlockFriendRequest(&listBlock)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.GetPendingListResponseModel
	for _, val := range result {
		if val != nil {
			finalResult = append(finalResult, &pb.GetPendingListResponseModel{
				FriendID:        val.FriendID,
				UpdateAt:        val.UpdateAt.String(),
				UniqueFriendsID: val.UniqueFriendID,
				UserID:          val.UserProfile.UserID,
				UserName:        val.UserProfile.UserName,
				Name:            val.UserProfile.Name,
				ProfilePhoto:    val.UserProfile.ProfilePhoto,
			})
		}
	}
	return &pb.GetBlockFriendRequestResponse{Block: finalResult}, nil
}

func (u *FriendServer) UpdateFriendshipStatus(ctx context.Context, req *pb.UpdateFriendshipStatusRequest) (*emptypb.Empty, error) {
	err := u.useCase.FriendShipStatusUpdate(req.FriendShipID, req.Status)
	if err != nil {
		return new(emptypb.Empty), err
	}
	return new(emptypb.Empty), nil
}
