package cli

import (
	"errors"
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
		err := c.tasks.Add(args[1])
		if err != nil {
			return err
		}
	default:
		return errors.New("такого метода не существует")
	}
	return nil
}
