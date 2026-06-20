package main

import "testing"

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)

	if err != nil {
		t.Fatalf("Не ожидалась ошибка: %v", err)
	}

	if result != 5 {
		t.Fatalf("Ожидалось 5, получено %v", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(10, 0)

	if err == nil {
		t.Fatal("Ожидалась ошибка при делении на ноль")
	}
}