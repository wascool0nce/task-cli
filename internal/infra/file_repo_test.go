package infra

import "time"

func TestSave(t *testing.T) {
	tasks := [2]Task {
		{
			Id: 1,
			Description: "Tests task 1",
			Status: "todo",
			CreatedAt: time.Now()
		},
		{
			Id: 2,
			Description: "Tests task 2",
			Status: "in-progress",
			CreatedAt: time.Now()
		}
	}

	t.Run("valid save", func(t *testing.T) {
		_, err := os.Create("tasks.json")
		repo := TaskFileRepository{path: "tasks.json"}
		tasks, err
	})
}
