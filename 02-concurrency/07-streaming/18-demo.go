package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func genNos(ch chan int) {
	ch <- 10
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- 20
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- 30
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- 40
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	ch <- 50
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
