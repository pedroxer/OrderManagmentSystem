package user

import (
	"context"
	proto_gen "github.com/pedroxer/ordermanagmentsystem/internal/proto_gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email, password string) (token string, err error)
	RegisterUser(ctx context.Context, email, password string) (userID int64, err error)
}

type userServerAPI struct {
	proto_gen.UnimplementedUserServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	proto_gen.RegisterUserServer(gRPC, &userServerAPI{auth: auth})
}

func (s *userServerAPI) RegisterUser(ctx context.Context, req *proto_gen.CreateUser) (*proto_gen.UserResp, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	// impl auth service
	userID, err := s.auth.RegisterUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &proto_gen.UserResp{UserID: int64(userID)}, nil
}

func (s *userServerAPI) Login(ctx context.Context, req *proto_gen.LoginRequest) (*proto_gen.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	// impl auth service
	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto_gen.LoginResponse{Token: token}, nil

}

func (s *userServerAPI) Logout(ctx context.Context, req *proto_gen.LogoutRequest) (*proto_gen.LogoutResponse, error) {
	return nil, nil
}
