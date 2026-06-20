
package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderPage(w http.ResponseWriter, page string) {
	tmpl, err := template.ParseFiles("templates/" + page)
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка отображения страницы", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "index.html")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "about.html")
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "services.html")
	})

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "gallery.html")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, "contact.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}