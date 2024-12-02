package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	for i := 0;i<j;i++ {
		for j:=len(arr)-1{
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}
	fmt.Println(arr)

}
