package main

import "fmt"

func main() {
	countries := make(map[string]map[string]int)

	countries["USA"] = map[string]int{
		"NewYork":     1479359,
		"Los Angeles": 3579320,
	}

	countries["India"] = map[string]int{
		"Mumbai": 3457293,
		"Delhi":  23510825,
	}

	countries["USA"]["Chicago"] = 23509235
	fmt.Println(countries)

	if population, exists := countries["India"]["Mumbai"]; exists {
		fmt.Println("Country %s exists", population)
	} else {
		fmt.Println("The country  not exitsts", population)
	}
	delete(countries["USA"], "NewYork")
}
