package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	count := 90

	if len(os.Args) > 1 {
		arg := os.Args[1]
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("invalid argument:", arg)
			os.Exit(1)
		}
		count = n
	}

	resp, err := Collector(count)
	if err != nil {
		os.Exit(1)
	}
	data, err := Converter(resp)
	if err != nil {
		os.Exit(1)
	}
	result := Agregation(data)
	Min(result)
	Max(result)
	AVG(result)
}
