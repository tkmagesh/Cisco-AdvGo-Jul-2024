package main

import (
	"fmt"
	"sync"
)

/* Encapsulate the concurrent safe manipulation of "count" logic in a custom type */
type Counter struct {
	sync.Mutex
	count int
}

func (counter *Counter) Increment() {
	counter.Lock()
	{
		counter.count++
	}
	counter.Unlock()
}

func (counter *Counter) Count() int {
	return counter.count
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}

	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.Count())
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
