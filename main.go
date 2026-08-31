package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	alfa       uint32 = 2013
	z          uint32 = 1997
	millennium uint32 = 1981
	x          uint32 = 1965
	boomer     uint32 = 1946
)

const errorValueMessage = "Значение года ВНЕ диапазона!"

func main() {
	year, err := getYearFromInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := getGeneration(year)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Привет, " + result)
}

func getYearFromInput() (uint32, error) {
	// Сначала проверяем переменную окружения
	if envYear := os.Getenv("TEST_YEAR"); envYear != "" {
		fmt.Printf("Используем год из окружения: %s\n", envYear)
		year, err := strconv.ParseUint(envYear, 10, 32)
		if err != nil {
			return 0, fmt.Errorf("неверный формат TEST_YEAR: %v", err)
		}
		return uint32(year), nil
	}

	// Если переменной нет - запрашиваем ввод
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите год: ")

	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		fmt.Println("Вы ввели:", input)

		year, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return 0, fmt.Errorf(errorValueMessage)
		}
		return uint32(year), nil
	}

	return 0, fmt.Errorf("не удалось прочитать ввод")
}

func getGeneration(year uint32) (string, error) {
	if boomer > year {
		return "", fmt.Errorf(errorValueMessage)
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
	case year >= boomer:
		result = "Бумер"
	default:
		result = "Неизвестное поколение"
	}
	return result, nil
}
