package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	count := 0

	fmt.Print("Введите строку: ")
	fmt.Scanln(&s)

	s = strings.ToLower(s)

	for _, буква := range s {
		if буква == 'а' || буква == 'е' || буква == 'ё' || буква == 'и' ||
			буква == 'о' || буква == 'у' || буква == 'ы' || буква == 'э' ||
			буква == 'ю' || буква == 'я' {
			count++
		}
	}

	fmt.Println("Количество гласных:", count)
}