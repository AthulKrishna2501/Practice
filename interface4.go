package main

import "fmt"

func total(arr []int, cha chan int) {
	result := 0
	for _, n := range arr {
		result += n
	}
	cha <- result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	cha := make(chan int)

	go total(arr[:len(arr)/2], cha)
	go total(arr[len(arr)/2:], cha)

	sum1 := <-cha
	sum2 := <-cha

	result := sum1 + sum2
	fmt.Println(result)
}
r