package main

import "fmt"

func sum(s []int, cha chan int) {
	total := 0
	for _, r := range s {
		total += r
	}
	cha <- total
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}


	cha := make(chan int)

	go sum(arr[:len(arr)/2], cha)
	go sum(arr[len(arr)/2:], cha)

	sum1 := <-cha
	sum2 := <-cha

	result := sum1 + sum2
	fmt.Println(result)
}
