package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var operand1, operand2, result float64
	var operation string
	var err error
	var lenOperation int

	fmt.Printf("Введите первое число: ")
	_, err = fmt.Scanln(&operand1)
	if err != nil {
		fmt.Printf("Введено некорректное значение!\n")
		os.Exit(1)
	}

	fmt.Printf("Введите второе число: ")
	_, err = fmt.Scanln(&operand2)
	if err != nil {
		fmt.Printf("Введено некорректное значение!\n")
		os.Exit(1)
	}

	fmt.Printf("Введите арифметическую операцию (+, -, *, /, ^): ")
	lenOperation, err = fmt.Scanln(&operation)
	if err != nil || lenOperation != 1  {
		fmt.Printf("Введено некорректное значение!\n")
		os.Exit(1)
	}

	switch operation {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Printf("Деление на ноль недопустимо\n")
			os.Exit(1)
		}
		result = operand1 / operand2
	case "^":
		result = math.Pow(operand1, operand2)
	default:
		fmt.Printf("Операция указана неверно\n")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", result)
}