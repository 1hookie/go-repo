package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func createTask(id int, title string) (Task, error) {
	if id <= 0 || title == "" {
		return Task{}, errors.New("ID должен быть > 0 и название не должно быть пустым")
	}
	return Task{ID: id, Title: title, Completed: false}, nil
}

func main() {
	task, err := createTask(1, "Изучить ошибки в GoLang")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(task)
	}

	task, err = createTask(2, "Закрепить тему Errors в GoLang")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(task)
	}

	task, err = createTask(0, "Проверить обработчик ошибки 0-го айди")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(task)
	}

	//Проверить обработчик ошибки пустого названия
	task, err = createTask(4, "")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(task)
	}
}
