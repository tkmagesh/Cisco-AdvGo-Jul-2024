/*
channel is used not only to communicate data, but also to synchronize goroutines
*/
package main

import (
	"fmt"
)

// consumer
func main() {
	ch := make(chan int)

	// converting a "sync" into an "async" without changing the code
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	fmt.Println("Add Result :", <-ch)
}

// 3rd party api (cannot change the code)
func add(x, y int) int {
	return x + y
}
