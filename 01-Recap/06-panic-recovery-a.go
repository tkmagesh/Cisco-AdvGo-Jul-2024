/*
panic = state of the application where the application execution is unable to proceed further
*/

package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic occurred, err = %q\n", e)
			return
		}
		fmt.Println("Thank You!")
	}()
	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[divide] calculating quotient")
	quotient = x / y
	fmt.Println("[divide] calculating remainder")
	remainder = x % y
	return
}
