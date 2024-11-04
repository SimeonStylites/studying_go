package main

import (
	"fmt"
	"math"
)

func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func main() {
	//var diff int = 3
	var limit int = 10e7
	for i := 1; powInt(i, 3) <= limit; i++ {
		fmt.Println(powInt(i,3))
	}
}