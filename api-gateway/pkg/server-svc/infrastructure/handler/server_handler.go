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

// Handler function for creating a server.
// @Summary Create Server
// @Description Create a new server.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.Server true "Request body for creating a server"
// @Success 201 {object} resonsemodel_server_svc.Response "Server created successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server [post]
// @Security UserAuthorization
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

// Handler function for creating a category.
// @Summary Create Category
// @Description Create a new category.
// @Tags Category
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.CreateCategory true "Request body for creating a category"
// @Success 201 {object} resonsemodel_server_svc.Response "Category created successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/category [post]
// @Security UserAuthorization
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

// Handler function for creating a channel.
// @Summary Create Channel
// @Description Create a new channel.
// @Tags Channel
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.CreateChannel true "Request body for creating a channel"
// @Success 201 {object} resonsemodel_server_svc.Response "Channel created successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/channel [post]
// @Security UserAuthorization
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

// Handler function for joining a server.
// @Summary Join Server
// @Description Join an existing server.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.JoinToServer true "Request body for joining a server"
// @Success 200 {object} resonsemodel_server_svc.Response "Joined server successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /join [post]
// @Security UserAuthorization
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

// Handler function for getting categories of a server.
// @Summary Get Categories of Server
// @Description Retrieve categories of a server.
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Server ID" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} resonsemodel_server_svc.Response "Categories retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/category [get]
// @Security UserAuthorization
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

// Handler function for getting channels of a server.
// @Summary Get Channels of Server
// @Description Retrieve channels of a server.
// @Tags Channel
// @Accept json
// @Produce json
// @Param id path string true "Server ID" format:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"
// @Success 200 {object} resonsemodel_server_svc.Response "Channels retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/channel [get]
// @Security UserAuthorization

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

// Handler function for getting a user's server.
// @Summary Get User's Server
// @Description Retrieve a user's server information.
// @Tags Server
// @Accept json
// @Produce json
// @Success 200 {object} resonsemodel_server_svc.Response "User's server information retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/userserver [get]
// @Security UserAuthorization
func (c *ServerService) GetUserServer(ctx echo.Context) error {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := c.Clind.GetUserServer(context, &pb.GetUserServerRequest{UserID: ctx.Get("userID").(string)})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, resonsemodel_server_svc.Responses(http.StatusBadRequest, "", "", err.Error()))
	}
	return ctx.JSON(http.StatusOK, resonsemodel_server_svc.Responses(http.StatusOK, "", result, nil))
}


// Handler function for getting a server by ID.
// @Summary Get Server by ID
// @Description Retrieve a server by ID.
// @Tags Server
// @Accept json
// @Produce json
// @Param id path string true "Server ID"
// @Success 200 {object} resonsemodel_server_svc.Response "Server information retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/{id} [get]
// @Security UserAuthorization
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

// Handler function for searching servers.
// @Summary Search Servers
// @Description Search servers based on query parameters.
// @Tags Server
// @Accept json
// @Produce json
// @Param limit query string false "Limit for pagination"
// @Param page query string false "Offset for pagination"
// @Param name query string false "Server name"
// @Success 200 {object} resonsemodel_server_svc.Response "Servers retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/search [get]
// @Security UserAuthorization
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

// Handler function for getting channel messages.
// @Summary Get Channel Messages
// @Description Retrieve messages of a channel.
// @Tags Channel
// @Accept json
// @Produce json
// @Param ChannelID query string true "Channel ID"
// @Param Page query string true "Offset for pagination"
// @Param Limit query string true "Limit for pagination"
// @Success 200 {object} resonsemodel_server_svc.Response "Messages retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/message [get]
// @Security UserAuthorization
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

// Handler function for updating a server's photo.
// @Summary Update Server Photo
// @Description Update a server's photo.
// @Tags Server
// @Accept multipart/form-data
// @Produce json
// @Param Image formData file true "Image file"
// @Success 200 {object} resonsemodel_server_svc.Response "Server photo updated successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/image [patch]
// @Security UserAuthorization
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

// Handler function for updating a server's description.
// @Summary Update Server Description
// @Description Update a server's description.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.ServerDescription true "Request body for updating server description"
// @Success 200 {object} resonsemodel_server_svc.Response "Server description updated successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/description [patch]
// @Security UserAuthorization
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

// Handler function for getting server members.
// @Summary Get Server Members
// @Description Retrieve members of a server.
// @Tags Server
// @Accept json
// @Produce json
// @Param ServerID query string true "Server ID"
// @Param Page query string true "Offset for pagination"
// @Param Limit query string true "Limit for pagination"
// @Success 200 {object} resonsemodel_server_svc.Response "Server members retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/members [get]
// @Security UserAuthorization
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

// Handler function for removing a user from a server.
// @Summary Remove User from Server
// @Description Remove a user from a server.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.RemoveUser true "Request body for removing user from server"
// @Success 200 {object} resonsemodel_server_svc.Response "User removed from server successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/remove [delete]
// @Security UserAuthorization
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

// Handler function for updating a member's role in a server.
// @Summary Update Member Role
// @Description Update a member's role in a server.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.UpdateMemberRole true "Request body for updating member role"
// @Success 200 {object} resonsemodel_server_svc.Response "Member role updated successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/role [patch]
// @Security UserAuthorization
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

// Handler function for deleting a server.
// @Summary Delete Server
// @Description Delete a server.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.ServerReq true "Request body for deleting server"
// @Success 200 {object} resonsemodel_server_svc.Response "Server deleted successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/destroy [delete]
// @Security UserAuthorization
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

// Handler function for leaving a server.
// @Summary Leave Server
// @Description Leave a server.
// @Tags Server
// @Accept json
// @Produce json
// @Param body body requestmodel_server_svc.ServerReq true "Request body for leaving server"
// @Success 200 {object} resonsemodel_server_svc.Response "Left from server successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/left [delete]
// @Security UserAuthorization
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

// Handler function for getting forum posts.
// @Summary Get Forum Posts
// @Description Retrieve forum posts.
// @Tags Forum
// @Accept json
// @Produce json
// @Param limit query string false "Limit for pagination"
// @Param page query string false "Offset for pagination"
// @Param channelID query string false "Channel ID"
// @Success 200 {object} resonsemodel_server_svc.Response "Forum posts retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/forum [get]
// @Security UserAuthorization
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

// Handler function for getting a single forum post.
// @Summary Get Single Forum Post
// @Description Retrieve a single forum post.
// @Tags Forum
// @Accept json
// @Produce json
// @Param postid path string true "Post ID"
// @Success 200 {object} resonsemodel_server_svc.Response "Single forum post retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/forum/{postid} [get]
// @Security UserAuthorization
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

// Handler function for getting forum commands.
// @Summary Get Forum Commands
// @Description Retrieve forum commands.
// @Tags Forum
// @Accept json
// @Produce json
// @Param limit query string false "Limit for pagination"
// @Param page query string false "Offset for pagination"
// @Param PostID query string false "Post ID"
// @Success 200 {object} resonsemodel_server_svc.Response "Forum commands retrieved successfully"
// @Failure 400 {object} resonsemodel_server_svc.Response "Bad request"
// @Router /server/forum/command [get]
// @Security UserAuthorization
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