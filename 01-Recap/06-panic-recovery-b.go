/*
connverting a panic into an error to allow course correction
*/

package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic occurred, err = %q\n", e)
			return
		}
		fmt.Println("Thank You!")
	}()
	var divisor int
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		// q, r := divide(100, divisor)
		q, r, err := divideClient(100, divisor)
		if err == ErrDivideByZero {
			fmt.Println("divisor cannot be 0, try again!")
			continue
		}
		if err != nil {
			fmt.Println("unknown error occurred, err =", err)
			continue
		}
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
		break
	}
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		// handle the panics to convert the panic into an error
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (cannot change)
func divide(x, y int) (quotient, remainder int) {

	fmt.Println("[divide] calculating quotient")
	// raise a panic programmatically
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("[divide] calculating remainder")
	remainder = x % y
	return
}
