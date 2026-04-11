package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	fmt.Println("Количество символов:", utf8.RuneCountInString(s))
}