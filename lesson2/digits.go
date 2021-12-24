package main

import "fmt"

func main()  {
	fmt.Print("Введите трехзначное число: ")
	var number uint
	fmt.Scan(&number)
	if number < 1000 && number > 99 {
		// используем особенность, что деление двух целых чисел дает целое число
		hundreds := number / 100
		tens := (number - hundreds * 100) / 10
		ones := (number - hundreds * 100 - tens * 10)
		fmt.Println("Состав числа:")
		fmt.Printf("Сотен: %d\n", hundreds)
		fmt.Printf("Десятков: %d\n", tens)
		fmt.Printf("Единиц: %d\n", ones)
	} else {
		fmt.Printf("Необходимо ввести трехзначное число!\n")
	}

}