package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Result struct {
	evenTotal int
	oddTotal  int
}

var ErrFileRead = errors.New("Error reading file")

func main() {
	result1 := processFile("data1.dat")
	result2 := processFile("data2.dat")
	evenTotal := result1.evenTotal + result2.evenTotal
	oddTotal := result1.oddTotal + result2.oddTotal
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "Even Total : %d\n", evenTotal)
	fmt.Fprintf(file, "Odd Total : %d\n", oddTotal)
}

func processFile(fileName string) Result {
	var result Result
	file, err := os.Open(fileName)
	if err != nil {
		panic(ErrFileRead)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if val, err := strconv.Atoi(txt); err == nil {
			if val%2 == 0 {
				result.evenTotal += val
			} else {
				result.oddTotal += val
			}
		}
	}
	return result

}
