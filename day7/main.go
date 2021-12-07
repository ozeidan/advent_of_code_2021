package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	positions, maxPosition := readPositionsAndMax(scanner.Text())

	fmt.Println("solution of part1")
	part1(positions, maxPosition)
	fmt.Println("solution of part2")
	part2(positions, maxPosition)
}

func part1(positions []int, maxPosition int) {
	positionCounts := make([]int, maxPosition+1)

	for _, position := range positions {
		positionCounts[position]++
	}

	crabsRightOfI := make([]int, maxPosition+1)
	crabsLeftOfI := make([]int, maxPosition+1)

	cost := 0
	for i, acc := 0, 0; i <= maxPosition; i++ {
		crabsRightOfI[i] = acc
		acc += positionCounts[i]
	}

	for i, acc := maxPosition, 0; i >= 0; i-- {
		crabsLeftOfI[i] = acc
		cost += acc
		acc += positionCounts[i]
	}

	minCost := cost

	for i := 1; i <= maxPosition; i++ {
		cost -= crabsLeftOfI[i-1]
		cost += crabsRightOfI[i]

		if cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}

func part2(positions []int, maxPosition int) {
	minCost := math.MaxInt
	for i := 0; i <= maxPosition; i++ {
		cost := 0

		for _, position := range positions {
			dist := abs(i - position)
			cost += (dist * (dist + 1)) / 2
		}

		if cost < minCost {
			minCost = cost
		}
	}

	fmt.Println(minCost)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func readPositionsAndMax(input string) ([]int, int) {
	max := 0
	var positions []int
	for _, positionStr := range strings.Split(input, ",") {
		position, err := strconv.Atoi(positionStr)
		if err != nil {
			panic(err)
		}
		positions = append(positions, position)
		if position > max {
			max = position
		}
	}
	return positions, max
}
