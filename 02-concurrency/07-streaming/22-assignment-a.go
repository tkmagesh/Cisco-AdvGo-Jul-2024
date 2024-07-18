/*
write a "genPrimes()" (producer) function that generates prime numbers within the given "start" and "end" asynchronously and returns the prime numbers one at a time as and when they are generated
print the generated prime numbers in the main function (consumer) as and when they are generated
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := genPrimes(1000, 2000)
	for primeNo := range ch {
		fmt.Println("primeNo :", primeNo)
	}
	fmt.Println("Done!")
	elapsed := time.Since(start)
	fmt.Println("time taken :", elapsed)
}

func genPrimes(start, end int) chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
