package main

import "fmt"

func generateSumFunc() func(int) int {
	total := 0

	return func(n int) int {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}
		total += sum
		return total
	}
}

func main() {
	sumFunc := generateSumFunc()

	fmt.Println(sumFunc(3))
	fmt.Println(sumFunc(5))
	fmt.Println(sumFunc(2))
}