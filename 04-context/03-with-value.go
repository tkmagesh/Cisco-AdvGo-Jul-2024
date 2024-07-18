package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go doSomething(valCtx, wg)
	wg.Wait()
}

func doSomething(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("root-key =", ctx.Value("root-key"))
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}

}
