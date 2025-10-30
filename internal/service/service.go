package service

import (
	"task-cli/internal/domain"
	"task-cli/internal/infra"
)

const Local_path = "tasks.json"

func AddTask(decsription string) {
	task := domain.Task{
		Id:          1,
		Description: decsription,
		Status:      "todo",
	}

	repo := infra.FileTaskRepository{path: Local_path}

}
