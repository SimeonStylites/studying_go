package main

import "fmt"

const (
    one = iota*2+1
    three
    five
    seven
    nine
    eleven
)

func main() {
    fmt.Println(one, three, five, seven, nine, eleven)
} 