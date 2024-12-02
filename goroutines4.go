package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func main() {
	cha := make(chan int)

	fmt.Println("Main fn started")

	wg.Add(1)
	go findSum(23, 45, cha)

	wg.Wait()
	fmt.Println("Main fn Ended")

}

func findSum(num1, num2 int, cha chan int) {
	result := num1 + num2

	cha <- result
	wg.Done()
}
