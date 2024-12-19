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
	// Find all don't()
	dont_parts := strings.Split(string(input_b), "don't()")
	do_parts := make([]string, len(dont_parts))
	// In the beginnig operations are enabled
	do_parts[0] = dont_parts[0]
	// Remove all beginnings in dont_parts untill there will be do()
	for i := 1; i < len(dont_parts); i++ {
		if len(strings.SplitN(dont_parts[i], "do()", 2)) != 1 {
			do_parts[i] = strings.SplitN(dont_parts[i], "do()", 2)[1]
		} else {
			do_parts[i] = ""
		}
	} // now in do_parts there are only enabled mul()
	for i := 0; i < 10; i++ {
		fmt.Println(len(do_parts[i]))
	}
	// Let's put them in one string and run code from d3p1
	enabled_string := ""
	for _, l := range do_parts {
		enabled_string += l
	}
	fmt.Println(enabled_string)
	// Find all mul(
	input_lines := strings.Split(enabled_string, "mul(")
	// Save all content in () if there is any
	for i, line := range input_lines {
		input_lines[i] = strings.Split(line, ")")[0]
	}
	// sum products of n,m in input_lines
	res := 0
	for _, line := range input_lines {
		res += product_handler(line)
	}
	fmt.Println(res)
}
