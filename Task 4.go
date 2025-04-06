package main

import "fmt"

func multiplierFactory(factor int) func(int) int {
    return func(x int) int {
        return factor * x
    }
}

func main() {
    double := multiplierFactory(2)
    triple := multiplierFactory(3)
    quintuple := multiplierFactory(5)

    fmt.Println(double(10))    // 20 (2 * 10)
    fmt.Println(triple(10))    // 30 (3 * 10)
    fmt.Println(quintuple(10)) // 50 (5 * 10)

}