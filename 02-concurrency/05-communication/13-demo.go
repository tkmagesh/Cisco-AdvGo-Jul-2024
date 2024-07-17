/* 
execute the add() as a goroutine and print the result in main() 
*/
package main

import "fmt"

func main() {
	result := go add(100, 200)
	fmt.Println("Add Result :", result)
}

func add(x, y int) int {
	return x + y
}
