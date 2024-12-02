package main

import "fmt"

func getResult(arr []int, channel chan int) {
	sum := 0
	for _, s := range arr {
		sum += s
	}
	channel <- sum
}

func main() {
	input1 := []int{1, 2, 3, 4, 54}
	input2 := []int{6, 7, 8, 9}

	channnel1 := make(chan int)
	channnel2 := make(chan int)

	go getResult(input1, channnel1)
	go getResult(input2, channnel2)

	result := <-channnel1
	result2 := <-channnel2
	if result > result2 {
		fmt.Println("result is greater", result)
	} else {
		fmt.Println("result 2 is greater", result)
	}

}
