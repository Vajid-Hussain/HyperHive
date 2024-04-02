package server_friend_server

import (
	"context"
	"fmt"

	requestmodel_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/infrastructure/model/requestModel"
	"github.com/Vajid-Hussain/HyperHive/friend-service/pkg/pb"
	interface_usecase_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/usecase/interface"
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
	var req reque
	result, err := u.useCase.GetFriendRequest(req.UserID)
	if err != nil {
		return nil, err
	}

	for _, val := range result {
		fmt.Println("--", val)
	}
	// return &pb.FriendListResponse{
	// 	FriendID: result.FriendsID,
	// }

	return nil, nil
}
