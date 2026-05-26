package domain

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	ToDo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

func NewTaskStatus(s string) (TaskStatus, error) {
	switch s {
	case "todo":
		return ToDo, nil
	case "in-progress":
		return InProgress, nil
	case "done":
		return Done, nil
	}
	return "", fmt.Errorf("unknown task status: %s", s)
}

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
