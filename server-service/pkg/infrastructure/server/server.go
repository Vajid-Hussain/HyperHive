package server_server_service

import (
	"context"
	"strconv"

	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	responsemodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/responseModel"
	"github.com/Vajid-Hussain/HyperHive/server-service/pkg/pb"
	interface_useCase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase/interface"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerServer struct {
	useCase interface_useCase_server_service.IServerUseCase
	pb.UnimplementedServerServer
}

func NewServerServer(usecase interface_useCase_server_service.IServerUseCase) *ServerServer {
	return &ServerServer{useCase: usecase}
}

func (u *ServerServer) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerResponse, error) {
	result, err := u.useCase.CreateServer(&requestmodel_server_service.Server{Name: req.ServerName, UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	return &pb.CreateServerResponse{
		ServerID:   result.ServerID,
		ServerName: result.Name,
	}, nil
}

func (u *ServerServer) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*emptypb.Empty, error) {
	categoryData := requestmodel_server_service.CreateCategory{
		UserID:       req.UserID,
		ServerID:     req.ServerID,
		CategoryName: req.CategoryName,
	}

	return new(emptypb.Empty), u.useCase.CreateCategory(&categoryData)
}

func (u *ServerServer) CreateChannel(ctx context.Context, req *pb.CreateChannelRequest) (*emptypb.Empty, error) {
	channelData := requestmodel_server_service.CreateChannel{
		ChannelName: req.ChannelName,
		UserID:      req.UserID,
		ServerID:    req.ServerID,
		CategoryID:  req.CategoryID,
		Type:        req.ChannelType,
	}

	return new(emptypb.Empty), u.useCase.CreateChannel(&channelData)
}

func (u *ServerServer) JoinToServer(ctx context.Context, req *pb.JoinToServerRequest) (*emptypb.Empty, error) {
	joinServer := requestmodel_server_service.JoinToServer{
		UserID:   req.UserID,
		ServerID: req.ServerID,
	}
	return new(emptypb.Empty), u.useCase.JoinToServer(&joinServer)
}

func (u *ServerServer) GetCategoryOfServer(ctx context.Context, req *pb.GetCategoryOfServerRequest) (*pb.GetCategoryOfServerResponse, error) {
	result, err := u.useCase.GetServerCategory(req.ServerID)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.FullServerChannel
	for _, val := range result {
		finalResult = append(finalResult, &pb.FullServerChannel{CategoryID: val.CategoryID, Name: val.Name})
	}

	return &pb.GetCategoryOfServerResponse{
		Data: finalResult,
	}, nil
}

func (u *ServerServer) GetChannelsOfServer(ctx context.Context, req *pb.GetChannelsOfServerRequest) (*pb.GetChannelsOfServerResponse, error) {
	result, err := u.useCase.GetChannels(req.ServerID)
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.FullServerChannel
	for _, val := range result {
		var allChannel []*pb.Channel
		for _, val := range val.Channel {
			allChannel = append(allChannel, &pb.Channel{ChannelID: val.ChannelID, CategoryID: val.CategoryID, Name: val.Name, Type: val.Type})
		}
		finalResult = append(finalResult, &pb.FullServerChannel{CategoryID: val.CategoryID, Name: val.Name, Channel: allChannel})
	}

	return &pb.GetChannelsOfServerResponse{Data: finalResult}, nil
}

func (u *ServerServer) GetUserServer(ctx context.Context, req *pb.GetUserServerRequest) (*pb.GetUserServerResponse, error) {
	result, err := u.useCase.GetUserServers(req.UserID)
	if err != nil {
		return nil, err
	}

	var finalResult []string
	for _, val := range result {
		finalResult = append(finalResult, val.ServerID)
	}
	return &pb.GetUserServerResponse{ServerId: finalResult}, nil
}

func (u *ServerServer) GetServer(ctx context.Context, req *pb.GetServerRequest) (*pb.GetServerResponse, error) {
	result, err := u.useCase.GetServer(req.ServerID)
	if err != nil {
		return nil, err
	}

	return &pb.GetServerResponse{
		ServerId:    req.ServerID,
		Name:        result.Name,
		Description: result.Description,
		Icon:        result.Icon,
		CoverPhoto:  result.CoverPhoto,
	}, nil
}

func (u *ServerServer) GetChannelMessage(ctx context.Context, req *pb.GetChannelMessageRequest) (*pb.GetChannelMessageResponse, error) {
	result, err := u.useCase.GetChannelMessages(req.ChannelID, requestmodel_server_service.Pagination{Limit: req.Limit, OffSet: req.OffSet})
	if err != nil {
		return nil, err
	}

	var finalResult []*pb.ChannelMessage
	for _, val := range result {
		var msg pb.ChannelMessage
		msg.UserProfile= val.UserProfile
		msg.UserName= val.UserName
		msg.ChannelID = strconv.Itoa(val.ChannelID)
		msg.Content = val.Content
		msg.ID = val.ID
		msg.ServerID = strconv.Itoa(val.ServerID)
		msg.TimeStamp = val.TimeStamp.String()
		msg.Type = val.Type
		msg.UserID = strconv.Itoa(val.UserID)
		finalResult = append(finalResult, &msg)
	}

	return &pb.GetChannelMessageResponse{Messages: finalResult}, nil
}

func (u *ServerServer) UpdateServerPhoto(ctx context.Context, req *pb.UpdateServerPhotoRequest) (*emptypb.Empty, error) {
	var request requestmodel_server_service.ServerImages
	request.Image = req.Image
	request.ServerID = req.ServerID
	request.Type = req.Type
	request.UserID = req.UserID

	err := u.useCase.UpdateServerPhoto(&request)
	if err != nil {
		return new(emptypb.Empty), err
	}
	return new(emptypb.Empty), nil
}

func (u *ServerServer) UpdateServerDiscription(ctx context.Context, req *pb.UpdateServerDiscriptionRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.useCase.UpdateServerDiscription(&requestmodel_server_service.Description{UserID: req.UserID, Description: req.Description, ServerID: req.ServerID})
}

func (u *ServerServer) GetServerMembers(ctx context.Context, req *pb.GetServerMembersRequest) (*pb.GetServerMembersResponse, error) {
	result, err := u.useCase.GetServerMembers(req.ServerID, requestmodel_server_service.Pagination{Limit: req.Limit, OffSet: req.OffSet})
	if err != nil {
		return nil, err
	}
	var finalResult []*pb.ServerMemberModel
	for _, member := range result {
		var memberDetails pb.ServerMemberModel
		memberDetails.Name = member.Name
		memberDetails.Role = member.Role
		memberDetails.UserId = member.UserID
		memberDetails.UserName = member.UserName
		memberDetails.UserProfile = member.UserProfile
		finalResult = append(finalResult, &memberDetails)
	}
	return &pb.GetServerMembersResponse{List: finalResult}, nil
}

func (u *ServerServer) UpdateMemberRole(ctx context.Context, req *pb.UpdateMemberRoleRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.useCase.UpdateMemberRole(requestmodel_server_service.UpdateMemberRole{
		UserID:       req.UserID,
		TargetUserID: req.TargetUserID,
		TargetRole:   req.TargetRole,
		ServerID:     req.ServerID,
	})
}

func (u *ServerServer) RemoveUserFromServer(ctx context.Context, req *pb.RemoveUserFromServerRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.useCase.RemoveUserFromServer(&requestmodel_server_service.RemoveUser{
		UserID:    req.UserID,
		RemoverID: req.RemoverID,
		ServerID:  req.ServerID,
	})
}

func (u *ServerServer) DeleteServer(ctx context.Context, req *pb.DeleteServerRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.useCase.DeleteServer(req.UserID, req.ServerID)
}

func (u *ServerServer) LeftFromServer(ctx context.Context, req *pb.LeftFromServerRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), u.useCase.LeftFromServer(req.UserID, req.ServerID)
}

func (u *ServerServer) GetForumPost(ctx context.Context, req *pb.GetForumPostRequest) (*pb.GetForumPostResponse, error) {
	result, err := u.useCase.GetForumPost(req.ChannelID, requestmodel_server_service.Pagination{
		Limit:  req.Limit,
		OffSet: req.Offset,
	})
	if err != nil {
		return nil, err
	}

	var post []*pb.GetForumPostModel
	for _, value := range result {
		post = append(post, &pb.GetForumPostModel{
			PostID:          value.ID,
			UserProfile:     value.UserProfile,
			UserName:        value.UserName,
			UserId:          int32(value.UserID),
			ChannelId:       int32(value.ChannelID),
			ServerId:        int32(value.ServerID),
			Content:         value.Content,
			MainContentType: value.MainContentType,
			SubContent:      value.SubContent,
			TimeStamp:       value.TimeStamp.String(),
			Type:            value.Type,
			CommandContent:  value.CommandContent,
		})
	}
	return &pb.GetForumPostResponse{Post: post}, nil
}

func (u *ServerServer) GetSingleForumPost(ctx context.Context, req *pb.GetSingleForumPostRequest) (*pb.GetSingleForumPostResponse, error) {
	value, err := u.useCase.GetFormSinglePost(req.PostID)
	if err != nil {
		return nil, err
	}
	return &pb.GetSingleForumPostResponse{
		Post: &pb.GetForumPostModel{
			PostID:          value.ID,
			UserProfile:     value.UserProfile,
			UserName:        value.UserName,
			UserId:          int32(value.UserID),
			ChannelId:       int32(value.ChannelID),
			ServerId:        int32(value.ServerID),
			Content:         value.Content,
			MainContentType: value.MainContentType,
			SubContent:      value.SubContent,
			TimeStamp:       value.TimeStamp.String(),
			Type:            value.Type,
		},
	}, nil
}

func (u *ServerServer) GetPostCommand(ctx context.Context, req *pb.GetPostCommandRequest) (*pb.GetPostCommandResponse, error) {
	result, err := u.useCase.GetPostCommand(req.PostID, requestmodel_server_service.Pagination{
		Limit:  req.Limit,
		OffSet: req.Offset,
	})
	if err != nil {
		return nil, err
	}

	// limitInt, _ := strconv.Atoi(req.Limit)

	var finalResult []*pb.ForumCommandModel
	// for _, val := range result.Commands {
	// 	finalResult = append(finalResult, &pb.ForumCommandModel{
	// 		ID:          val.ID,
	// 		UserProfile: val.UserProfile,
	// 		UserName:    val.UserName,
	// 		UserId:      int32(val.UserID),
	// 		ChannelId:   int32(val.ChannelID),
	// 		ServerId:    int32(val.ServerID),
	// 		ParentId:    val.ParentID,
	// 		Content:     val.Content,
	// 		TimeStamp:   val.TimeStamp.String(),
	// 		Type:        val.Type,
	// 		Thread:      val.Thread,
	// 	})
	// }

	finalResult = u.arrageAllCommand(result.Commands, finalResult)

	return &pb.GetPostCommandResponse{Command: finalResult}, err
}

func (u *ServerServer) arrageAllCommand(command []*responsemodel_server_service.ForumCommand, finalResult []*pb.ForumCommandModel) []*pb.ForumCommandModel {
	for i, val := range command {
		finalResult = append(finalResult, &pb.ForumCommandModel{
			ID:          val.ID,
			UserProfile: val.UserProfile,
			UserName:    val.UserName,
			UserId:      int32(val.UserID),
			ChannelId:   int32(val.ChannelID),
			ServerId:    int32(val.ServerID),
			ParentId:    val.ParentID,
			Content:     val.Content,
			TimeStamp:   val.TimeStamp.String(),
			Type:        val.Type,
		})
		finalResult[i].Thread = u.arrageAllCommand(command[i].Thread, finalResult[i].Thread)
	}

	return finalResult
}
