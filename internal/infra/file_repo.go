package infra

import (
	"encoding/json"
	"errors"
	"os"
	"sort"
	"task-cli/internal/domain"
	"time"
)

type TaskRepository interface {
	Save(task domain.Task) error
	GetAll() ([]domain.Task, error)
}

type FileTaskRepository struct {
	path string
}

func NewFileTaskRepository(path string) *FileTaskRepository {
	return &FileTaskRepository{path: path}
}

func (r *FileTaskRepository) read() ([]domain.Task, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []domain.Task{}, nil
	}
	var tasks []domain.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, err
}

func (r *FileTaskRepository) save(tasks []domain.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, data, 0644)
}

func (r FileTaskRepository) Add(description, status string) error {

	tasks, err := r.read()
	if err != nil {
		return err
	}

	newId := 1
	if len(tasks) > 1 {
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Id < tasks[j].Id
		})
		newId = tasks[len(tasks)-1].Id + 1
	}
	task := domain.Task{
		Id:          newId,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, task)

	return r.save(tasks)
}

func (r FileTaskRepository) GetAll() ([]domain.Task, error) {
	return r.read()
}

func (r FileTaskRepository) GetId(id int) (*domain.Task, error) {
	tasks, err := r.read()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.Id == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}
