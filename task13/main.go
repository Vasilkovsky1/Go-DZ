package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/igorvasilkovskij/Downloads/logfile.log")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	errorCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(strings.ToLower(line), "error") {
			errorCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Printf("Количество строк, содержащих слово 'error': %d\n", errorCount)
}