package main

import (
	"encoding/json"
	"fmt"
	"learn/task"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func exampleTaskHandler(w http.ResponseWriter, r *http.Request) {
	newTask, err := task.CreateTask(1, "Изучить JSON")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}
