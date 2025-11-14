package domain

type TaskRepository interface {
	GetAll() ([]Task, error)
	Add(description, status string) error
	GetId(id int) (*Task, error)
	Update(id int, description, status string) error
	Delete(id int) error
}
