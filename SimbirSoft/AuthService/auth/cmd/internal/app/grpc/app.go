package grpcapp

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	server "gitlab.simbirsoft/verify/m.zemtsov/auth/cmd/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func unaryInterceptorLogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	log.Println(md)

	m, err := handler(ctx, req)
	logger(info.FullMethod)
	if err != nil {
		logger("RPC failed with error: %v", err)
	}
	return m, err
}

func New(log *slog.Logger, authService server.Auth, port int) *App {
	// Объединяем перехватчики в один с помощью ChainUnaryInterceptor
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			unaryInterceptorLogger,                 // перехватчик для логгирования
			grpc_prometheus.UnaryServerInterceptor, // перехватчик для метрик Prometheus
		),
	)

	// Регистрируем метрики
	grpc_prometheus.Register(grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	// Регистрируем сервисы
	server.Register(grpcServer, authService)

	return &App{
		log:        log,
		gRPCServer: grpcServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}

func logger(format string, a ...any) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}
