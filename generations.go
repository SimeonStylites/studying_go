package main

import "fmt"

func main() {
	var s int
	fmt.Println("Какой твой год рождения?")
	fmt.Scan(&s)
	var w string
	switch {
	case s>=1946 && s<=1964:
		w = ", бумер"
	case s>=1965 && s<=1980:
		w = ", представитель X"
	case s>=1981 && s<=1996:
		w = ", миллениал"
	case s>=1997 && s<=2012:
		w = ", зумер"
	case s>=2013:
		w = ", альфа"
	}
	fmt.Println("Привет"+w+"!")
}