package jsonstore

import (
	"encoding/json"
	"os"
	"path/filepath"

	"task-cli/internal/entity"
)

type Store struct {
	path string
}

func New(path string) *Store {
	return &Store{path: path}
}

// Load Читаем файл по пути
// Если ничего не найдено ты мы возвращаем пустой slice структур
// Если ошибка то возвращаеи nil и ошибку
// Если длинна  данных 0 то мы возвращаем пустой slice структур
// Проверяем возможность преобразовать json если не получилось возвращаем ошибку
func (s *Store) Load() ([]entity.Task, error) {
	data, err := os.ReadFile(s.path)
	if os.IsNotExist(err) {
		return []entity.Task{}, nil
	}

	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []entity.Task{}, nil
	}

	var tasks []entity.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *Store) Save(tasks []entity.Task) error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}
