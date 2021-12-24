package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вычисляем диаметр и длину окружности")
	fmt.Print("Введите площадь круга: ")
	
	var area float64
	fmt.Scan(&area)

	diameter := math.Sqrt(area * 4 / math.Pi)
	length := diameter * math.Pi

	fmt.Printf("Диаметр окружности: %f. Длина окружности: %f\n", diameter, length)
}