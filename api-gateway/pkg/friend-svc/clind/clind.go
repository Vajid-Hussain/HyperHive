package clind_friend_svc

import (
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/friend-svc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clind struct {
	Clind pb.FriendServiceClient
}

func InitClind(port string) (*Clind, error) {
	cc, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Clind{Clind: pb.NewFriendServiceClient(cc)}, nil
}
