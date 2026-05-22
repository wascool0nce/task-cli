package usecase

import (
	"errors"
	"task-cli/internal/domain"
	"time"
)

type TaskRepository interface {
	Load() ([]domain.Task, error)
	Save([]domain.Task) error
}

type TaskService struct {
	repo TaskRepository
	now  func() time.Time
}

func NewTaskService(repo TaskRepository, now func() time.Time) *TaskService {
	return &TaskService{
		repo: repo,
		now:  now,
	}
}

func (s *TaskService) Add(description string) error {
	if description == "" {
		return errors.New("опиcание задачи пустое")
	}
	// загружаем все задачи
	tasks, err := s.repo.Load()
	if err != nil {
		return err
	}

	// формируем новый ID
	newId := 1
	if len(tasks) != 0 {
		newId = tasks[len(tasks)-1].Id + 1
	}

	// создаем новую задачу
	newTask := domain.Task{
		Id:          newId,
		Description: description,
		Status:      domain.ToDo,
		CreatedAt:   s.now(),
		UpdatedAt:   s.now(),
	}

	// сохраняем ее
	tasks = append(tasks, newTask)

	err = s.repo.Save(tasks)
	if err != nil {
		return err
	}

	return nil
}
