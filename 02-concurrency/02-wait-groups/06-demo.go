package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for idx := range 100 {
		wg.Add(1)      // increment the counter by 1
		go fn(idx, wg) // schedule the execution of f1 through the scheduler
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
