package app

import (
	grpcapp "github.com/pedroxer/ordermanagmentsystem/internal/app/grpc"
	"github.com/pedroxer/ordermanagmentsystem/internal/storage"
	log "github.com/sirupsen/logrus"
	"time"
)

type App struct {
	GRPCSrvr *grpcapp.App
}

func NewApp(log *log.Logger, grpcPort int, storage storage.DBTX, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.NewApp(log, grpcPort)
	return &App{
		GRPCSrvr: grpcApp,
	}
}
