package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string // Имя
	Email       string
	DateOfBirth time.Time
}

func main() {

	man := Person{
		Name:  "Alex",
		Email: "alex@yandex.ru",
	}

	result, err := json.Marshal(man)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}
