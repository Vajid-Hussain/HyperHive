package clind_server_svc

import (
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServerClind struct {
	Clind pb.ServerClient
}

func InitClind(port string) (*ServerClind, error) {
	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &ServerClind{Clind: pb.NewServerClient(cc)}, nil
}
