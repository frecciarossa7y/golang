// FizzBuzz問題（通常版）
package main

import (
    "fmt"
)

func main() {
    n := 1
    for i := 0; i < 100; i++ {
        switch {
        case n%15 == 0:
            fmt.Println("FizzBuzz")
        case n%5 == 0:
            fmt.Println("Buzz")
        case n%3 == 0:
            fmt.Println("Fizz")
        default:
            fmt.Println(n)
        }
        n++
    }
}
