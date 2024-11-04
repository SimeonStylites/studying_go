package main

import "fmt"

func main() {
	var a int = 5
	p := &a

	fmt.Println(a,p) //a=5 p=0xc0000b2008
	
	i := 42
	p = &i
	fmt.Println(*p) // читаем значение переменной i через указатель p
	*p = *p/2         // записываем в переменную i значение 21 через указатель p
	fmt.Println(i)
}