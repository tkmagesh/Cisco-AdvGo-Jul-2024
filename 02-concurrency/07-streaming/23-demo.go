package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(no int) {
			fmt.Println(no)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
}
