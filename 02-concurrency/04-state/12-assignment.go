/*
Assignment:
Refactor the program to find the prime numbers concurrently
*/
package main

import "fmt"

func main() {
	var primes []int
	start, end := 1000, 2000
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	fmt.Println(primes)
}
func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
