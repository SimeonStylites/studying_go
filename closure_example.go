package main

import "fmt"

func fib() func() int {
    x1, x2 := 0, 1
    // возвращаемая функция замыкает x1, x2
    return func() int {
        x1, x2 = x2, x1+x2
        return x1
    }
}

func main() {
    f := fib() // получили функцию-замыкание. f() — захватила x1, x2. x1 = 0, x2 = 1
    fmt.Println(f()) // x1 = 1, x2 = 1
    fmt.Println(f()) // x1 = 1, x2 = 2
    fmt.Println(f()) // x1 = 2, x2 = 3
    fmt.Println(f()) // x1 = 3, x2 = 5
    fmt.Println(f()) // x1 = 5, x2 = 8
    fmt.Println(f()) // x1 = 8, x2 = 13

} 