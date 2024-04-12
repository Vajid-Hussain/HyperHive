package main

import (
	"log"
	"net"

	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
	di_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/di"
	"github.com/Vajid-Hussain/HyperHive/server-service/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	config, err := config_server_service.ConfigInit()
	if err != nil {
		log.Fatal(err)
	}

	server, err := di_server_service.ServerInitialize(config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.ServerCredential.ServerPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServerServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}
