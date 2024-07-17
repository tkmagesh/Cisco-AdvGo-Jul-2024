package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "Anim ullamco sint do sit aute tempor id minim."
	x = 99.99
	// x = struct{}{}
	/* x = func() {
		fmt.Println("fn invoked")
	} */
	fmt.Println(x)

	// x = 100
	x = "Culpa in est nisi labore velit reprehenderit aliquip et quis."
	// y := x * 2
	// y := x.(int) * 2

	// type assertion (if)
	if val, ok := x.(int); ok {
		y := val * 2
		fmt.Println(y)
	} else {
		fmt.Println("x has a non int value")
	}

	// type assertion (type switch)
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x * 2 =", val*2)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case func():
		fmt.Println("x is a func")
		val()
	default:
		fmt.Println("x is an unknown type")
	}

}
