package cli

import (
	"errors"
	"fmt"
	"io"
	"task-cli/internal/usecase"
)

type CLI struct {
	tasks *usecase.TaskService
	out   io.Writer
}

func New(tasks *usecase.TaskService, out io.Writer) *CLI {
	return &CLI{
		tasks: tasks,
		out:   out,
	}
}

func (c *CLI) Run(args []string) error {
	if len(args) < 1 {
		return errors.New("нет аргументов")
	}

	method := args[0]
	switch method {
	case "add":
		if len(args) < 2 {
			return errors.New("нет описания задачи")
		}
		err := c.tasks.Add(args[1])
		if err != nil {
			return err
		}
		fmt.Fprintf("Task added successfully (ID: %d)")
	default:
		return errors.New("такого метода не существует")
	}
	return nil
}
