package grpcapp

import (
	"fmt"
	goodsGrpc "github.com/pedroxer/ordermanagmentsystem/internal/grpc/goods"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	logger     *log.Logger
	grpcServer *grpc.Server
	port       int
}

func NewApp(log *log.Logger, port int) *App {
	server := grpc.NewServer()

	goodsGrpc.Register(server)
	return &App{
		logger:     log,
		grpcServer: server,
		port:       port,
	}
}

func (app *App) Run() error {
	const op = "grpcapp.Run"
	log := app.logger.WithField("op", op)
	log.Info("starting grpc server on port ", app.port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		log.Fatalf("%s: %s", op, err)
	}

	log.Info("grpc server is running")
	return app.grpcServer.Serve(l)
}

func (app *App) Stop() {
	const op = "grpcapp.Stop"
	log.WithField("op", op).Info("stopping gRPC server")

	app.grpcServer.GracefulStop()
}
