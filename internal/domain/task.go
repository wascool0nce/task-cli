package domain

import (
	"time"
	"errors"
)

const (
	toDo = "todo"
	inProgress = "in-progress"
	done = "done"
)

type Task struct {
	Id			int
	Description	string
	Status		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func (t *Task) ChangeStatus(status string) error {
	if !isValidStatus(status) {
		return errors.New("invalid status: " + status)
	}
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}

func isValidStatus(status string) bool {
	switch status {
	case toDo, inProgress, done:
		return true
	default:
		return false
	}
}
