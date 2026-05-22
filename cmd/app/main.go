package main

import (
	"fmt"
	"os"
	"task-cli/internal/cli"
	"task-cli/internal/storage/jsonstore"
	"task-cli/internal/usecase"
	"time"
)

func main() {
	path := "tasks.json"
	store := jsonstore.New(path)
	service := usecase.NewTaskService(store, time.Now)
	app := cli.New(service, os.Stdout)
	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
