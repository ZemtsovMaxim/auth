package main

import (
	"bank_service/internal/app"
	"bank_service/internal/config"
	"bank_service/pkg/logger/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	log.Info("starting application", slog.Any("config", cfg))
	application, err := app.New(log, cfg.MyGRPC.Network, cfg.MyGRPC.Address, cfg.Storage.Driver, cfg.Storage.Info, cfg.Storage.URL, cfg.MigrationPath)
	if err != nil {
		log.Error("failed to init app", sl.Err(err))
		os.Exit(1)
	}

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("application stopped")
}
