package grpcapp

import (
	"fmt"
	"github.com/pedroxer/ordermanagmentsystem/internal/grpc/auth_interceptor"
	goodsGrpc "github.com/pedroxer/ordermanagmentsystem/internal/grpc/goods"
	userGrpc "github.com/pedroxer/ordermanagmentsystem/internal/grpc/user"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	logger     *log.Logger
	grpcServer *grpc.Server
	port       int
}

func NewApp(log *log.Logger, port int, auth userGrpc.Auth, maker auth_interceptor.TokenMaker) *App {
	interceptor := auth_interceptor.New(maker)
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()), grpc.StreamInterceptor(interceptor.Stream()))

	goodsGrpc.Register(server)
	userGrpc.Register(server, auth)
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
