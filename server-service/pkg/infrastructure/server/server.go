package server_server_service

import (
	"context"

	requestmodel_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/infrastructure/model/requestModel"
	"github.com/Vajid-Hussain/HyperHive/server-service/pkg/pb"
	interface_useCase_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/usecase/interface"
)

type ServerServer struct {
	useCase interface_useCase_server_service.IServerUseCase
	pb.UnimplementedServerServer
}

func NewServerServer(usecase interface_useCase_server_service.IServerUseCase) *ServerServer {
	return &ServerServer{useCase: usecase}
}

func (u *ServerServer) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerResponse, error) {
	result, err := u.useCase.CreateServer(&requestmodel_server_service.Server{Name: req.ServerName})
	if err != nil {
		return nil, err
	}
	return &pb.CreateServerResponse{
		ServerID:   result.ServerID,
		ServerName: result.Name,
	}, nil
}
