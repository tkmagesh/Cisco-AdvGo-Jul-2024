package main

import (
	"context"
	"log"
	"net"

	"github.com/tkmagesh/cisco-advgo-jul-2024/07-grpc/proto"
	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer //inheriting from UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	// extract the data from the request payload
	x := req.GetX()
	y := req.GetY()
	log.Printf("[AppService - Add] x = %d and y = %d\n", x, y)

	// process the data
	result := x + y

	// return the response payload
	resp := &proto.AddResponse{
		Sum: result,
	}
	return resp, nil
}

func main() {
	// create an instance of the service implementation
	asi := &AppServiceImpl{}

	// create a tcp listener
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	// create a grpc server instance
	grpcServer := grpc.NewServer()

	// register the service implementation with the grpc server
	proto.RegisterAppServiceServer(grpcServer, asi)

	// listen for the requests
	grpcServer.Serve(listener)
}
