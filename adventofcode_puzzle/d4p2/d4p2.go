package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func findXMASes2(lines []string) int {
	res := 0
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == byte('A') {
				res += check(lines, i, j)
			}
		}
	}
	return res
}

func check(lines []string, i, j int) int {
	n := len(lines)
	m := len(lines[0])
	if i > 0 && i < n-1 && j > 0 && j < m-1 {
		M1 := lines[i+1][j+1]
		S1 := lines[i-1][j-1]
		if ((M1 == byte('M')) && (S1 == byte('S'))) || ((M1 == byte('S')) && (S1 == byte('M'))) {
			M2 := lines[i+1][j-1]
			S2 := lines[i-1][j+1]
			if ((M2 == byte('M')) && (S2 == byte('S'))) || ((M2 == byte('S')) && (S2 == byte('M'))) {
				return 1
			}
		}
	}
	return 0
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\r\n")
	fmt.Println(findXMASes2(lines))
}
