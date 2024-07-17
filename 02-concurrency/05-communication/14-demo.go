/*
execute the add() as a goroutine and print the result in main()
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	// communicate by sharing memory
	var result int
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(wg, 100, 200, &result)
	wg.Wait()
	fmt.Println("Add Result :", result)
}

func add(wg *sync.WaitGroup, x, y int, result *int) {
	defer wg.Done()
	*result = x + y
}
