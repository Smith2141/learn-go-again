package main

import (
	"fmt"
	"math"
)

type figure int

const pi float64 = 3.1415

const (
	square   figure = iota // квадрат
	circle                 // круг
	triangle               // равносторонний треугольник
	unknown
)

func area(f figure) (func(float64) float64, bool) {

	switch f {
	case square:
		return func(f float64) float64 { return f * f }, true
	case circle:
		return func(f float64) float64 { return pi * f * f }, true
	case triangle:
		return func(f float64) float64 { return math.Sqrt(3) / 4 * f * f }, true
	default:
		return nil, false
	}
}

func main() {
	// var myFigure figure = square
	// var myFigure figure = triangle
	// var myFigure figure = circle
	var myFigure figure = unknown

	ar, ok := area(myFigure)
	x := 10.0

	if !ok {
		fmt.Println("Ошибка, фигура не известна")
		return
	}
	myArea := ar(x)

	fmt.Println(myArea)
}
