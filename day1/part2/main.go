package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	window := make([]int, 0, 3)
	currentWindowIndex := 0
	result := 0

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if len(window) < 3 {
			window = append(window, measurement)
			continue
		}

		if measurement > window[currentWindowIndex] {
			result++
		}

		window[currentWindowIndex] = measurement
		currentWindowIndex = (currentWindowIndex + 1) % 3
	}

	fmt.Println(result)
}
