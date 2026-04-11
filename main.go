package main

import "fmt"

func main() {
	var s string
	open := 0
	close := 0
	balance := 0
	ok := true

	fmt.Print("Введите строку: ")
	fmt.Scanln(&s)

	for _, ch := range s {
		if ch == '(' {
			open++
			balance++
		}
		if ch == ')' {
			close++
			balance--
		}

		if balance < 0 {
			ok = false
		}
	}

	if balance != 0 {
		ok = false
	}

	if ok {
		fmt.Printf("Скобки расставлены верно, %d открывающиеся, %d закрывающиеся\n", open, close)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся\n", open, close)
	}
}