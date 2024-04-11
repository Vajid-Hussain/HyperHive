package main

import (
	"log"

	config_server_service "github.com/Vajid-Hussain/HyperHive/server-service/pkg/config"
)

func main(){
	config,err:= config_server_service.ConfigInit()
	if err!=nil{
		log.Fatal(err)
	}

	
}