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
	result := <-ch
	wg.Wait()
	fmt.Println("Add Result :", result)
}

func add(wg *sync.WaitGroup, x, y int, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
