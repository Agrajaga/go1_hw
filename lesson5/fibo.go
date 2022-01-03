package main

import (
	"fmt"
	"os"
)

func main() {
	var number uint
	
	fmt.Printf("Рассчитываем N-ый член ряда Фибоначчи\n")
	fmt.Printf("Введите число N: ")
	if _, err := fmt.Scanln(&number); err != nil {
		fmt.Printf("Введено некорректное значение\n")
		os.Exit(1)
	}

	fmt.Printf("Результат: %d\n", fibo(number))
}

func fibo(n uint) uint  {
	if n < 2 {
		return n
	} else {
		return fibo(n - 1) + fibo(n - 2)
	}
}