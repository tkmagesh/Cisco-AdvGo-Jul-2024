package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	main - deferred")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	f1[1] - deferred")
		}()
		defer func() {
			fmt.Println("	f1[2] - deferred")
		}()
		defer func() {
			fmt.Println("	f1[3] - deferred")
		}() */
	defer fmt.Println("	f1[1] - deferred")
	defer fmt.Println("	f1[2] - deferred")
	defer fmt.Println("	f1[3] - deferred")
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("	f2 - deferred")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
