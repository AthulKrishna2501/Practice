package main

import "fmt"

func getSum(s []int, channel chan int) {
	sum := 0
	for _, n := range s {
		sum += n
	}
	channel <- sum
}

func main() {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []int{6, 7, 8, 35, 3}

	cha1 := make(chan int)
	cha2 := make(chan int)

	go getSum(input1, cha1)
	go getSum(input2, cha2)

	result1 := <-cha1
	result2 := <-cha2

	if result1 > result1 {
		fmt.Println("Result 1 is Greater", result1)
	} else {
		fmt.Println("Result 2 is Greater", result2)
	}

}
