package app

import (
	grpcapp "github.com/pedroxer/ordermanagmentsystem/internal/app/grpc"
	"github.com/pedroxer/ordermanagmentsystem/internal/config"
	jwt_maker "github.com/pedroxer/ordermanagmentsystem/internal/lib/jwt"
	authService "github.com/pedroxer/ordermanagmentsystem/internal/service/auth"
	"github.com/pedroxer/ordermanagmentsystem/internal/storage"
	log "github.com/sirupsen/logrus"
	"time"
)

type App struct {
	GRPCSrvr *grpcapp.App
}

func NewApp(log *log.Logger, grpcPort int, store storage.DBTX, conf config.Config) *App {
	// init storage
	queries := storage.New(store)
	//init service
	tokenTTL, err := time.ParseDuration(conf.TokenTTL)
	if err != nil {
		log.Fatal(err)
	}
	jwtToken := jwt_maker.NewJwtMaker(conf.Secret, tokenTTL)
	authServ := authService.New(log, queries, queries, jwtToken)
	grpcApp := grpcapp.NewApp(log, grpcPort, authServ, jwtToken)
	return &App{
		GRPCSrvr: grpcApp,
	}
}
