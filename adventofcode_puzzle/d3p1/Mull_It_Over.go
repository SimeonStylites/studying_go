package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func product_handler(line string) int {
	two := strings.Split(line, ",")
	if len(two) != 2 {
		return 0
	}
	a, err := strconv.Atoi(two[0])
	if err != nil {
		return 0
	}
	b, err := strconv.Atoi(two[1])
	if err != nil {
		return 0
	}
	return a * b
}

func main() {
	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Find all mul(
	input_lines := strings.Split(string(input_b), "mul(")
	for i := 0; i < 100; i++ {
		fmt.Println(input_lines[i])
	}
	// Save all content in () if there is any
	for i, line := range input_lines {
		input_lines[i] = strings.Split(line, ")")[0]
	}
	for i := 0; i < 100; i++ {
		fmt.Println(input_lines[i])
	}
	// sum products of n,m in input_lines
	res := 0
	for _, line := range input_lines {
		res += product_handler(line)
	}
	fmt.Println(res)
}
