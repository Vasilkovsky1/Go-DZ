package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(map[string]string{
		"message": "hello, world!",
	})
	if err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	log.Println("Сервер запущен: http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
