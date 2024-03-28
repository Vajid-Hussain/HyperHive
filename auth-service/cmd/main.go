package main

import (
	"fmt"
	"log"
	"net"

	configl_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/config"
	dil_auth_server "github.com/Vajid-Hussain/HiperHive/auth-service/pkg/di"
	"github.com/Vajid-Hussain/HiperHive/auth-service/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	defer handlepanic() 
	config, err := configl_auth_server.InitServer()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.DB.Port)
	grpcServer := grpc.NewServer()

	server, err := dil_auth_server.InitAuthServer(config)
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterAuthServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("failed to serve", err)
	}
}

func handlepanic() {

	if a := recover(); a != nil {

		fmt.Println("RECOVER", a)
	}
}
