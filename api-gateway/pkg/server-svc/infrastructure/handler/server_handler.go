package handler_server_svc

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	requestmodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/requestModel"
	resonsemodel_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/model/resonseModel"
	interface_server_svc "github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/infrastructure/useCase/interface"
	"github.com/Vajid-Hussain/HiperHive/api-gateway/pkg/server-svc/pb"
	helper_api_gateway "github.com/Vajid-Hussain/HiperHive/api-gateway/utils"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

type ServerService struct {
	Clind         pb.ServerClient
	SoketioServer *socketio.Server
	serverUseCase interface_server_svc.IserverServiceUseCase
}

func NewServerService(clind pb.ServerClient, soketioServer *socketio.Server, serverUseCase interface_server_svc.IserverServiceUseCase) *ServerService {
	return &ServerService{
		Clind:         clind,
		SoketioServer: soketioServer,
		serverUseCase: serverUseCase,
	}
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
	result.OnlineUsers = strconv.Itoa(c.SoketioServer.RoomLen("/", result.ServerId))
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) SearchServer(ctx echo.Context)error{
	var req requestmodel_server_svc.SearchServer
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.SearchServer(context, &pb.SearchServerRequest{
		ServerID: req.ServerID,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) InitSoketio(ctx echo.Context) {

	c.SoketioServer.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		fmt.Println("connected:=", conn.ID())
		err := c.serverUseCase.JoinToServerRoom(ctx.Get("userID").(string), c.SoketioServer, conn)
		if err != nil {
			conn.Emit("error", err.Error())
		}
		return nil
	})

	c.SoketioServer.OnEvent("/", "friendly chat", func(conn socketio.Conn, msg string) {
		c.serverUseCase.SendFriendChat([]byte(msg), c.SoketioServer, conn)
	})

	c.SoketioServer.OnEvent("/", "server chat", func(s socketio.Conn, msg string) {
		c.serverUseCase.BroadcastMessage([]byte(msg), c.SoketioServer, s)
	})

	c.SoketioServer.OnEvent("/", "forum", func(conn socketio.Conn, msg string) {
		c.serverUseCase.BroadcastForum([]byte(msg), *c.SoketioServer, conn)
	})

	c.SoketioServer.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		fmt.Println("closed =>", reason)
		c.SoketioServer.LeaveAllRooms("/", conn)
	})
}

func (c *ServerService) GetChannelMessage(ctx echo.Context) error {
	var req requestmodel_server_svc.ChatRequest
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetChannelMessage(context, &pb.GetChannelMessageRequest{
		ChannelID: req.ChannelID,
		OffSet:    req.Offset,
		Limit:     req.Limit,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))

}

func (c *ServerService) UpdateServerPhoto(ctx echo.Context) error {
	var validImageExtention = map[string]struct{}{}

	validImageExtention["image/jpb"] = struct{}{}
	validImageExtention["image/png"] = struct{}{}
	validImageExtention["image/gif"] = struct{}{}

	file, err := ctx.FormFile("Image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, resonsemodel_server_svc.ErrNoImageInRequest.Error(), "", err.Error()))
	}

	if file.Size/(1024) > 1024 {
		return ctx.JSON(http.StatusRequestEntityTooLarge, resonsemodel_server_svc.Responses(http.StatusRequestEntityTooLarge, "", "", resonsemodel_server_svc.ErrImageOverSize.Error()))
	}

	// if _, ok := validImageExtention[file.Header.Get("Content-Type")]; !ok {
	// 	return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", resonsemodel_server_svc.ErrUnsupportImageType.Error()))
	// }

	image, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	buffer := make([]byte, file.Size)
	image.Read(buffer)

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = c.Clind.UpdateServerPhoto(context, &pb.UpdateServerPhotoRequest{
		Image:    buffer,
		UserID:   ctx.Get("userID").(string),
		ServerID: ctx.FormValue("ServerID"),
		Type:     ctx.FormValue("Type"),
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}

	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", ctx.FormValue("Type")+resonsemodel_server_svc.ServerImageUpdateSuccesFully, nil))
}

func (c *ServerService) UpdateServerDescription(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerDescription
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := c.Clind.UpdateServerDiscription(context, &pb.UpdateServerDiscriptionRequest{
		UserID:      ctx.Get("userID").(string),
		ServerID:    req.ServerID,
		Description: req.Description,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, resonsemodel_server_svc.ServerDescriptionUpdate, "", nil))
}

func (c *ServerService) GetServerMembers(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerMember
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetServerMembers(context, &pb.GetServerMembersRequest{
		ServerID: req.ServerID,
		OffSet:   req.Offset,
		Limit:    req.Limit,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))

}

func (c *ServerService) RemoveUserFromServer(ctx echo.Context) error {
	var req requestmodel_server_svc.RemoveUser
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := c.Clind.RemoveUserFromServer(context, &pb.RemoveUserFromServerRequest{
		UserID:    ctx.Get("userID").(string),
		RemoverID: req.RemoveUserID,
		ServerID:  req.ServerID},
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", "succesfully remove user "+req.RemoveUserID, nil))
}

func (c *ServerService) UpdateMemberRole(ctx echo.Context) error {
	var req requestmodel_server_svc.UpdateMemberRole
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := c.Clind.UpdateMemberRole(context, &pb.UpdateMemberRoleRequest{
		UserID:       ctx.Get("userID").(string),
		TargetUserID: req.TargetUserID,
		TargetRole:   req.TargetRole,
		ServerID:     req.ServerID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", "update sussefully to "+req.TargetRole, nil))
}

func (c *ServerService) DeleteServer(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerReq
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := c.Clind.DeleteServer(context, &pb.DeleteServerRequest{UserID: ctx.Get("userID").(string), ServerID: req.ServerID})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", "delete sussefully "+req.ServerID, nil))
}

func (c *ServerService) LeftFromServer(ctx echo.Context) error {
	var req requestmodel_server_svc.ServerReq
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := c.Clind.LeftFromServer(context, &pb.LeftFromServerRequest{UserID: ctx.Get("userID").(string), ServerID: req.ServerID})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", "left sussefully from "+req.ServerID, nil))
}

func (c *ServerService) GetForumPost(ctx echo.Context) error {
	var req requestmodel_server_svc.ReqGetForumPost
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetForumPost(context, &pb.GetForumPostRequest{
		ChannelID: req.ChannelID,
		Limit:     req.Limit,
		Offset:    req.Offset,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) GetSinglePost(ctx echo.Context) error {
	postID := ctx.Param("postid")
	if postID == "" {
		return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", "", resonsemodel_server_svc.ErrNoPostIDINQueryParams))
	}
	fmt.Println("===", postID)

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetSingleForumPost(context, &pb.GetSingleForumPostRequest{
		PostID: postID,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) GetPostCommand(ctx echo.Context) error {
	var req requestmodel_server_svc.ReqGetForumCommand
	ctx.Bind(&req)
	validateError := helper_api_gateway.Validator(req)
	if len(validateError) > 0 {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", validateError))
	}

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetPostCommand(context, &pb.GetPostCommandRequest{
		PostID: req.PostID,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}

func (c *ServerService) CheckReddis(ctx echo.Context) error{
	c.serverUseCase.SetDataInReddis()
	return ctx.String(200, "all perfect")
}