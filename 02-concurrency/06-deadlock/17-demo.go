package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go fn(wg)
	wg.Wait()
	fmt.Println("main completed")
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fn started")
	time.Sleep(3 * time.Second)
	fmt.Println("fn completed")
}
