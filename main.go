package main

import (
	"fmt"
	"learn/task"
)

func main() {
	t, err := task.CreateTask(1, "Изучить ошибки в GoLang")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(t)
	}

	t, err = task.CreateTask(2, "Закрепить тему Errors в GoLang")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(t)
	}

	t, err = task.CreateTask(0, "Проверить обработчик ошибки 0-го айди")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(t)
	}

	//Проверить обработчик ошибки пустого названия
	t, err = task.CreateTask(4, "")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(t)
	}
}
