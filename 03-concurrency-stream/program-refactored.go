package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type ResultStats struct {
	total int
	count int
}

func main() {
	fileWg := &sync.WaitGroup{}
	dataCh := make(chan int)
	fileWg.Add(1)
	go Source("data1.dat", dataCh, fileWg)
	fileWg.Add(1)
	go Source("data2.dat", dataCh, fileWg)

	done := ProcessData(dataCh)

	fileWg.Wait()
	close(dataCh)
	<-done

}

func ProcessData(dataCh chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		processWg := &sync.WaitGroup{}
		processWg.Add(1)
		evenCh, oddCh := Splitter(dataCh, processWg)

		processWg.Add(1)
		evenSumCh := Sum(evenCh, processWg)

		processWg.Add(1)
		oddSumCh := Sum(oddCh, processWg)

		processWg.Add(1)
		go Merger("result.txt", evenSumCh, oddSumCh, processWg)
		processWg.Wait()
		// close(doneCh) // to unblock the channel receive operation in other part of the application
		doneCh <- struct{}{} // to unblock the channel receive operation in other part of the application
	}()
	return doneCh
}

func Source(fileName string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		if no, err := strconv.Atoi(val); err == nil {
			ch <- no
		}
	}
}

func Splitter(ch chan int, wg *sync.WaitGroup) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go func() {
		defer wg.Done()
		defer close(evenCh)
		defer close(oddCh)
		counter := 0
		for data := range ch {
			counter += 1
			if data%2 == 0 {
				evenCh <- data
			} else {
				oddCh <- data
			}
		}
		fmt.Println("counter :", counter)
	}()
	return evenCh, oddCh
}

func Sum(ch <-chan int, wg *sync.WaitGroup) <-chan ResultStats {
	sumCh := make(chan ResultStats)
	go func() {
		defer wg.Done()
		total := 0
		count := 0
		for val := range ch {
			total += val
			count += 1
		}
		result := ResultStats{
			total: total,
			count: count,
		}
		sumCh <- result
	}()
	return sumCh

}

func Merger(fileName string, evenSumCh, oddSumCh <-chan ResultStats, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for i := 0; i < 2; i++ {
		select {
		case evenResult := <-evenSumCh:
			fmt.Fprintf(file, "Even Total : %d, Even Count : %d\n", evenResult.total, evenResult.count)
		case oddResult := <-oddSumCh:
			fmt.Fprintf(file, "Odd Total : %d, Odd Count : %d\n", oddResult.total, oddResult.count)
		}
	}

}
