package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}
	flag.IntVar(&count, "count", 0, "# of goroutines to spin up")
	flag.Parse()
	fmt.Printf("spinning up %d goroutines... hit ENTER to start!\n", count)
	fmt.Scanln()
	for idx := range count {
		wg.Add(1)        // increment the counter by 1
		go fn(idx+1, wg) // schedule the execution of f1 through the scheduler
	}
	wg.Wait() // block this function's execution until the counter becomes 0 (default)
	fmt.Println("Thank you!")

}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
