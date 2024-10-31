package auth_interceptor

import (
	"context"
	models "github.com/pedroxer/ordermanagmentsystem/internal/storage"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type TokenMaker interface {
	VerifyToken(token string) (int64, error)
	NewToken(user models.User) (string, error)
}
type AuthInterceptor struct {
	tokenMaker TokenMaker
}

func New(maker TokenMaker) *AuthInterceptor {
	return &AuthInterceptor{
		tokenMaker: maker,
	}
}
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		const op = "auth_interceptor.Unary"
		log := log.WithField("op", op)
		if info.FullMethod == "/OMS.User/Login" || info.FullMethod == "/OMS.User/RegisterUser" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "failed getting metadata from context")
		}
		authToken := md.Get("authorization")
		if len(authToken) == 0 {
			return nil, status.Error(codes.Unauthenticated, "no token provided")
		}
		uid, err := i.tokenMaker.VerifyToken(authToken[0][7:])
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "internal error")
		}
		if uid == -1 {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}
		log.Info("auth successful")
		return handler(ctx, req)
	}

}

func (i *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		if info.FullMethod == "/OMS.User/Login" || info.FullMethod == "/OMS.User/RegisterUser" {
			return handler(srv, stream)
		}
		md, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return status.Error(codes.Internal, "failed getting metadata from context")
		}
		authToken := md.Get("authorization")
		if len(authToken) == 0 {
			return status.Error(codes.Unauthenticated, "no token provided")
		}
		uid, err := i.tokenMaker.VerifyToken(authToken[0][7:])
		if err != nil {
			return status.Error(codes.Unauthenticated, "internal error")
		}
		if uid == -1 {
			return status.Error(codes.Unauthenticated, "invalid token")
		}
		log.Info("auth successful")

		return handler(srv, stream)
	}
}
