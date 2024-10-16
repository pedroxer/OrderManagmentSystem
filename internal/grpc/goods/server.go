package goods

import (
	proto_gen "github.com/pedroxer/ordermanagmentsystem/internal/proto_gen/proto"
	"google.golang.org/grpc"
)

type serverAPI struct {
	proto_gen.UnimplementedGoodsServer
}

func Register(gRPC *grpc.Server) {
	proto_gen.RegisterGoodsServer(gRPC, &serverAPI{})
}
