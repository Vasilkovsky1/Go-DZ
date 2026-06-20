package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')

	// Приводим к нижнему регистру и разбиваем на слова
	words := strings.Fields(strings.ToLower(text))

	// Мапа для подсчета частот
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("\nСтатистика слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}