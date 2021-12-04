package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./example")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	previousMeasurement := 0
	result := 0
	resultIsInitialised := false

	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if measurement > previousMeasurement && resultIsInitialised {
			result++
		}

		previousMeasurement = measurement
		resultIsInitialised = true
	}

	fmt.Println(result)
}
