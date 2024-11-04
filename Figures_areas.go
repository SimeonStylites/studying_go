package main

import (
	"fmt"
	"math"
)

type figures int

//константы с типизацией (все типа figures)
const (
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
)
const pi = 3.14159

func area(fig figures) (func(a float64) float64, bool) {
	switch {
		case fig == square:
			return func(a float64) float64 {
				return a*a
			}, true
		case fig == circle:
			return func(a float64) float64 {
				return pi*a*a
			}, true
		case fig == triangle:
			return func(a float64) float64 {
				return a*a*math.Sqrt(3)/4.0
			}, true
		default:
			return nil, false
	}
}

func main() {
	a := 2.0
	myFigure := square
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Неизвестная фигура")
	} else {
		fmt.Println(ar(a))
	}
	
	ar, ok = area(1)
	if !ok {
		fmt.Println("Неизвестная фигура")
	} else {
		fmt.Println(ar(a))
	}
	
	ar, ok = area(triangle)
	if !ok {
		fmt.Println("Неизвестная фигура")
	} else {
		fmt.Println(ar(a))
	}
	
	ar, ok = area(4)
	if !ok {
		fmt.Println("Неизвестная фигура")
	} else {
		fmt.Println(ar(a))
	}
}