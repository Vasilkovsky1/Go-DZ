package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

type Post struct {
	ID          int
	Title       string
	Description string
	Content     string
	Author      string
	CreatedAt   string
}

type PageData struct {
	Title string
	User  string
	Posts []Post
	Post  Post
}

func initDB() {
	var err error

	db, err = sql.Open("sqlite", "site.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		created_at TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		message TEXT NOT NULL,
		created_at TEXT NOT NULL
	);
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentUser(r *http.Request) string {
	cookie, err := r.Cookie("username")
	if err != nil {
		return ""
	}

	return cookie.Value
}

func render(w http.ResponseWriter, r *http.Request, page string, data PageData) {
	data.User = getCurrentUser(r)

	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/"+page,
	)

	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы: "+err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "index.html", PageData{
		Title: "Главная",
	})
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		SELECT id, title, description, content, author, created_at
		FROM posts
		ORDER BY id DESC
	`)

	if err != nil {
		http.Error(w, "Ошибка получения статей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
		)

		if err != nil {
			http.Error(w, "Ошибка чтения статьи", http.StatusInternalServerError)
			return
		}

		posts = append(posts, post)
	}

	render(w, r, "articles.html", PageData{
		Title: "Статьи",
		Posts: posts,
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	render(w, r, "about.html", PageData{
		Title: "О нас",
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		_, err := db.Exec(`
			INSERT INTO contacts (name, email, message, created_at)
			VALUES (?, ?, ?, ?)
		`, name, email, message, time.Now().Format("2006-01-02 15:04"))

		if err != nil {
			http.Error(w, "Ошибка сохранения сообщения", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/contact", http.StatusSeeOther)
		return
	}

	render(w, r, "contact.html", PageData{
		Title: "Контакты",
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if getCurrentUser(r) != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := db.Exec(`
			INSERT INTO users (username, email, password)
			VALUES (?, ?, ?)
		`, username, email, password)

		if err != nil {
			http.Error(w, "Пользователь с таким email уже существует", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	render(w, r, "register.html", PageData{
		Title: "Регистрация",
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if getCurrentUser(r) != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var username string
		var storedPassword string

		err := db.QueryRow(`
			SELECT username, password
			FROM users
			WHERE email = ?
		`, email).Scan(&username, &storedPassword)

		if err == sql.ErrNoRows {
			http.Error(w, "Пользователь не найден", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
			return
		}

		if password != storedPassword {
			http.Error(w, "Пароль неверный", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "username",
			Value:    username,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	render(w, r, "login.html", PageData{
		Title: "Вход",
	})
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	username := getCurrentUser(r)

	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		description := r.FormValue("description")
		content := r.FormValue("content")

		_, err := db.Exec(`
			INSERT INTO posts (title, description, content, author, created_at)
			VALUES (?, ?, ?, ?, ?)
		`, title, description, content, username, time.Now().Format("2006-01-02 15:04"))

		if err != nil {
			http.Error(w, "Ошибка сохранения статьи", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/articles", http.StatusSeeOther)
		return
	}

	render(w, r, "create.html", PageData{
		Title: "Создать статью",
	})
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Некорректный ID статьи", http.StatusBadRequest)
		return
	}

	var post Post

	err = db.QueryRow(`
		SELECT id, title, description, content, author, created_at
		FROM posts
		WHERE id = ?
	`, id).Scan(
		&post.ID,
		&post.Title,
		&post.Description,
		&post.Content,
		&post.Author,
		&post.CreatedAt,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Статья не найдена", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}

	render(w, r, "post.html", PageData{
		Title: post.Title,
		Post:  post,
	})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/articles", articlesHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", createPostHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущен: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}