package main

import "fmt"

func main() {
    a := "some text"
    defer func(s string) {
        fmt.Println(s)
    }(a)
    a = "another text"
}