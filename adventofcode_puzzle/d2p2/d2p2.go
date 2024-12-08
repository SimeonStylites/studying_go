package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func issafe(list []int) int {
	d1 := list[0] - list[1]
	if d1 > 0 {
		for i := 0; i < len(list)-1; i++ {
			d2 := list[i] - list[i+1]
			if d2 > 3 || d2 < 1 {
				return 0
			}
		}
	} else {
		for i := 0; i < len(list)-1; i++ {
			d2 := list[i] - list[i+1]
			if d2 < -3 || d2 > -1 {
				return 0
			}
		}
	}
	return 1
}

func issafe2(list []int) int {
	if issafe(list) == 1 {
		return 1
	} else {
		for i, _ := range list {
			list_mod := make([]int, 0)
			list_mod = append(list_mod, list[:i]...)
			list_mod = append(list_mod, list[i+1:]...)
			fmt.Println(list_mod)
			if issafe(list_mod) == 1 {
				return 1
			}
		}
	}
	return 0
}

func toint(a []string) []int {
	b := make([]int, len(a), len(a))
	for i, s := range a {
		b[i], _ = strconv.Atoi(s)
	}
	return b
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\r\n")
	s := 0
	for _, line := range lines {
		numbers_str := strings.Split(line, " ")
		numbers := toint(numbers_str)
		s += issafe2(numbers)
	}
	fmt.Println(s)
	sequence1 := []int{1, 3, 2, 4, 5}
	fmt.Println(issafe2(sequence1))
}
