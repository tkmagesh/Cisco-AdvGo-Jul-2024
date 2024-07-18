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
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER key to stop...!")
	go func() {
		fmt.Scanln()
		close(stopCh)
	}()
	ch := genPrimes(1000, stopCh)
	for primeNo := range ch {
		fmt.Println("primeNo :", primeNo)
	}
	fmt.Println("Done!")
	elapsed := time.Since(start)
	fmt.Println("time taken :", elapsed)
}

func genPrimes(start int, stopCh chan struct{}) chan int {
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
