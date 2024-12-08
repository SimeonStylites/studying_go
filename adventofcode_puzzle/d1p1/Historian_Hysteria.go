package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diff(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	res := 0
	for i, _ := range a {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}
	return res
}

func sim(a, b []int) int {
	fmt.Print(a, b)
	return 0
}

func main() {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	out := diff(list1, list2)
	fmt.Println(out)

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
	fmt.Println(diff(list3, list4))

	fmt.Println(sim(list1, list2))
}
