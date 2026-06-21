package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
}

func insertUsers(db *sql.DB) {
	query := `
	INSERT OR IGNORE INTO users (email, password) VALUES
	('user1@mail.com', '12345'),
	('user2@mail.com', 'qwerty'),
	('admin@mail.com', 'admin123');
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Ошибка добавления пользователей:", err)
	}
}

func checkUser(db *sql.DB, email, password string) string {
	var storedPassword string

	err := db.QueryRow(
		"SELECT password FROM users WHERE email = ?",
		email,
	).Scan(&storedPassword)

	if err == sql.ErrNoRows {
		return "пользователь не найден"
	}

	if err != nil {
		return "ошибка базы данных"
	}

	if storedPassword != password {
		return "пароль неверный"
	}

	return "пароль верный"
}

func main() {
	db, err := sql.Open("sqlite", "users.db")
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}
	defer db.Close()

	createTable(db)
	insertUsers(db)

	var email string
	var password string

	fmt.Print("Введите email: ")
	fmt.Scanln(&email)

	fmt.Print("Введите пароль: ")
	fmt.Scanln(&password)

	result := checkUser(db, email, password)
	fmt.Println(result)
}