package main

import (
	"fmt"

	"github.com/Smith2141/gdp"
)

// Главная функция модуля
func main() {
	if sum := gdp.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
