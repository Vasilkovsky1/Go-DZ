package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создание и заполнение массива
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1 // числа от 1 до 100
	}

	// Копирование массива в слайс
	slice := make([]int, len(arr))
	copy(slice, arr[:])

	// Сортировка слайса
	sort.Ints(slice)

	// Вывод результатов
	fmt.Println("Исходный массив:", arr)
	fmt.Println("Отсортированный слайс:", slice)
}