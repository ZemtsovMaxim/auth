package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/app"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/metrics"
	"gitlab.simbirsoft/verify/m.zemtsov/auth/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// инициализация конфига

	cfg := config.MustLoad()

	fmt.Println(cfg)

	// инициализация логгера

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.Any("config", cfg))

	// инициализация приложения

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	// запустить gRPC сервер
	go application.GRPCDSrv.MustRun()

	//запускаем сервер по сбору метрик
	go func() {
		if err := metrics.Listen("0.0.0.0:8082"); err != nil {
			log.Error("Failed to start metrics server: %v", err)
		}
		log.Info("Serving metrics at :8082/metrics")
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCDSrv.Stop()

	log.Info("Gracefully stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
