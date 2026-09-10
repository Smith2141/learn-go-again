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

	for _, value := range input {
		if _, ok := storage[value]; !ok {
			// fmt.Println("is NOT present: ", value)
			// fmt.Println("input BEFORE: ", result)
			result = append(result, value)
			// fmt.Println("input AFTER: ", result)
            storage[value] = value
		}
	}

	return result
}
