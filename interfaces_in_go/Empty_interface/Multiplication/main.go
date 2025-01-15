package main

import "fmt"

func Mul(a interface{}, b int) interface{} {
	switch a1 := a.(type) {
	case int:
		return a1 * b
	case string:
		c := ""
		for i := 0; i < b; i++ {
			c = c + a1
		}
		return c
	default:
		return false
	}
}

func main() {
	a := "la"
	b := 37
	c := 3.3
	n := 3
	fmt.Println(Mul(a, n))
	fmt.Println(Mul(b, n))
	fmt.Println(Mul(c, n))
}
