package domain

import (
	"errors"
	"time"
)

const (
	toDo       = "todo"
	inProgress = "in-progress"
	done       = "done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
