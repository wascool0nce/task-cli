package infra

import (
	"os"
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
		},
		{
			Id:          2,
			Description: "Tests task 2",
			Status:      "in-progress",
			CreatedAt:   time.Now(),
		},
	}

	t.Run("valid save", func(t *testing.T) {
		_, err := os.Create("tasks.json")
		if err != nil {
			t.Errorf("dont creat file tasks.json")
		}
		repo := FileTaskRepository{path: "tasks.json"}
		if err := repo.Save(tasks[0]); err != nil {
			t.Errorf("in repo dont creat new task")
		}
	})

}
