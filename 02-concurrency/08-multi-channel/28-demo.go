/*
keep generating prime numbers until a stop signal is received
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// timeoutCh := timeout(5 * time.Second)
	timeoutCh := time.After(5 * time.Second)
	ch := genPrimes(1000, timeoutCh)
	for primeNo := range ch {
		fmt.Println("primeNo :", primeNo)
	}
	fmt.Println("Done!")
	elapsed := time.Since(start)
	fmt.Println("time taken :", elapsed)
}

/*
func timeout(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		ch <- time.Now()
	}()
	return ch
}
*/

func genPrimes(start int, stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for no := start; ; no++ {
			time.Sleep(100 * time.Millisecond)
			select {
			case <-stopCh:
				break LOOP
			default:
				if isPrime(no) {
					ch <- no
				}
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
