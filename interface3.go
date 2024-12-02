package main

import (
	"fmt"
)

func sendNums(channel chan int) {
	for i := 1; i <= 5; i++ {
		channel <- i

	}
	close(channel)
}

// func recieveNums(channel chan int) {
// 	for nums := range channel {
// 		fmt.Println(nums)
// 	}
// }

func main() {
	channel := make(chan int)
	go sendNums(channel)
	for nums := range channel {
		fmt.Println(nums)
	}
}
