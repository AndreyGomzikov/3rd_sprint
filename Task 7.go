package main

import "fmt"

func main() {
    numbers := []int{1, 5, 8, 2, 9}

    sum := func(nums []int) int {
        total := 0
        for _, num := range nums {
            total += num
        }
        return total
    }

    result := sum(numbers)
    fmt.Println("Сумма чисел:", result)
}