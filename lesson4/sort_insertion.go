package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	list := []int{12,6,7,3,2,1,9,15,10,14,3}

	fmt.Println("Unsorted list:")
	fmt.Println(list)

	sort(list)

	fmt.Println("Sorted list:")
	fmt.Println(list)
	return nil
}

func sort(list []int) {
	for i := 1; i < len(list); i++ {
		for j := i - 1; j >= 0; j-- {
			if list[i] >= list[j] {
				move(list, i, j + 1)
				break
			}
			if j == 0 {
				move(list, i, 0)
			}
		}	
	}	
}

func move(list []int, fromIndex int, toIndex int)  {
	tmpValue := list[fromIndex]	
	for i := fromIndex - 1; i >= toIndex; i-- {
		list[i + 1] = list[i]
	}
	list[toIndex] = tmpValue
}