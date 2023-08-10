package main

import (
	"math"
	"fmt"
)

type figures int

const (
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
) 

func area(figure figures) (func(float64) float64, bool) {
	switch figure {
	case square:
		return func(x float64) float64 {
			return x*x
		}, true
	case circle:
		return func(x float64) float64 {
			return math.Pi*x*x
		}, true
	case triangle:
		return func(x float64) float64 {
			return (math.Sqrt(3)/4)*x*x
		}, true
	default:
		return nil, false
	}

}

func main() {
	ar, ok := area(triangle)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(2)

	fmt.Println(myArea)
}