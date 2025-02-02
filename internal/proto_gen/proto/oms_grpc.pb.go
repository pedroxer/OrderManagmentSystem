// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: proto/oms.proto

package proto_gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*UserResp, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *CreateUser, opts ...grpc.CallOption) (*UserResp, error) {
	out := new(UserResp)
	err := c.cc.Invoke(ctx, "/OMS.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/OMS.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/OMS.User/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *CreateUser) (*UserResp, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *CreateUser) (*UserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*CreateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.User/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OMS.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _User_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/oms.proto",
}

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	GetAllOrdersByUser(ctx context.Context, in *GetAllOrdersByUserRequest, opts ...grpc.CallOption) (Order_GetAllOrdersByUserClient, error)
	ChangeOrderState(ctx context.Context, in *ChangeOrderStateRequest, opts ...grpc.CallOption) (*ChangeOrderStateResponse, error)
	AddGoodToOrder(ctx context.Context, in *AddGoodToOrderRequest, opts ...grpc.CallOption) (*AddGoodToOrderResponse, error)
	RemoveGoodFromOrder(ctx context.Context, in *RemoveGoodFromOrderRequest, opts ...grpc.CallOption) (*RemoveGoodFromOrderResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, "/OMS.Order/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetAllOrdersByUser(ctx context.Context, in *GetAllOrdersByUserRequest, opts ...grpc.CallOption) (Order_GetAllOrdersByUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Order_ServiceDesc.Streams[0], "/OMS.Order/GetAllOrdersByUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderGetAllOrdersByUserClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Order_GetAllOrdersByUserClient interface {
	Recv() (*GetAllOrdersByUserResponse, error)
	grpc.ClientStream
}

type orderGetAllOrdersByUserClient struct {
	grpc.ClientStream
}

func (x *orderGetAllOrdersByUserClient) Recv() (*GetAllOrdersByUserResponse, error) {
	m := new(GetAllOrdersByUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderClient) ChangeOrderState(ctx context.Context, in *ChangeOrderStateRequest, opts ...grpc.CallOption) (*ChangeOrderStateResponse, error) {
	out := new(ChangeOrderStateResponse)
	err := c.cc.Invoke(ctx, "/OMS.Order/ChangeOrderState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) AddGoodToOrder(ctx context.Context, in *AddGoodToOrderRequest, opts ...grpc.CallOption) (*AddGoodToOrderResponse, error) {
	out := new(AddGoodToOrderResponse)
	err := c.cc.Invoke(ctx, "/OMS.Order/AddGoodToOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) RemoveGoodFromOrder(ctx context.Context, in *RemoveGoodFromOrderRequest, opts ...grpc.CallOption) (*RemoveGoodFromOrderResponse, error) {
	out := new(RemoveGoodFromOrderResponse)
	err := c.cc.Invoke(ctx, "/OMS.Order/RemoveGoodFromOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	GetAllOrdersByUser(*GetAllOrdersByUserRequest, Order_GetAllOrdersByUserServer) error
	ChangeOrderState(context.Context, *ChangeOrderStateRequest) (*ChangeOrderStateResponse, error)
	AddGoodToOrder(context.Context, *AddGoodToOrderRequest) (*AddGoodToOrderResponse, error)
	RemoveGoodFromOrder(context.Context, *RemoveGoodFromOrderRequest) (*RemoveGoodFromOrderResponse, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedOrderServer) GetAllOrdersByUser(*GetAllOrdersByUserRequest, Order_GetAllOrdersByUserServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllOrdersByUser not implemented")
}
func (UnimplementedOrderServer) ChangeOrderState(context.Context, *ChangeOrderStateRequest) (*ChangeOrderStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOrderState not implemented")
}
func (UnimplementedOrderServer) AddGoodToOrder(context.Context, *AddGoodToOrderRequest) (*AddGoodToOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGoodToOrder not implemented")
}
func (UnimplementedOrderServer) RemoveGoodFromOrder(context.Context, *RemoveGoodFromOrderRequest) (*RemoveGoodFromOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGoodFromOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.Order/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetAllOrdersByUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllOrdersByUserRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServer).GetAllOrdersByUser(m, &orderGetAllOrdersByUserServer{stream})
}

type Order_GetAllOrdersByUserServer interface {
	Send(*GetAllOrdersByUserResponse) error
	grpc.ServerStream
}

type orderGetAllOrdersByUserServer struct {
	grpc.ServerStream
}

func (x *orderGetAllOrdersByUserServer) Send(m *GetAllOrdersByUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Order_ChangeOrderState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeOrderStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ChangeOrderState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.Order/ChangeOrderState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ChangeOrderState(ctx, req.(*ChangeOrderStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_AddGoodToOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodToOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AddGoodToOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.Order/AddGoodToOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AddGoodToOrder(ctx, req.(*AddGoodToOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_RemoveGoodFromOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGoodFromOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).RemoveGoodFromOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.Order/RemoveGoodFromOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).RemoveGoodFromOrder(ctx, req.(*RemoveGoodFromOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OMS.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _Order_PlaceOrder_Handler,
		},
		{
			MethodName: "ChangeOrderState",
			Handler:    _Order_ChangeOrderState_Handler,
		},
		{
			MethodName: "AddGoodToOrder",
			Handler:    _Order_AddGoodToOrder_Handler,
		},
		{
			MethodName: "RemoveGoodFromOrder",
			Handler:    _Order_RemoveGoodFromOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllOrdersByUser",
			Handler:       _Order_GetAllOrdersByUser_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/oms.proto",
}

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	AddGoods(ctx context.Context, opts ...grpc.CallOption) (Goods_AddGoodsClient, error)
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) AddGoods(ctx context.Context, opts ...grpc.CallOption) (Goods_AddGoodsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Goods_ServiceDesc.Streams[0], "/OMS.Goods/AddGoods", opts...)
	if err != nil {
		return nil, err
	}
	x := &goodsAddGoodsClient{stream}
	return x, nil
}

type Goods_AddGoodsClient interface {
	Send(*AddGoodsRequest) error
	CloseAndRecv() (*AddGoodsResponse, error)
	grpc.ClientStream
}

type goodsAddGoodsClient struct {
	grpc.ClientStream
}

func (x *goodsAddGoodsClient) Send(m *AddGoodsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *goodsAddGoodsClient) CloseAndRecv() (*AddGoodsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddGoodsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goodsClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := c.cc.Invoke(ctx, "/OMS.Goods/GetGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	AddGoods(Goods_AddGoodsServer) error
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) AddGoods(Goods_AddGoodsServer) error {
	return status.Errorf(codes.Unimplemented, "method AddGoods not implemented")
}
func (UnimplementedGoodsServer) GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGood not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_AddGoods_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GoodsServer).AddGoods(&goodsAddGoodsServer{stream})
}

type Goods_AddGoodsServer interface {
	SendAndClose(*AddGoodsResponse) error
	Recv() (*AddGoodsRequest, error)
	grpc.ServerStream
}

type goodsAddGoodsServer struct {
	grpc.ServerStream
}

func (x *goodsAddGoodsServer) SendAndClose(m *AddGoodsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *goodsAddGoodsServer) Recv() (*AddGoodsRequest, error) {
	m := new(AddGoodsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Goods_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.Goods/GetGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OMS.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGood",
			Handler:    _Goods_GetGood_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddGoods",
			Handler:       _Goods_AddGoods_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/oms.proto",
}

// PaymentsClient is the client API for Payments service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentsClient interface {
	ProvidePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
}

type paymentsClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsClient(cc grpc.ClientConnInterface) PaymentsClient {
	return &paymentsClient{cc}
}

func (c *paymentsClient) ProvidePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/OMS.Payments/ProvidePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsServer is the server API for Payments service.
// All implementations must embed UnimplementedPaymentsServer
// for forward compatibility
type PaymentsServer interface {
	ProvidePayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	mustEmbedUnimplementedPaymentsServer()
}

// UnimplementedPaymentsServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentsServer struct {
}

func (UnimplementedPaymentsServer) ProvidePayment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvidePayment not implemented")
}
func (UnimplementedPaymentsServer) mustEmbedUnimplementedPaymentsServer() {}

// UnsafePaymentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentsServer will
// result in compilation errors.
type UnsafePaymentsServer interface {
	mustEmbedUnimplementedPaymentsServer()
}

func RegisterPaymentsServer(s grpc.ServiceRegistrar, srv PaymentsServer) {
	s.RegisterService(&Payments_ServiceDesc, srv)
}

func _Payments_ProvidePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsServer).ProvidePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OMS.Payments/ProvidePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsServer).ProvidePayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Payments_ServiceDesc is the grpc.ServiceDesc for Payments service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payments_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OMS.Payments",
	HandlerType: (*PaymentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProvidePayment",
			Handler:    _Payments_ProvidePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/oms.proto",
}
