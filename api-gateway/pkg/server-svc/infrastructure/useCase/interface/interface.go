package interface_server_svc

import socketio "github.com/googollee/go-socket.io"

type IserverServiceUseCase interface {
	JoinToServerRoom(string, *socketio.Server, socketio.Conn) error
	BroadcastMessage( []byte, *socketio.Server, socketio.Conn)
	SendFriendChat( []byte, *socketio.Server, socketio.Conn)
}
