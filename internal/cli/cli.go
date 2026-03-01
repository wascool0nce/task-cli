package cli

import "errors"

func DispetcherComands(args []string) {
	method := parseMethod(args)
}

func parseMethod(args []string) (string, error) {
	if len(args) > 0 {
		method := args[0]
		return method, nil
	}
	return "", errors.New("Не получилось распарсить метод")
}
