package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Горутина %d начала работу\n", id)

	// Имитация выполнения задачи
	time.Sleep(1 * time.Second)

	fmt.Printf("Горутина %d завершила работу\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Запускаем 3 горутины
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("Все горутины завершили работу")
}