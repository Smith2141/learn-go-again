package main

import (
	"fmt"
)

var Global = 5

func useGlobal() {
	// 1. Ловим начальное значение Global
	defer func(checkout int) {
		// 3. Возвращаем Global исходное значение
		Global = checkout
	}(Global)

	// 2. Меняем значение
	Global = 42
	// 2.1 Демонстрируем измененное значение
	fmt.Println("change", Global)
}

func main() {
	// 0. Выводим начальное значение
	fmt.Println("first", Global)
	useGlobal()
	// 4. Демонстрируем возвращенное значение
	fmt.Println("final", Global)

}
