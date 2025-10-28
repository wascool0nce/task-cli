package infra

import (
	"os"
	"reflect"
	"task-cli/internal/domain"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	tasks := [2]domain.Task{
		{
			Id:          1,
			Description: "Tests task 1",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          2,
			Description: "Tests task 2",
			Status:      "in-progress",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
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
		// err = os.Remove("tasks.json")
		// if err != nil {
		// 	t.Errorf("dont remove file tasks.json")
		// }
	})
}
