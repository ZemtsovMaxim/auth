package server

import (
	"context"
	"strings"

	api "gitlab.simbirsoft/verify/m.zemtsov/auth/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		idSec int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (statusMsg string, err error)
	Logout(ctx context.Context, token string) (invalidToken string, err error)
	ValidateToken(ctx context.Context, token string, idSec int) (id int, err error)
}

type serverAPI struct {
	api.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	api.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	if err := validateLogin(req); err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), 1) //TODO: Здесь заглушка тип 1 секрет выбираем ибо протошник надо менять
	if err != nil {
		if strings.Contains(err.Error(), "invalid credentials") {
			return nil, status.Errorf(codes.InvalidArgument, "Wrong email or password")
		} else if strings.Contains(err.Error(), "no rows in result set") {
			return nil, status.Errorf(codes.InvalidArgument, "Wrong email or password")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &api.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {

	if err := validateRegister(req); err != nil {
		return nil, err
	}

	statusMsg, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &api.RegisterResponse{
		StatusMessage: statusMsg,
	}, nil
}

func (s *serverAPI) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) {
	if err := validateLogout(req); err != nil {
		return nil, err
	}

	invalidToken, err := s.auth.Logout(ctx, req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &api.LogoutResponse{
		Token: invalidToken,
	}, nil
}

func (s *serverAPI) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {

	id, err := s.auth.ValidateToken(ctx, req.GetToken(), 1)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "invalid token")
	}

	return &api.ValidateTokenResponse{
		Id: int64(id),
	}, nil
}

func validateLogin(req *api.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Errorf(codes.InvalidArgument, "email is required")
	}

	if !strings.Contains(req.GetEmail(), "@") {
		return status.Errorf(codes.InvalidArgument, "incorrect email")
	}

	if req.GetPassword() == "" {
		return status.Errorf(codes.InvalidArgument, "password is required")
	}
	if len(req.GetPassword()) <= 4 {
		return status.Errorf(codes.InvalidArgument, "password should be bigger than 4 letters")
	}
	return nil
}

func validateRegister(req *api.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Errorf(codes.InvalidArgument, "email is required")
	}

	if !strings.Contains(req.GetEmail(), "@") {
		return status.Errorf(codes.InvalidArgument, "incorrect email")
	}

	if req.GetPassword() == "" {
		return status.Errorf(codes.InvalidArgument, "password is required")
	}
	if len(req.GetPassword()) <= 4 {
		return status.Errorf(codes.InvalidArgument, "password should be bigger than 4 letters")
	}
	return nil
}

func validateLogout(req *api.LogoutRequest) error {
	if req.GetToken() == "" {
		return status.Errorf(codes.InvalidArgument, "Token is missed")
	}

	return nil
}
