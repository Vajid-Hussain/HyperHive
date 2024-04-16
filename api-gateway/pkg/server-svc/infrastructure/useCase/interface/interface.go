package interface_server_svc

import socketio "github.com/googollee/go-socket.io"

type IserverServiceUseCase interface {
	JoinToServerRoom(string, *socketio.Server, socketio.Conn) error
	EmitErrorMessage(socketio.Conn,  string)
	BroadcastMessage ( []byte,  *socketio.Server)
}
