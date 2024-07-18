package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go doSomething(cancelCtx, wg)
	wg.Wait()
}

func doSomething(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[doSomething] cancellation signal received")
			break LOOP
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Print(".")
		}
	}
}
