package main

import (
	"fmt"
)

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		fmt.Println("Ошибка: деление на ноль")
		return 0
	}
	return a / b
}

func main() {
	sum := applyOperation(5, 3, add)
	fmt.Printf("Сумма: %d\n", sum)

	difference := applyOperation(5, 3, subtract)
	fmt.Printf("Разность: %d\n", difference)

	product := applyOperation(5, 3, multiply)
	fmt.Printf("Произведение: %d\n", product)

	quotient := applyOperation(6, 3, divide)
	fmt.Printf("Частное: %d\n", quotient)

	applyOperation(6, 0, divide)
}