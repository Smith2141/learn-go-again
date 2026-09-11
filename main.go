package main

import (
	"encoding/json"
	"fmt"
	"time"
	"log"
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
		DateOfBirth:time.Now(),
	}

	result, err := json.Marshal(man)

	if err != nil {
        log.Fatalln("unable marshal to json")
	}

	fmt.Printf("Man %v\n",string(result))
}
