package main

import (
	"encoding/json"
	"fmt"
	"learn/task"
	"log"
	"net/http"
)

var tasks = []task.Task{
	{ID: 1, Title: "Изучить JSON", Completed: false},
	{ID: 2, Title: "Написать REST API", Completed: true},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func exampleTaskHandler(w http.ResponseWriter, r *http.Request) {
	newTask, err := task.CreateTask(1, "Изучить JSON")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newTask); err != nil {
		log.Println("ошибка кодирования JSON:", err)
	}
}

func listTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println("ошибка кодирования JSON:", err)
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
	}
}
