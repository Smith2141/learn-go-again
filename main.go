package main

import "fmt"

func main() {

	var a, b, d int = 1, 2, 4

	p_nums := [4]*int{&a, &b, nil, &d} // массив из 4 указателей на значения типа int

	for _, p := range p_nums {
		// если указатель не равен nil, выводим значение, которое хранится по его адресу
		if p != nil {
			fmt.Println(*p)
		}
	}
}
