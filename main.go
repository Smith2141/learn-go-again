package main

import "fmt"

func main() {

	var a int = 9

	var p_a *int = &a // p_a - указатель на переменную a

	// обращаемся к указателю p_a
	fmt.Println("Value of a:", *p_a)  // Value of a: 9
	fmt.Println("Address of a:", p_a) // Address of a: 0xc000010100

	var p_p_a **int = &p_a // p_p_a - указатель на переменную p_a

	// обращаемся к указателю p_a
	fmt.Println("Value of p_a:", *p_p_a)  // Value of p_a: 0xc000010100
	fmt.Println("Address of p_a:", p_p_a) // Address of p_a: 0xc000076040
}
