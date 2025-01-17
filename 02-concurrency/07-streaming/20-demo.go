package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		fmt.Println("Channel closed... exiting...!")
		break
	}
	fmt.Println("Done!")
}

// producer
func genNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Println("count :", count)
	for i := range count {
		ch <- (i + 1) * 10
	}
	close(ch)
}
