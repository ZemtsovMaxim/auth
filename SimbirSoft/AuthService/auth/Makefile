.PHONY: generate migrate lint test

generate:
	protoc -I api/service api/service/auth.proto --go_out=api/gen --go_opt=paths=source_relative --go-grpc_out=api/gen --go-grpc_opt=paths=source_relative 

migrate:
	go run ./cmd/migrator --migrations-path=./migrations --storage-path "postgres://myUser:12345@db:5432/myDb?sslmode=disable"

test:
	go test -v ./tests		

# Предварительно установите golangci-lint https://golangci-lint.run
lint:
	golangci-lint run 

