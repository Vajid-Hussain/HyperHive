run:
	go run cmd/main.go

protoAuthService:
	protoc ./pkg/pb/auth.proto --go_out=./ --go-grpc_out=./

autoserver:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go

mock:
	mockgen -source=pkg/repository/interface/interface_repo.go -destination=pkg/mock/mockRepo/user.go -package=mockrepository
	mockgen -source=pkg/usecase/interface/interface.go -destination=pkg/mock/mockUseCase/redis.go interface_usecase_auth_server IAuthCache 
