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

func (s *TaskService) Add(description string) (int, error) {
	if description == "" {
		return 0, errors.New("task description must not be empty")
	}
	// Load all tasks.
	tasks, err := s.repo.Load()
	if err != nil {
		return 0, err
	}

	// Build the next task ID.
	newId := 1
	if len(tasks) != 0 {
		newId = tasks[len(tasks)-1].Id + 1
	}

	// Create a new task.
	newTask := domain.Task{
		Id:          newId,
		Description: description,
		Status:      domain.ToDo,
		CreatedAt:   s.now(),
		UpdatedAt:   s.now(),
	}

	// Save the new task.
	tasks = append(tasks, newTask)

	err = s.repo.Save(tasks)
	if err != nil {
		return 0, err
	}

	return newTask.Id, nil
}

func (s *TaskService) List(filter domain.TaskStatus) ([]domain.Task, error) {
	tasks, err := s.repo.Load()
	if err != nil {
		return nil, err
	}
	var filterTasks []domain.Task

	if filter != "" {
		for _, task := range tasks {
			if task.Status == filter {
				filterTasks = append(filterTasks, task)
			}
		}
	} else {
		return tasks, nil
	}
	return filterTasks, nil
}

func (s *TaskService) Delete(taskID int) (bool, error) {
	tasks, err := s.repo.Load()
	isDelete := false
	if err != nil {
		return isDelete, err
	}
	for idx, task := range tasks {
		if task.Id == taskID {
			copy(tasks[idx:], tasks[idx+1:])
			tasks = tasks[:len(tasks)-1]
			isDelete = true
		}
	}
	err = s.repo.Save(tasks)
	if err != nil {
		return false, err
	}
	return isDelete, nil
}

func (s *TaskService) Update(taskID int, newDescription string) (bool, error) {
	isUpdate := false
	tasks, err := s.repo.Load()
	if err != nil {
		return isUpdate, err
	}
	for idx, task := range tasks {
		if task.Id == taskID {
			tasks[idx].Description = newDescription
			tasks[idx].UpdatedAt = s.now()

			isUpdate = true
		}
	}

	err = s.repo.Save(tasks)
	if err != nil {
		return false, err
	}
	return isUpdate, nil
}

func (s *TaskService) UpdateStatus(taskID int, status domain.TaskStatus) (bool, error) {
	isUpdate := false
	tasks, err := s.repo.Load()
	if err != nil {
		return isUpdate, err
	}
	for idx, task := range tasks {
		if task.Id == taskID {
			tasks[idx].Status = status
			tasks[idx].UpdatedAt = s.now()
			isUpdate = true
		}
	}
	err = s.repo.Save(tasks)
	if err != nil {
		return false, err
	}
	return isUpdate, nil
}
