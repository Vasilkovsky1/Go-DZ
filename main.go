package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers [5]int
	sum := 0

	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
		sum = sum + numbers[i]
	}

	sort.Ints(numbers[:])

	fmt.Print("Отсортированные элементы: ")
	for i := 4; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()

	fmt.Println("Самое большое число:", numbers[4])
	fmt.Println("Самое маленькое число:", numbers[0])
	fmt.Println("Среднее арифметическое:", sum/5)
}