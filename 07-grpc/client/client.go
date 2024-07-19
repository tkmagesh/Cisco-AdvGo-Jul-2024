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
func (clientApp *ClientApp) doClientStream(ctx context.Context) {
	nos := []int64{3, 1, 4, 2, 5, 9, 6, 8, 7}
	clientStream, err := clientApp.serviceClient.Aggregate(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		fmt.Println("Sending no :", no)
		req := &proto.AggregateRequest{
			No: no,
		}
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Client finished sending all the data")
	if res, err := clientStream.CloseAndRecv(); err == io.EOF || err == nil {
		fmt.Println("Sum :", res.GetSum())
		fmt.Println("Min :", res.GetMin())
		fmt.Println("Max :", res.GetMax())
	} else {
		log.Fatalln(err)
	}

}

func (ClientApp *ClientApp) doBidirectionalStream(ctx context.Context) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	clientStream, err := ClientApp.serviceClient.Greet(timeoutCtx)

	if err != nil {
		log.Fatalln(err)
	}
	go sendRequests(ctx, clientStream)
	done := make(chan struct{})
	go func() {
		fmt.Println("Press ENTER to cancel")
		fmt.Scanln()
		clientStream.CloseSend()
		close(done)
	}()
	go recvResponse(ctx, clientStream)
	// return done
	<-done
}

func sendRequests(ctx context.Context, clientStream proto.AppService_GreetClient) {
	persons := []*proto.PersonName{
		{FirstName: "Magesh", LastName: "Kuppan"},
		{FirstName: "Suresh", LastName: "Kannan"},
		{FirstName: "Ramesh", LastName: "Jayaraman"},
		{FirstName: "Rajesh", LastName: "Pandit"},
		{FirstName: "Ganesh", LastName: "Kumar"},
	}

	// done := make(chan struct{})

	for _, person := range persons {
		req := &proto.GreetRequest{
			Person: person,
		}
		log.Printf("Sending Person : %s %s\n", person.FirstName, person.LastName)
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func recvResponse(ctx context.Context, clientStream proto.AppService_GreetClient) {
	for {
		res, err := clientStream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.GetMessage())
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
	// clientApp.doServerStream(ctx)
	// clientApp.doClientStream(ctx)
	clientApp.doBidirectionalStream(ctx)
}
