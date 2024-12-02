package main

import "fmt"

type address struct {
	Place string
}

type person struct {
	Name    string
	Age     int
	Gender  string
	address address
}

func main() {

	map1 := map[string]interface{}{"key1": "hello", "key2": 10, "key3": map[string]interface{}{"subkey1": 30, "subkey2": "range", "subkey3": []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "subkey4": person{Name: "shuhaib", Age: 19, Gender: "male", address: address{Place: "thrissur"}}, "subkey5": map[interface{}]interface{}{1: []interface{}{"hello", false, 1}, "subinnerkey1": 1000}}}
	map1["key3"].(map[string]interface{})["subkey3"] = append(map1["key3"].(map[string]interface{})["subkey3"].([]int), 11)
	fmt.Println(map1)

}
