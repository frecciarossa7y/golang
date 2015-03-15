// fizzbuzz問題（剰余を使わないバージョン）
package main

import (
    "fmt"
)

func main() {
    var n_3 int = 3
    var n_5 int = 5
    var fizz_str string = ""
    var buzz_str string = ""
    for i := 1; i < 100; i++ {
        fizz_str = ""
        buzz_str = ""
        if i == n_3 {
            fizz_str = "Fizz"
            n_3 += 3
        } 
        if i == n_5 { 
            buzz_str = "Buzz"
            n_5 += 5
        }
        if fizz_str == buzz_str {
            fmt.Println(i)
        } else {
            fmt.Println(fizz_str,buzz_str)
        }
    }
}
