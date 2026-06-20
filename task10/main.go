package main

import (
	"fmt"
)

// Структура Book
type Book struct {
	Title  string
	Author string
	Year   int
}

// Метод GetInfo
func (b Book) GetInfo() string {
	return fmt.Sprintf("\"%s\" by %s, %d", b.Title, b.Author, b.Year)
}

func main() {
	// Создание экземпляра книги
	book := Book{
		Title:  "1984",
		Author: "George Orwell",
		Year:   1949,
	}

	// Вывод информации о книге
	fmt.Println(book.GetInfo())
}