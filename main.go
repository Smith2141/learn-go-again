package main

import "fmt"

const (
	Black = iota
	Gray
	White
)

// счётчик обнуляется
const (
	Yellow = iota
	Red
	Green = iota // это присваивание не обнулит iota
	Blue
)

func main() {
	fmt.Println(Black, Gray, White)
	fmt.Println(Yellow, Red, Green, Blue)
}
