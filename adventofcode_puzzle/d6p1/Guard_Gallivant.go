package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// 4 vectors define 4 directions, can be changing in cycle
var vecs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func guard_moving(initial_map []string) []string {
	n, m := guardposition(initial_map)
	mod4 := 0                                             //initial direction will be vecs[mod4]=vecs[0]={-1,0}
	stop := initial_map[n+vecs[mod4][0]][m+vecs[mod4][1]] // what is in front of guard
	i, j := n, m
	for i > -1 && j > -1 && i < 130 && j < 130 { //Are we in map after changing directions?
		for stop != byte('#') {
			i, j = i+vecs[mod4][0], j+vecs[mod4][1]
			// Are we in map?
			if i+vecs[mod4][0] < 0 || j+vecs[mod4][1] < 0 || i+vecs[mod4][0] > 129 || j+vecs[mod4][1] > 129 {
				return initial_map
			}
			stop = initial_map[i+vecs[mod4][0]][j+vecs[mod4][1]]
			initial_map[i] = initial_map[i][:j] + "X" + initial_map[i][j+1:]
		}
		stop = initial_map[i-vecs[mod4][0]][j-vecs[mod4][1]]
		mod4 = (mod4 + 1) % 4
		stop = initial_map[i+vecs[mod4][0]][j+vecs[mod4][1]]
	}
	return initial_map
}

// Finding where is ^ - initial guard position
func guardposition(initial_map []string) (i, j int) {
	for i, line := range initial_map {
		j = strings.Index(line, "^")
		if j != -1 {
			return i, j
		}
	}
	return -1, -1
}

// Counting X in []string
func distinct_positions(initial_map []string) int {
	res := 0
	for _, line := range initial_map {
		for _, l := range line {
			if l == 'X' {
				res += 1
			}
		}
	}
	return res
}

func main() {
	input, err := os.ReadFile("input.txt") //Read file
	if err != nil {
		log.Fatal(err)
	}
	initial_map := strings.Split(string(input), "\r\n") //Make []string
	map_with_route := guard_moving(initial_map)         //Draw X everywhere the guard steps

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, line := range map_with_route {
		fmt.Fprint(f, line+"\r\n") //Print map with route in data.txt
	}

	fmt.Println(distinct_positions(map_with_route) + 2) // Count all X, +2: 1.initial position, 2.last position
}
