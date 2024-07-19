package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/tkmagesh/cisco-advgo-jul-2024/07-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (asi *AppServiceImpl) Aggregate(serverStream proto.AppService_AggregateServer) error {
	var sum, min, max int64 = 0, 9223372036854775807, -9223372036854775808
LOOP:
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			log.Println("[AppService - Aggregate] All the data have been received")
			res := &proto.AggregateResponse{
				Sum: sum,
				Min: min,
				Max: max,
			}
			if err := serverStream.SendAndClose(res); err != io.EOF && err != nil {
				log.Fatalln(err)
			}
			break LOOP
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(req)
		no := req.GetNo()
		sum += no
		if no < min {
			min = no
		}
		if no > max {
			max = no
		}
	}
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (asi *AppServiceImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		greetReq, err := serverStream.Recv()
		if code := status.Code(err); code == codes.Unavailable {
			fmt.Println("Client connection closed")
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		person := greetReq.GetPerson()
		firstName := person.GetFirstName()
		lastName := person.GetLastName()
		log.Printf("Received greet request for %q and %q\n", firstName, lastName)
		message := fmt.Sprintf("Hi %s %s, Have a nice day!", firstName, lastName)
		time.Sleep(2 * time.Second)
		log.Printf("Sending response : %q\n", message)
		greetResp := &proto.GreetResponse{
			Message: message,
		}
		if err := serverStream.Send(greetResp); err != nil {
			if code := status.Code(err); code == codes.Unavailable {
				fmt.Println("Client connection closed")
				break
			}
		}
	}
	return nil
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
