/*
Assignment:
Refactor the program to find the prime numbers concurrently
*/
package main

import (
	"fmt"
	"sync"
)

var primes []int
var mutex sync.Mutex

func main() {

	start, end := 1000, 2000
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go processNo(wg, no)
	}
	wg.Wait()
	fmt.Println("primes count :", len(primes))
	fmt.Println(primes)
}

func processNo(wg *sync.WaitGroup, no int) {
	defer wg.Done()
	if isPrime(no) {
		mutex.Lock()
		{
			primes = append(primes, no)
		}
		mutex.Unlock()
	}
}
func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
