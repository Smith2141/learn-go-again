package main

import (
	"fmt"
)

func EvaluationOrder() {
	fmt.Println("Hello")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("World")
}

func main() {
	EvaluationOrder()
}
