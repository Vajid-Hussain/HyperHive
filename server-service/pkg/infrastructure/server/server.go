package server_server_service

import (
	"context"

	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
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
