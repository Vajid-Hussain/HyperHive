autoserver:
	nodemon --watch './api-gateway/**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go
	nodemon --watch './auth-service/**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go

deploy:
	docker compose up --detach

deploydown:
	docker compose down

dockerPostgres:
	docker exec -it hyperhive_postgres bash -c 'psql -U postgres'


swaggo:
	swag init -g ./api-gateway/cmd/main.go


dockerall:
	docker compose down
	docker compose build
	docker compose up