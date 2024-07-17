package main

import "fmt"

type SimpleFunc func( /* no arguments */ ) /* no return values */
type OperationFunc func(int, int)

func main() {

	// functions as values
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/
	// var fn func()
	var fn SimpleFunc

	fn = func() {
		fmt.Println("fn[1] invoked")
	}
	fn()

	fn = func() {
		fmt.Println("fn[2] invoked")
	}
	fn()

	var operation OperationFunc
	operation = func(i1, i2 int) {
		fmt.Println(i1 + i2)
	}
	operation(100, 200)

	operation = func(i1, i2 int) {
		fmt.Println(i1 * i2)
	}
	operation(100, 200)
}
