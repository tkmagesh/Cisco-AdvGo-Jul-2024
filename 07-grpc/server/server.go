package main

import (
	"context"
	"fmt"
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

func (asi *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Printf("[AppService - GeneratePrimes] start = %d and end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			log.Printf("[AppService - GeneratePrimes] sending prime no : %d\n", no)
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			if err := serverStream.Send(res); err != nil {
				log.Fatalln(err)
			}
		}
	}
	fmt.Println("[AppService - GeneratePrimes] Done!")
	return nil // io.EOF will be returned
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
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
