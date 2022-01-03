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
	fibo := fibonacci()
	fmt.Printf("Результат: %d\n", fibo(number))
}

func fibonacci() func(uint) uint64 {
	var cache = make(map[uint]uint64)
	var fibo func(uint) uint64

	fibo = func(n uint) (result uint64) {
		if result, ok := cache[n]; ok {
			return result
		}
		defer func(){
			cache[n] = result
		}()
		if n <= 1 {
			return uint64(n)
		} else {
			return fibo(n - 1) + fibo(n - 2)
		}	
	}
	return fibo
}