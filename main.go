package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks/example", exampleTaskHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера: ", err)
	}
	//http.HandleFunc("/health", healthHandler)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Println("Ошибка запуска сервера: ", err)
	//}
}
