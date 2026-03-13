package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args, err := validateArgs(os.Args[1:])
	if err != nil {
		log.Fatalf("Ошибка запуска taks-cli: %v", err)
	}

	_, err := DispetcherComands(args)
}

func validateArgs(args []string) ([]string, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("не хватает аргументов командной строки")
	}
	return args, nil
}
