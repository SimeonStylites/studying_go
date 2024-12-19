package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func findXMASes(lines []string) int {
	res := 0
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == byte('X') {
				res += check(lines, i, j)
			}
		}
	}
	return res
}

var directs = [8][2]int{
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func check(lines []string, i, j int) int {
	res := 0
	n := len(lines)
	m := len(lines[0])
	vectors_exist := possible(i, j, n, m)
	for d, p := range vectors_exist {
		if p {
			M := lines[i+directs[d][0]][j+directs[d][1]]
			A := lines[i+2*directs[d][0]][j+2*directs[d][1]]
			S := lines[i+3*directs[d][0]][j+3*directs[d][1]]
			if (M == byte('M')) && (A == byte('A')) && (S == byte('S')) {
				res += 1
			}
		}
	}
	return res
}

func possible(i, j, n, m int) []bool {
	vectors_exist := []bool{true, true, true, true, true, true, true, true}
	if i < 3 {
		vectors_exist[1], vectors_exist[2], vectors_exist[3] = false, false, false
	}
	if i > n-4 {
		vectors_exist[5], vectors_exist[6], vectors_exist[7] = false, false, false
	}
	if j < 3 {
		vectors_exist[3], vectors_exist[4], vectors_exist[5] = false, false, false
	}
	if j > m-4 {
		vectors_exist[7], vectors_exist[0], vectors_exist[1] = false, false, false
	}
	return vectors_exist
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\r\n")
	fmt.Println(findXMASes(lines))
}
