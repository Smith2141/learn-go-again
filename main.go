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

	for product, price := range products {
		if price > 500 {
			fmt.Println(product)
		}
	}

	// delimiter
	delimiter := strings.Repeat("-", 50)
	fmt.Println(delimiter)

	var order = []string{"хлеб", "буженина", "сыр", "огурцы"}

	var total int = 0
	for _, item := range order {
		if price, ok := products[item]; ok {
			// добавить форматирование !!!
			fmt.Printf("Товар: %s, стоимость %d\n", item, price)
			total += price
		}
	}

	fmt.Println(delimiter)
	fmt.Println("Стоимость заказа ", total)
}
