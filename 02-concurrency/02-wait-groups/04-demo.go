package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the counter by 1
	go f1()   // schedule the execution of f1 through the scheduler
	f2()

	// blocking the execution of main() function so that the scheduler can look for other goroutines scheduled and execute them

	wg.Wait() // block this function's execution until the counter becomes 0 (default)

}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
