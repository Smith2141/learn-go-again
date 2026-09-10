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

	fmt.Println("total: ", RemoveDuplicates(input))
}

func RemoveDuplicates(input []string) []string {
	var storage = map[string]string{}
	var result []string

	for i, value := range input {
		// fmt.Println(value)
		if _, ok := storage[value]; ok {
			fmt.Println("is present: ", value)

			// result = input[:i]
		} else {
			fmt.Println("is NOT present: ", value)
			storage[value] = value
			fmt.Println("input BEFORE: ", result)
			result = append(result, input[i])
			fmt.Println("input AFTER: ", result)
		}
	}

	return result
}
