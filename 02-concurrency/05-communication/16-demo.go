package main

import (
	"fmt"
)

func main() {
	// create a channel
	ch := make(chan int)

	/*
		// "Send" operation in a goroutine
		go func() {
			// send the data to the channel
			ch <- 100
		}()

		// receive & print data from the channel
		data := <-ch
		fmt.Println(data)
	*/

	/*
		// "Recieve" operation in a goroutine
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			// receive & print data from the channel
			data := <-ch // (nb)
			fmt.Println(data)
			wg.Done()
		}()

		// send the data to the channel
		ch <- 100 // (b)
		wg.Wait()
	*/

	// "Recieve" operation in a goroutine (using a channel instead of a waitgroup)
	doneCh := make(chan struct{})
	go func() {
		// receive & print data from the channel
		data := <-ch // (nb)
		fmt.Println(data)
		doneCh <- struct{}{}
	}()

	// send the data to the channel
	ch <- 100 // (b)
	<-doneCh
}
