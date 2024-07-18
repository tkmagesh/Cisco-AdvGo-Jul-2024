package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	rootValCtx := context.WithValue(rootCtx, "root-key", "root-value")
	cancelCtx, cancel := context.WithCancel(rootValCtx)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go printNos(cancelCtx, wg)
	wg.Wait()
}

func printNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printNos] root-key = ", ctx.Value("root-key"))

	// send more data through context
	parentValCtx := context.WithValue(ctx, "parent-key", "parent-value")
	// start sub-task 1
	wg.Add(1)
	evenTimeCtx, cancel := context.WithTimeout(parentValCtx, 5*time.Second)
	defer cancel()
	go printEvenNos(evenTimeCtx, wg)

	//start sub-task 2
	wg.Add(1)
	oddTimeCtx, cancel := context.WithTimeout(parentValCtx, 10*time.Second)
	defer cancel()
	go printOddNos(oddTimeCtx, wg)
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("[printNos] cancellation signal received!")
			break LOOP
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("[printNos] no :", i)
		}

	}
}

func printEvenNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printEvenNos] root-key = ", ctx.Value("root-key"))
	fmt.Println("[printEvenNos] parent-key = ", ctx.Value("parent-key"))
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("[printEvenNos] cancellation signal received!")
			break LOOP
		default:
			if i%2 == 0 {
				time.Sleep(300 * time.Millisecond)
				fmt.Println("[printEvenNos] no :", i)
			}
		}

	}
}

func printOddNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printOddNos] root-key = ", ctx.Value("root-key"))
	fmt.Println("[printOddNos] parent-key = ", ctx.Value("parent-key"))
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("[printOddNos] cancellation signal received!")
			break LOOP
		default:
			if i%2 != 0 {
				time.Sleep(500 * time.Millisecond)
				fmt.Println("[printOddNos] no :", i)
			}
		}
	}
}
