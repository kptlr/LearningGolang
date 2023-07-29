package main

import (
	"fmt"
	"runtime"
)

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	fmt.Println("Я сделяль")

	runtime.GC()
	_ = structure[0]
}
