package task

import (
	"errors"
	"fmt"
)

var ErrInvalidTaskID = errors.New("ID задачи должен быть больше нуля")
var ErrEmptyTaskTitle = errors.New("название задачи не должно быть пустым")

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func CreateTask(id int, title string) (Task, error) {
	if id <= 0 {
		return Task{}, ErrInvalidTaskID
	}
	if title == "" {
		return Task{}, ErrEmptyTaskTitle
	}

	return Task{ID: id, Title: title, Completed: false}, nil
}

func AddTask(tasks []Task, id int, title string) ([]Task, error) {
	task, err := CreateTask(id, title)
	if err != nil {
		return tasks, fmt.Errorf("добавление задачи %d: %w", id, err)
	}
	tasks = append(tasks, task)
	return tasks, nil
}
