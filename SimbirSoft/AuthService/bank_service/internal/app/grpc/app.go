package grpcapp

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	server "bank_service/internal/server/grpc"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	GRPCServer *grpc.Server
	network    string
	address    string
}

func New(log *slog.Logger, bank server.Bank, network string, address string) *App {
	GRPCServer := grpc.NewServer()

	server.Register(GRPCServer, bank)

	return &App{
		log:        log,
		GRPCServer: GRPCServer,
		network:    network,
		address:    address,
	}
}

func (a *App) MustRun() {
	err := a.Run()
	if err != nil {
		log.Fatalf("app couldn't run: %s", err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.String("address", a.address),
	)

	listner, err := net.Listen(a.network, a.address)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("gRPC server is running", slog.String("address", listner.Addr().String()))

	err = a.GRPCServer.Serve(listner)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op))
	a.log.Info("stopping gRPC server", slog.String("address", a.address))

	a.GRPCServer.GracefulStop()
}
