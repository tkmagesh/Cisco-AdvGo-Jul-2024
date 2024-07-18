/*
channel is used not only to communicate data, but also to synchronize goroutines
*/
package main

import (
	"fmt"
)

// consumer
func main() {
	ch := add(100, 200)
	result := <-ch
	fmt.Println("Add Result :", result)
}

// producer
func add(x, y int) chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result
	}()
	return ch
}
