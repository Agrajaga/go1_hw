package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var n int
	fmt.Printf("Программа ищет простые чисела от 0 до N\n")
	fmt.Printf("Укажите число N: ")

	if _, err := fmt.Scanln(&n); err != nil {
		return err
	}

	if n < 0 {
		return errors.New("expected N > 0") 
	}

	runNaiveMethod(n)
	
	return nil
}

func runNaiveMethod(maxNumber int) {
	for number := 1; number <= int(maxNumber); number++ {
		if isPrimeNumber(number) {
			fmt.Println(number)
		}
	}
}

func isPrimeNumber(number int) bool {
	switch number {
	case 1, 2:
		return true
	default:
		for divider := 2; divider < number / 2; divider++ {
			if number % divider == 0 {
				return false
			}
		} 
		return true
	}
}