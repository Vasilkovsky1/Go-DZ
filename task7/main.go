package main

import "fmt"

// Добавление города
func addCity(cities []string, city string) []string {
	return append(cities, city)
}

// Удаление города по имени
func removeCity(cities []string, city string) []string {
	for i, c := range cities {
		if c == city {
			return append(cities[:i], cities[i+1:]...)
		}
	}
	return cities
}

// Поиск города
func findCity(cities []string, city string) bool {
	for _, c := range cities {
		if c == city {
			return true
		}
	}
	return false
}

func main() {
	// Исходный список городов
	cities := []string{"Москва", "Санкт-Петербург", "Казань"}

	fmt.Println("Исходный список:", cities)

	// Добавление города
	cities = addCity(cities, "Новосибирск")
	fmt.Println("После добавления:", cities)

	// Поиск города
	cityToFind := "Казань"
	if findCity(cities, cityToFind) {
		fmt.Printf("Город %q найден.\n", cityToFind)
	} else {
		fmt.Printf("Город %q не найден.\n", cityToFind)
	}

	// Удаление города
	cities = removeCity(cities, "Санкт-Петербург")
	fmt.Println("После удаления:", cities)
}