package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontalPosition := 0
	depth := 0

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		length, err := strconv.Atoi(parts[1])

		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontalPosition += length
		case "down":
			depth += length
		case "up":
			depth -= length
		}
	}

	fmt.Println(horizontalPosition * depth)
}
