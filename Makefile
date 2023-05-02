include .env

TEST_DB="host=$(DB_HOST) port=$(DB_PORT) user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_TEST_NAME) sslmode=$(DB_SSL_MODE)"
CMD= cmd/main.go

db:
		psql -c "drop database if exists $(DB_TEST_NAME)"
		psql -c "create database $(DB_TEST_NAME)"
		goose -dir migrations postgres $(TEST_DB) up

test: 	db
		go test ./...

protogen:
		protoc --go_out=internal/protobuf \
               --go-grpc_out=internal/protobuf \
               proto/shorter.proto

postgresql:
		docker-compose --profile postgresql up --build


inmemory:
		docker-compose --profile inmemory up --build

linter:
		golangci-lint --config .golangci.yaml run

.PHONY:	db test protogen postgresql inmemory linter
