package main

import (
	"errors"
	"testing"
)

func TestCreateTaskSuccess(t *testing.T) {
	task, err := createTask(1, "Изучить тесты")

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
			_, err := createTask(tt.id, tt.title)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ошибка = %v, ожидали %v", err, tt.wantErr)
			}
		})
	}
}
