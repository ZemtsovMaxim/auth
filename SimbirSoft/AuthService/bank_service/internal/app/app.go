package app

import (
	grpcapp "bank_service/internal/app/grpc"
	"bank_service/internal/service/bank"
	"bank_service/internal/storage/postgres"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcNetwork, grpcAddress, storageDriver, storageInfo, storageURL, migrationPath string) (*App, error) {

	storage, err := postgres.New(storageDriver, storageInfo)
	if err != nil {
		return nil, err
	}

	err = storage.MigrationUp(storageURL, migrationPath)
	if err != nil {
		return nil, err
	}

	bank := bank.New(log, storage)

	grpcApp := grpcapp.New(log, bank, grpcNetwork, grpcAddress)

	return &App{
		GRPCServer: grpcApp,
	}, nil
}
