package client

import (
	api "gitlab.simbirsoft/verify/m.zemtsov/auth/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientGRPC struct {
	api.AuthClient
}

func NewClientGRPC(address string) (*ClientGRPC, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := api.NewAuthClient(conn)
	return &ClientGRPC{
		client,
	}, nil
}
