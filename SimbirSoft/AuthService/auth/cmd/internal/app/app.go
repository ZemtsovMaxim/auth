package app

import (
	"log/slog"
	"time"

	grpcapp "gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/app/grpc"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/services/auth"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/storage/postgresql"
)

type App struct {
	GRPCDSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// инициализация хранилища

	// инициализация auth

	storage, err := postgresql.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCDSrv: grpcApp,
	}
}
