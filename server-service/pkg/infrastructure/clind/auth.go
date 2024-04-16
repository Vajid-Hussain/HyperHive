package clind_srv_on_server_service

import (
	"github.com/Vajid-Hussain/HyperHive/server-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClind struct {
	Clind pb.AuthServiceClient
}

func InitAuthClind(port string) (pb.AuthServiceClient, error) {
	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	clind := AuthClind{Clind: pb.NewAuthServiceClient(cc)}
	return clind.Clind, nil
}
