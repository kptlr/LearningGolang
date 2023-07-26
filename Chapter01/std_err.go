package main

import (
	"io"
	"os"
)

func main() {
	str := ""
	arguments := os.Args
	if len(arguments) == 1 {
		str = "Пожалуйста, мне нужен один аргумент!"
	} else {
		str = arguments[1]
	}

	io.WriteString(os.Stdout, "Это - стандартный вывод\n")
	io.WriteString(os.Stderr, str)
	io.WriteString(os.Stderr, "\n")
}
