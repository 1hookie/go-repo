package main

import "fmt"

func createTask(id int, title string) (Task, error) {
	if id <= 0 {
		return Task{}, ErrInvalidTaskID
	}
	if title == "" {
		return Task{}, ErrEmptyTaskTitle
	}

	return Task{ID: id, Title: title, Completed: false}, nil
}

func addTask(tasks []Task, id int, title string) ([]Task, error) {
	task, err := createTask(id, title)
	if err != nil {
		return tasks, fmt.Errorf("добавление задачи %d: %w", id, err)
	}
	tasks = append(tasks, task)
	return tasks, nil
}
