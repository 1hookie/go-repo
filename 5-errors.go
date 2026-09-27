package main

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
