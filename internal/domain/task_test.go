package domain

import (
	"testing"
	"time"
)

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   string
		expected bool
	}{
		{"valid todo", toDo, true},
		{"valid in-progress", inProgress, true},
		{"valid done", done, true},
		{"invalid", "archived", false},
		{"invalid TODO", "TODO", false},
		{"invalid zero value", "", false},
		{"invalid ' done '", " done ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidStatus(tt.status)
			if result != tt.expected {
				t.Errorf("isValidStatus(%q) = %v; want %v", tt.status, result, tt.expected)
			}
		})
	}
}

func TestChangeStatus(t *testing.T) {
	task := Task{
		Id:          1,
		Description: "Test Task",
		Status:      toDo,
		CreatedAt:   time.Now(),
	}

	t.Run("valid status", func(t *testing.T) {
		err := task.ChangeStatus(inProgress)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if task.Status != inProgress {
			t.Errorf("expected status %q, got %q", inProgress, task.Status)
		}
		if !task.UpdatedAt.After(task.CreatedAt) {
			t.Errorf("UpdatedAt not after CreatedAt")
		}
	})

	t.Run("invalid status", func(t *testing.T) {
		err := task.ChangeStatus("invalid status")
		if err == nil {
			t.Errorf("excepted error, got nil")
		}
	})

	t.Run("status zero value", func(t *testing.T) {
		err := task.ChangeStatus("")
		if err == nil {
			t.Errorf("excepted error, got nil")
		}
	})
}
