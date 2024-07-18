package main

import "fmt"

func main() {
	// create a channel
	ch := make(chan int)

	// receive & print data from the channel
	data := <-ch
	fmt.Println(data)

	// send the data to the channel
	ch <- 100

}
