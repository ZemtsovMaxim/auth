package client

import (
	auth_v1 "bank_service/api/gen/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientGRPC struct {
	auth_v1.AuthClient
}

func NewClientGRPC(address string) (*ClientGRPC, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := auth_v1.NewAuthClient(conn)

	return &ClientGRPC{client}, nil
}
