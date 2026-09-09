package main

import (
	"fmt"
	"strings"
)

func main() {
	var products = map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for p := range products {
		if products[p] > 500 {
			fmt.Println(p)
		}
	}

	// delimiter
	delimiter := strings.Repeat("-", 50)
	fmt.Println(delimiter)

	var order = []string{"хлеб", "буженина", "сыр", "огурцы", "чай"}

	var order_sum int = 0
	for _, item := range order {
		if price, ok := products[item]; ok {
			// добавить форматирование !!!
			fmt.Println(products[item], price)
			order_sum += price
		}
	}

	fmt.Println(delimiter)
	fmt.Println("Order total: ", order_sum)
}
