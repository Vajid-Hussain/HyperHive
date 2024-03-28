package clind_auth_svc

import (
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClind(port string) (pb.AuthServiceClient, error) {
	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewAuthServiceClient(cc), nil
}
