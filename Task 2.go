package main

import (
	"fmt"
	"math"
)

func getRectangleData(width, height int) (area int, perimeter int, diagonal float64) {
	// Вычисляем площадь
	area = width * height

	perimeter = 2 * (width + height)

	diagonal = math.Sqrt(float64(width*width + height*height))

	return area, perimeter, diagonal
}

func main() {
	width := 4
	height := 5

	area, perimeter, diagonal := getRectangleData(width, height)

	fmt.Printf("Площадь: %d\n", area)
	fmt.Printf("Периметр: %d\n", perimeter)
	fmt.Printf("Диагональ: %.4f\n", diagonal)
}