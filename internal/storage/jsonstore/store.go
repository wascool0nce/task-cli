package jsonstore

import (
	"encoding/json"
	"os"
	"task-cli/internal/domain"
)

type Store struct {
	path string
}

func New(path string) *Store {
	return &Store{path: path}
}

func (s *Store) Load() ([]domain.Task, error) {
	data, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		err := s.Save([]domain.Task{})
		if err != nil {
			return nil, err
		}
		return []domain.Task{}, nil
	}

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []domain.Task{}, nil
	}

	var tasks []domain.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *Store) Save(tasks []domain.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(s.path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
