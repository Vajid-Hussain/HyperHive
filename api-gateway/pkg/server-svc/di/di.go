package di_server_svc

import (
	"fmt"

	middlewire_auth_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/auth-svc/infrastructure/middlewire"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/config"
	clind_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/clind"
	handler_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/handler"
	router_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/router"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

func InitServerClind(engin *echo.Echo, config *config.Config, middleWire *middlewire_auth_svc.Middlewire) error {
	serverClind, err := clind_server_svc.InitClind(config.Server_service_port)
	if err != nil {
		return err
	}

	soketioServer := InitSoketio()
	serverHandler := handler_server_svc.NewServerService(serverClind.Clind, soketioServer)

	router_server_svc.ServerRouter(engin.Group("/server"), serverHandler, middleWire,soketioServer)

	return nil
}

func InitSoketio() *socketio.Server {
	server := socketio.NewServer(nil)
	go server.Serve()

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:=", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed =>", reason)
	})
	return server
}
