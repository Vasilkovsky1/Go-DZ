package main

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalizeWords(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		r := []rune(strings.ToLower(words[i]))

		if len(r) > 0 {
			r[0] = unicode.ToUpper(r[0])
		}

		words[i] = string(r)
	}

	return strings.Join(words, " ")
}

func main() {
	var s string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&s)

	fmt.Println(capitalizeWords(s))
}