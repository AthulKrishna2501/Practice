package main

import "fmt"

func main() {

	nestedMap := map[string]interface{}{
		"key1": map[string]interface{}{
			"innerKey1": "value1", "innerKey2": 42, "innerKey3": true,
		},
		"key2": map[string]interface{}{
			"innerKey1": []int{1, 2, 3, 4},
			"innerKey2": map[string]string{
				"subInnerKey1": "subValue1", "subInnerKey2": "subValue2",
			},
		},
		"key3": "simpleValue",
	}

	nestedMap["key1"].(map[string]interface{})["innerKey3"] = false
	fmt.Println(nestedMap)
	nestedMap["key2"].(map[string]interface{})["innerKey1"] = append(nestedMap["key2"].(map[string]interface{})["innerKey1"].([]int), 5)
	fmt.Println(nestedMap)

}
