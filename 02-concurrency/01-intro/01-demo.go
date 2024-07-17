package main

import (
	"fmt"
)

func main() {
	go f1() // schedule the execution of f1 through the scheduler
	f2()

	// primitive synchronization (DO NOT DO THIS!!)

	// blocking the execution of main() function so that the scheduler can look for other goroutines scheduled and execute them

	// time.Sleep(500 * time.Millisecond)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
