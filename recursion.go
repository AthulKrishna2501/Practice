package main

import "fmt"

func reversestring(str string) string {
	if len(str) == 0 {
		return ""
	}
	return string(str[len(str)-1]) + reversestring(str[:len(str)-1])
}

func main() {
	original := "This is a string"
	reversed := reversestring(original)

	fmt.Println(original)
	fmt.Println(reversed)

}
