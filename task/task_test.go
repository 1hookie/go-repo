package task

import (
	"errors"
	"testing"
)

func TestCreateTaskSuccess(t *testing.T) {
	task, err := CreateTask(1, "Изучить тесты")

	if err != nil {
		t.Fatalf("не ожидали ошибку, получили: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("ID = %d, ожидали 1", task.ID)
	}

	if task.Title != "Изучить тесты" {
		t.Errorf("Title = %q, ожидали %q", task.Title, "Изучить тесты")
	}

	if task.Completed {
		t.Error("новая задача должна быть невыполненной")
	}
}

func TestAddTaskSuccess(t *testing.T) {
	tasks := []Task{}
	result, err := AddTask(tasks, 1, "Buy a token")

	if err != nil {
		t.Fatalf("не ожидали ошибку, получили: %v", err)
	}
	if len(result) != 1 {
		t.Errorf("задача не добавилась, длина списка %d", len(result))
	}
	if result[0].ID != 1 {
		t.Errorf("ID = %d, ожидали 1", result[0].ID)
	}
	if result[0].Title != "Buy a token" {
		t.Errorf("Title = %q, ожидали %q", result[0].Title, "Buy a token")
	}

}

func TestAddTaskInvalidInput(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		title   string
		wantErr error
	}{
		{
			name:    "нулевой ID",
			id:      0,
			title:   "задача",
			wantErr: ErrInvalidTaskID,
		},
		{
			name:    "отрицательный ID",
			id:      -1,
			title:   "задача",
			wantErr: ErrInvalidTaskID,
		},
		{
			name:    "пустая задача",
			id:      1,
			title:   "",
			wantErr: ErrEmptyTaskTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AddTask([]Task{}, tt.id, tt.title)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ошибка = %v, ожидали %v", err, tt.wantErr)
			}
			if len(result) != 0 {
				t.Errorf("задача добавилась, длина списка %d", len(result))
			}

		})
	}
}

func TestCreateTaskInvalidInput(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		title   string
		wantErr error
	}{
		{
			name:    "нулевой ID",
			id:      0,
			title:   "задача",
			wantErr: ErrInvalidTaskID,
		},
		{
			name:    "отрицательный ID",
			id:      -1,
			title:   "задача",
			wantErr: ErrInvalidTaskID,
		},
		{
			name:    "пустая задача",
			id:      1,
			title:   "",
			wantErr: ErrEmptyTaskTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateTask(tt.id, tt.title)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ошибка = %v, ожидали %v", err, tt.wantErr)
			}

		})
	}
}
