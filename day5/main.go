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
	scanner := bufio.NewScanner(file)

	coordinateScoresNoDiags := make(map[point]int)
	coordinateScores := make(map[point]int)
	for scanner.Scan() {
		line := parseLine(scanner.Text())

		line.coordinates(func(x, y int) {
			coordinateScores[point{x, y}]++

			if !line.isDiagonal() {
				coordinateScoresNoDiags[point{x, y}]++
			}
		})
	}

	part1Result := 0
	for _, score := range coordinateScoresNoDiags {
		if score >= 2 {
			part1Result++
		}
	}

	part2Result := 0
	for _, score := range coordinateScores {
		if score >= 2 {
			part2Result++
		}
	}

	fmt.Println("result for part1:")
	fmt.Println(part1Result)
	fmt.Println("result for part2:")
	fmt.Println(part2Result)
}

type point struct {
	x, y int
}

type line struct {
	x1, y1 int
	x2, y2 int
}

func (l line) isDiagonal() bool {
	return l.x1 != l.x2 && l.y1 != l.y2
}

func (l line) coordinates(callback func(x, y int)) {
	deltaX := 0
	deltaY := 0

	if l.x1 < l.x2 {
		deltaX = 1
	} else if l.x1 > l.x2 {
		deltaX = -1
	}

	if l.y1 < l.y2 {
		deltaY = 1
	} else if l.y1 > l.y2 {
		deltaY = -1
	}

	for x, y := l.x1, l.y1; x != l.x2 || y != l.y2; x, y = x+deltaX, y+deltaY {
		callback(x, y)
	}

	callback(l.x2, l.y2)
}

func parseLine(inputLine string) line {
	parts := strings.Split(inputLine, " ")
	coordinateStrings := append(strings.Split(parts[0], ","), strings.Split(parts[2], ",")...)
	var coordinates []int
	for _, coordinateString := range coordinateStrings {
		coordinate, err := strconv.Atoi(coordinateString)
		if err != nil {
			panic(err)
		}
		coordinates = append(coordinates, coordinate)
	}
	return line{
		x1: coordinates[0],
		y1: coordinates[1],
		x2: coordinates[2],
		y2: coordinates[3],
	}
}
