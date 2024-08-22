.PHONY: generate_proto build_container run_container docker_compose_build docker_compose_up

generate_proto:
	protoc -I api/bank api/bank/bank.proto --go_out=./api/gen/bank --go_opt=paths=source_relative --go-grpc_out=./api/gen/bank --go-grpc_opt=paths=source_relative

build_container:
	docker build . -t go-containerized:latest

run_container:
	docker run -p 0.0.0.0:8000:8000 go-containerized:latest

docker_compose_build:
	docker-compose up --build bank_service

docker_compose_up:
	docker-compose up bank_service