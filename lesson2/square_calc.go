package main

import "fmt"

func main()  {
	fmt.Println("Вычисляем площадь прямоугольника")
	var dimension1, dimension2 int
	fmt.Print("Введите длины сторон: ")
	fmt.Scan(&dimension1, &dimension2)

	fmt.Printf("Площадь прямоугольника: %d\n", dimension1 * dimension2)
}