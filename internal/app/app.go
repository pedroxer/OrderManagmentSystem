package app

import (
	grpcapp "github.com/pedroxer/ordermanagmentsystem/internal/app/grpc"
	authService "github.com/pedroxer/ordermanagmentsystem/internal/service/auth"
	"github.com/pedroxer/ordermanagmentsystem/internal/storage"
	log "github.com/sirupsen/logrus"
	"time"
)

type App struct {
	GRPCSrvr *grpcapp.App
}

func NewApp(log *log.Logger, grpcPort int, store storage.DBTX, tokenTTL time.Duration) *App {
	// init storage
	queries := storage.New(store)
	//init service
	authServ := authService.New(log, queries, queries, tokenTTL)
	grpcApp := grpcapp.NewApp(log, grpcPort, authServ)
	return &App{
		GRPCSrvr: grpcApp,
	}
}
