package domain

import "time"

type TaskStatus string

const (
	ToDo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
