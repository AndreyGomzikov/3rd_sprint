package main

import "fmt"

func recursiveFactorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * recursiveFactorial(n-1)
}

func main() {
    num := 5
    result := recursiveFactorial(num)
    fmt.Printf("Факториал числа %d: %d\n", num, result)
}