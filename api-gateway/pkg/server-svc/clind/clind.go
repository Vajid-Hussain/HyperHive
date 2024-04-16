package clind_server_svc

import (
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	server "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServerClind struct {
	Clind server.ServerClient
}

type AuthClind struct {
	Clind pb.AuthServiceClient
}

func InitClind(port string) (*ServerClind, error) {
	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServerClind{Clind: server.NewServerClient(cc)}, nil
}

func InitAuthClind(port string) (*AuthClind, error) {
	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &AuthClind{Clind: pb.NewAuthServiceClient(cc)}, nil
}
