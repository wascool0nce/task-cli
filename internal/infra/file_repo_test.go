package infra

import (
	"os"
	"reflect"
	"task-cli/internal/domain"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	now := time.Now().Truncate(time.Second)
	tasks := []domain.Task{
		{
			Id:          1,
			Description: "Tests task 1",
			Status:      "todo",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			Id:          2,
			Description: "Tests task 2",
			Status:      "in-progress",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	t.Run("valid save", func(t *testing.T) {
		_, err := os.Create("tasks.json")
		if err != nil {
			t.Errorf("dont create file tasks.json")
		}
		repo := FileTaskRepository{path: "tasks.json"}
		if err := repo.Save(tasks[0]); err != nil {
			t.Errorf("in repo dont create new task")
		}
		if err := repo.Save(tasks[1]); err != nil {
			t.Errorf("in repo dont create new task")
		}
		data, err := repo.GetAll()
		if err != nil {
			t.Errorf("dont get all tasks")
		}

		ok := reflect.DeepEqual(tasks, data)
		if !ok {
			t.Errorf("struct dont equal")
		}
		err = os.Remove("tasks.json")
		if err != nil {
			t.Errorf("dont remove file tasks.json")
		}
	})

	t.Run("valid read", func(t *testing.T) {
		_, err := os.Create("tasks.json")
		if err != nil {
			t.Errorf("dont create file tasks.json")
		}
		repo := FileTaskRepository{path: "tasks.json"}
		if err := repo.Save(tasks[0]); err != nil {
			t.Errorf("in repo dont create new task")
		}
		if err := repo.Save(tasks[1]); err != nil {
			t.Errorf("in repo dont create new task")
		}
		data, err := repo.GetAll()
		if err != nil {
			t.Errorf("dont get all tasks")
		}

		ok := reflect.DeepEqual(tasks, data)
		if !ok {
			t.Errorf("struct dont equal")
		}
		err = os.Remove("tasks.json")
		if err != nil {
			t.Errorf("dont remove file tasks.json")
		}
	})
}
