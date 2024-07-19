package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tkmagesh/cisco-advgo-jul-2024/07-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientApp struct {
	serviceClient proto.AppServiceClient
}

func NewClientApp(clientConn *grpc.ClientConn) *ClientApp {
	return &ClientApp{
		serviceClient: proto.NewAppServiceClient(clientConn),
	}
}

func (clientApp *ClientApp) doRequestResponse(ctx context.Context) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResult, err := clientApp.serviceClient.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Sum : ", addResult.GetSum())
}

func (clientApp *ClientApp) doServerStream(ctx context.Context) {
	primeRequest := &proto.PrimeRequest{
		Start: 1000,
		End:   1100,
	}
	clientStream, err := clientApp.serviceClient.GeneratePrimes(ctx, primeRequest)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		// introduce a delay
		time.Sleep(500 * time.Millisecond)
		primeRes, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All prime numbers have been received...")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Prime No : %d\n", primeRes.GetPrimeNo())
	}
}
func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	clientApp := NewClientApp(clientConn)
	// clientApp.doRequestResponse(ctx)
	clientApp.doServerStream(ctx)
}
