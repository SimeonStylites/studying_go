package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        found := false

        // проверяем, что кратно 3, и выводим Fizz
        if i%3 == 0 {
            fmt.Printf("Fizz")
            found = true
        }
        // проверяем, что кратно 5, и выводим Buzz
        if i%5 == 0 {
            fmt.Printf("Buzz")
            found = true
        }
        // если число кратно 3 и 5, выводим FizzBuzz

        if !found {
            // если не выполнилось ни одно из условий, выводим число
            fmt.Println(i)
            continue
        }

        fmt.Println()
    }
}