package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, val := range os.Args[1:] {
		s += sep + val
		sep = " "
	}
	fmt.Println(s)
}
