package service

import (
	"task-cli/internal/domain"
	"task-cli/internal/infra"
	"time"
)

const Local_path = "tasks.json"

func AddTask(decsription string) error {
	task := domain.Task{
		Id:          1,
		Description: decsription,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repo := infra.FileTaskRepository{Path: "tasks.json"}
	err := repo.Save(task)

	if err != nil {
		return err
	}

	return nil
}
