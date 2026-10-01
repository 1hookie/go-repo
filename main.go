package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tasks/example", exampleTaskHandler)
	http.HandleFunc("GET /tasks", listTaskHandler)
	http.HandleFunc("POST /tasks", createTaskHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ", err)
	}
}
