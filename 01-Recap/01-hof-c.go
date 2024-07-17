package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed")

		log.Println("Operation started")
		subtract(100, 200)
		log.Println("Operation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(i1, i2 int) {
		fmt.Println("Multiply Result :", i1*i2)
	}, 100, 200)
}

// ver 2.0
func logOperation(operationFn func(int, int), x, y int) {
	log.Println("Operation started")
	operationFn(x, y)
	log.Println("Operation completed")
}

/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

// ver 1.0 (frozen)
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
