package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("GET /tasks/example", exampleTaskHandler)
	http.HandleFunc("GET /tasks", listTaskHandler)
	http.HandleFunc("POST /tasks", createTaskHandler)
	http.HandleFunc("GET /tasks/{id}", getTaskHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ", err)
	}
}
