/*
channel is used not only to communicate data, but also to synchronize goroutines
*/
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("Add Result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
