package app

import (
	grpcapp "github.com/goget-milk/sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: Init storage
	// TODO: Init auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
