package main

import (
	"fmt"
)

type Person struct {
	Name     string // Имя
	NumChild int    // Количество детей
	Age      int    // Возраст
}

func main() {

	man := Person{
		Name:     "Alex",
		Age:      30,
		NumChild: 2,
	}

	fmt.Printf("Man %#v\n", man)
	additional(man)
}

func additional(p Person) {
	fmt.Printf("Man %#v\n", p)

}
