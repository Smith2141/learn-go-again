package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string `json:"Имя"`
	Email       string `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {

	man := Person{
		Name:  "Alex",
		Email: "alex@yandex.ru",
	}

	result, err := json.Marshal(man)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(result))
}
