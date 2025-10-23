package infra

type TaskRepository interface {
	Save(task Task) error
	GetAll() ([]Task, error)
}

type FileTaskRepository struct {
	path string
}

func (r FileRepository) Save(task Task) error {
	data, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.path, data, 0644)
}

func (r FileRepository) GetAll() ([]Task, error) {
	content, err := os.ReadFile(r.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
	}
	return nil, err
}
