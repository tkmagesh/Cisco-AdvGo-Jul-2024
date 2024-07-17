/*
execute the add() as a goroutine and print the result in main()
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	// share memory by communicating
	/*
		var ch chan int
		ch = make(chan int)
	*/
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200, ch)
	wg.Wait()
	result := <-ch
	fmt.Println("Add Result :", result)
}

func add(wg *sync.WaitGroup, x, y int, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
