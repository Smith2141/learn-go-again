package main

import (
	"fmt"
)

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println("filtered result: ", RemoveDuplicates(input))
}

func RemoveDuplicates(input []string) []string {
	var storage = map[string]string{}
	var result []string

	for i, value := range input {
		if _, ok := storage[value]; !ok {
			// fmt.Println("is NOT present: ", value)
			storage[value] = value
			// fmt.Println("input BEFORE: ", result)
			result = append(result, input[i])
			// fmt.Println("input AFTER: ", result)
		}
	}

	return result
}
