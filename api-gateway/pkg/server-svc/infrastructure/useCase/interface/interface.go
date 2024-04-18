package interface_server_svc

import socketio "github.com/googollee/go-socket.io"

type IserverServiceUseCase interface {
	JoinToServerRoom(string, *socketio.Server, socketio.Conn) error
	// EmitErrorMessage(socketio.Conn,  string)
	BroadcastMessage(string, []byte, *socketio.Server,socketio.Conn)
	SendFriendChat(string, []byte,  *socketio.Server,  socketio.Conn)
}
