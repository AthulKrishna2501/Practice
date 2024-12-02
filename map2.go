package main

import "fmt"

func findMap(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		result[key] = value
	}
	for key, value := range map2 {
		result[key] += value
	}
	return result
}

func main() {
	map1 := map[string]int{"a": 5, "b": 10, "c": 15}
	map2 := map[string]int{"a": 10, "d": 2, "f": 25}
	fmt.Println(findMap(map1, map2))

}
