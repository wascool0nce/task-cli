package infra

import (
	"encoding/json"
	"os"
	"task-cli/internal/domain"
)

type TaskRepository interface {
	Save(task domain.Task) error
	GetAll() ([]domain.Task, error)
}

type FileTaskRepository struct {
	path string
}

func (r FileTaskRepository) Save(task domain.Task) error {
	tasks, _ := r.GetAll()
	tasks = append(tasks, task)
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.path, data, 0644)
}

func (r FileTaskRepository) GetAll() ([]domain.Task, error) {
	content, err := os.ReadFile(r.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Task{}, nil
		}
		return nil, err
	}

	var tasks []domain.Task
	if err := json.Unmarshal(content, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
