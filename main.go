package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	alfa       uint32 = 2013
	z          uint32 = 1997
	millennium uint32 = 1981
	x          uint32 = 1965
	boomer     uint32 = 1997

	oldest  uint32 = 1946
	current uint32 = 2026
)

const errorValueMessage = "Значение года ВНЕ диапазона!"

func main() {
	input, error := getInput()

	if error != nil {
		fmt.Println(error)
		return
	}

	year64, error := strconv.ParseUint(input, 10, 8)
	if error != nil {
		// fmt.Println(error)
		fmt.Println(errorValueMessage)
		return
	}

	year := uint32(year64)

	result, error := getGeneration(year)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("Привет, " + result)
}

func getInput() (string, error) {
	// Create a scanner that wraps standard input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите год: ")

	// Scan() blocks until user hits Enter
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("You typed:", input)
		return input, nil
	}

	error := fmt.Errorf("Не удалось прочитать ввод")
	return "", error
}

func getGeneration(year uint32) (string, error) {
	if oldest > year || year > current {
		error := fmt.Errorf(errorValueMessage)
		return "", error
	}

	var result string

	switch {
	case year > alfa:
		result = "Альфа"
	case year > z:
		result = "Зумер"
	case year > millennium:
		result = "Миллениал"
	case year > x:
		result = "X"
	case year > boomer:
		result = "Бумер"
	}
	return result, nil
}
