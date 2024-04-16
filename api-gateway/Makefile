run:
	go run cmd/main.go

protoAuthService:
	protoc ./pkg/auth-svc/pb/auth.proto --go_out=./ --go-grpc_out=./

autoserver:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go

protoFriendService:
	protoc ./pkg/friend-svc/pb/friend.proto --go_out=./ --go-grpc_out=./

protoServerService:
	protoc ./pkg/server-svc/pb/server.proto --go_out=./ --go-grpc_out=./

protoAuthServiceINServerService:
		protoc ./pkg/server-svc/pb/auth.proto --go_out=./ --go-grpc_out=./
