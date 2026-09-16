package main

import (
	"fmt"
)

func EvaluationOrder() {
	defer fmt.Println("deferred 1")
	fmt.Println("evaluated 1")
	defer fmt.Println("deferred 2")
	fmt.Println("evaluated 2")
}

func main() {
	EvaluationOrder()
}
