package cli

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"task-cli/internal/domain"
)

type TaskService interface {
	Add(description string) (int, error)
	List(filter domain.TaskStatus) ([]domain.Task, error)
	Delete(taskID int) (bool, error)
	Update(taskID int, description string) (bool, error)
	UpdateStatus(taskID int, status domain.TaskStatus) (bool, error)
}

type CLI struct {
	tasks TaskService
	out   io.Writer
}

func New(tasks TaskService, out io.Writer) *CLI {
	return &CLI{
		tasks: tasks,
		out:   out,
	}
}

func (c *CLI) Run(args []string) error {
	if len(args) < 1 {
		return errors.New("no command provided")
	}

	method := args[0]
	switch method {
	case "add":
		if len(args) < 2 {
			return errors.New("task description is required")
		}
		id, err := c.tasks.Add(args[1])
		if err != nil {
			return err
		}
		fmt.Fprintf(c.out, "Task added successfully (ID: %d)", id)
	case "list":
		var filter domain.TaskStatus
		var err error
		if len(args) > 1 {
			filter, err = domain.NewTaskStatus(args[1])
			if err != nil {
				return err
			}
		}
		tasks, err := c.tasks.List(filter)
		if err != nil {
			return err
		}
		fmt.Fprintf(c.out, "Tasks: %d", len(tasks))
		for _, task := range tasks {
			fmt.Fprintf(c.out, "\n----------------------------------------\n")
			fmt.Fprintf(c.out, "Task ID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			fmt.Fprintf(c.out, "\n----------------------------------------\n")
		}
	case "delete":
		if len(args) < 2 {
			return errors.New("task ID is required for delete")
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		isDelete, err := c.tasks.Delete(taskID)
		if err != nil {
			return err
		}
		if isDelete {
			fmt.Fprintf(c.out, "Task with ID %d was deleted", taskID)
		} else {
			fmt.Fprintf(c.out, "Task with ID %d was not found", taskID)
		}
	case "update":
		if len(args) < 3 {
			return errors.New("task ID and description are required for update")
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.New("task ID must be a number")
		}
		newDescription := args[2]
		if newDescription == "" {
			return errors.New("task description must not be empty")
		}
		isUpdate, err := c.tasks.Update(taskID, newDescription)
		if err != nil {
			return err
		}
		if isUpdate {
			fmt.Fprintf(c.out, "Task with ID %d was updated", taskID)
		} else {
			fmt.Fprintf(c.out, "Task with ID %d was not found", taskID)
		}
	case "mark-in-progress":
		if len(args) < 2 {
			return errors.New("task ID is required to update status")
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.New("task ID must be a number")
		}
		isUpdate, err := c.tasks.UpdateStatus(taskID, domain.InProgress)
		if err != nil {
			return err
		}
		if isUpdate {
			fmt.Fprintf(c.out, "Task with ID %d status was updated", taskID)
		} else {
			fmt.Fprintf(c.out, "Task with ID %d was not found", taskID)
		}
	case "mark-done":
		if len(args) < 2 {
			return errors.New("task ID is required to update status")
		}
		taskID, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.New("task ID must be a number")
		}
		isUpdate, err := c.tasks.UpdateStatus(taskID, domain.Done)
		if err != nil {
			return err
		}
		if isUpdate {
			fmt.Fprintf(c.out, "Task with ID %d status was updated", taskID)
		} else {
			fmt.Fprintf(c.out, "Task with ID %d was not found", taskID)
		}
	default:
		return errors.New("unknown command")
	}
	return nil
}
