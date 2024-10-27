package goods

import (
	"context"
	proto_gen "github.com/pedroxer/ordermanagmentsystem/internal/proto_gen/proto"
	"google.golang.org/grpc"
)

type goodsServerAPI struct {
	proto_gen.UnimplementedGoodsServer
}

func Register(gRPC *grpc.Server) {
	proto_gen.RegisterGoodsServer(gRPC, &goodsServerAPI{})
}

func (s *goodsServerAPI) AddGoods(serv proto_gen.Goods_AddGoodsServer) error {
	const op = "goods.AddGoods"
	var req proto_gen.AddGoodsRequest
	err := serv.RecvMsg(&req)
	if err != nil {
		return err
	}
	return nil
}
func (s *goodsServerAPI) GetGood(ctx context.Context, req *proto_gen.GetGoodRequest) (*proto_gen.GetGoodResponse, error) {
	return nil, nil
}
