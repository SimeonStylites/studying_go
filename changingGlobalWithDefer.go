package main

import "fmt"

var Global = 5

func ChangeGlobal() {
	defer func(a int) { //В аргументе запоминается начальное значение Global
						//После выполнения ChangeGlobal значение Global восстанавливается
		Global = a
	}(Global)
	Global++
	fmt.Println(Global)
}

func main() {
	fmt.Println(Global)
	ChangeGlobal()
	fmt.Println(Global)
}