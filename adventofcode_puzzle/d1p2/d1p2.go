package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sim(a, b []int) int {
	res := 0
	for _, n := range a {
		for _, m := range b {
			if n == m {
				res += n
			}
		}
	}
	return res
}

func main() {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	out := sim(list1, list2)
	fmt.Println(out)

	//Copy from input.txt to list3, list4
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(content), "\r\n")
	list3 := make([]int, 1000, 1000)
	list4 := make([]int, 1000, 1000)
	for i, line := range words {
		two := strings.Split(line, "   ")
		list3[i], _ = strconv.Atoi(two[0])
		list4[i], err = strconv.Atoi(two[1])
	}

	out = sim(list3, list4)
	fmt.Println(out)
}
