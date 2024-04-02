package main

import (
	"log"
	"net"

	config_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/config"
	di_friend_server "github.com/Vajid-Hussain/HyperHive/friend-service/pkg/di"
	"github.com/Vajid-Hussain/HyperHive/friend-service/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config_friend_server.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := di_friend_server.InitFriendService(config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.Friend.Friend_Service_Port)
	grpcServer := grpc.NewServer()

	pb.RegisterFriendServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
