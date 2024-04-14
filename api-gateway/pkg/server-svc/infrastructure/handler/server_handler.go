package handler_server_svc

import (
	"context"
	"net/http"
	"time"

	requestmodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/requestModel"
	resonsemodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/resonseModel"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
	"github.com/labstack/echo/v4"
)

type ServerService struct {
	Clind pb.ServerClient
}

func NewServerService(clind pb.ServerClient) *ServerService {
	return &ServerService{Clind: clind}
}

func (c *ServerService) CreateServer(ctx echo.Context) error {
	var serverReq requestmodel_server_svc.Server

	err := ctx.Bind(&serverReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(serverReq)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.CreateServer(context, &pb.CreateServerRequest{ServerName: serverReq.Name, UserID: ctx.Get("userID").(string)})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "server create succesfully", result, nil))
}

func (c *ServerService) CreateCategory(ctx echo.Context) error {
	var req requestmodel_server_svc.CreateCategory
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err = c.Clind.CreateCategory(context, &pb.CreateCategoryRequest{
		UserID:       ctx.Get("userID").(string),
		ServerID:     req.ServerID,
		CategoryName: req.CategoryName,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "category succesfully created", "", nil))
}

func (c *ServerService) CreateChannel(ctx echo.Context) error {
	var req requestmodel_server_svc.CreateChannel
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = c.Clind.CreateChannel(context, &pb.CreateChannelRequest{
		UserID:      ctx.Get("userID").(string),
		ServerID:    req.ServerID,
		CategoryID:  req.CategoryID,
		ChannelName: req.ChannelName,
		ChannelType: req.Type,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "channel succesfully created", "", nil))
}

func (c *ServerService) JoinToServer(ctx echo.Context) error {
	var req requestmodel_server_svc.JoinToServer
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = c.Clind.JoinToServer(context, &pb.JoinToServerRequest{
		UserID:   ctx.Get("userID").(string),
		ServerID: req.ServerID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "Now you are a member", "", nil))
}

func (c *ServerService) GetCategoryOfServer(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerReq
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.GetCategoryOfServer(context, &pb.GetCategoryOfServerRequest{ServerID: req.ServerID})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) GetChannelsOfServer(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerReq
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, resonsemodel_server_svc.Responses(http.StatusUnsupportedMediaType, "", "", err.Error()))
	}

	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.GetChannelsOfServer(context, &pb.GetChannelsOfServerRequest{ServerID: req.ServerID})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) GetUserServer(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetUserServer(context, &pb.GetUserServerRequest{UserID: ctx.Get("userID").(string)})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) GetServer(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := c.Clind.GetServer(context, &pb.GetServerRequest{ServerID: ctx.Param("id")})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

// func (c *ServerService) SoketIO(ctx echo.Context) error {
// 	fmt.Println("called=====================", ctx.Request())
// 	server := socketio.NewServer(nil)

// 	server.OnConnect("/", func(s socketio.Conn) error {
// 		s.SetContext("")
// 		fmt.Println("connected:", s.ID())
// 		return nil
// 	})

// 	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
// 		fmt.Println("notice:", msg)
// 		s.Emit("reply", "have "+msg)
// 	})

// 	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
// 		s.SetContext(msg)
// 		return "recv " + msg
// 	})

// 	server.OnEvent("/", "bye", func(s socketio.Conn) string {
// 		last := s.Context().(string)
// 		s.Emit("bye", last)
// 		s.Close()
// 		return last
// 	})

// 	server.OnError("/", func(s socketio.Conn, e error) {
// 		// server.Remove(s.ID())
// 		fmt.Println("meet error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {

// 		// Add the Remove session id. Fixed the connection & mem leak
// 		// server.Remove(s.ID())
// 		fmt.Println("closed =>", reason)
// 	})

// 	return errors.New("i dont know what is the errro")
// }

// server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
// 	log.Println("notice:", msg)
// 	s.Emit("reply", "have "+msg)
// })

// server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
// 	s.SetContext(msg)
// 	return "recv " + msg
// })

// server.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
// 	s.Emit("echo", msg)
// })

// server.OnEvent("/", "bye", func(s socketio.Conn) string {
// 	last := s.Context().(string)
// 	s.Emit("bye", last)
// 	s.Close()
// 	return last
// })

// server.OnError("/", func(s socketio.Conn, e error) {
// 	log.Println("meet error:", e)
// })

// server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 	log.Println("closed", reason)
// })
