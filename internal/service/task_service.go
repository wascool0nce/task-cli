package service

import "task-cli/internal/domain"

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(description, status string) error {
	return s.repo.Add(description, status)
}

func (s *TaskService) ListTasks() ([]domain.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) UpdateTask(id int, description, status string) error {
	return s.repo.Update(id, description, status)
}

func (s *TaskService) GetById(id int) (*domain.Task, error) {
	return s.repo.GetId(id)
}
