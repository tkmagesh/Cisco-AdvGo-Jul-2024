// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/service.proto

package proto

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

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	// Request & Response
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Server Streaming
	GeneratePrimes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (AppService_GeneratePrimesClient, error)
	// Client Streaming
	Aggregate(ctx context.Context, opts ...grpc.CallOption) (AppService_AggregateClient, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.AppService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GeneratePrimes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (AppService_GeneratePrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppService_ServiceDesc.Streams[0], "/proto.AppService/GeneratePrimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &appServiceGeneratePrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AppService_GeneratePrimesClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type appServiceGeneratePrimesClient struct {
	grpc.ClientStream
}

func (x *appServiceGeneratePrimesClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appServiceClient) Aggregate(ctx context.Context, opts ...grpc.CallOption) (AppService_AggregateClient, error) {
	stream, err := c.cc.NewStream(ctx, &AppService_ServiceDesc.Streams[1], "/proto.AppService/Aggregate", opts...)
	if err != nil {
		return nil, err
	}
	x := &appServiceAggregateClient{stream}
	return x, nil
}

type AppService_AggregateClient interface {
	Send(*AggregateRequest) error
	CloseAndRecv() (*AggregateResponse, error)
	grpc.ClientStream
}

type appServiceAggregateClient struct {
	grpc.ClientStream
}

func (x *appServiceAggregateClient) Send(m *AggregateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appServiceAggregateClient) CloseAndRecv() (*AggregateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AggregateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	// Request & Response
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Server Streaming
	GeneratePrimes(*PrimeRequest, AppService_GeneratePrimesServer) error
	// Client Streaming
	Aggregate(AppService_AggregateServer) error
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAppServiceServer) GeneratePrimes(*PrimeRequest, AppService_GeneratePrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GeneratePrimes not implemented")
}
func (UnimplementedAppServiceServer) Aggregate(AppService_AggregateServer) error {
	return status.Errorf(codes.Unimplemented, "method Aggregate not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AppService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GeneratePrimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServiceServer).GeneratePrimes(m, &appServiceGeneratePrimesServer{stream})
}

type AppService_GeneratePrimesServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type appServiceGeneratePrimesServer struct {
	grpc.ServerStream
}

func (x *appServiceGeneratePrimesServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AppService_Aggregate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServiceServer).Aggregate(&appServiceAggregateServer{stream})
}

type AppService_AggregateServer interface {
	SendAndClose(*AggregateResponse) error
	Recv() (*AggregateRequest, error)
	grpc.ServerStream
}

type appServiceAggregateServer struct {
	grpc.ServerStream
}

func (x *appServiceAggregateServer) SendAndClose(m *AggregateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appServiceAggregateServer) Recv() (*AggregateRequest, error) {
	m := new(AggregateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AppService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GeneratePrimes",
			Handler:       _AppService_GeneratePrimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Aggregate",
			Handler:       _AppService_Aggregate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}