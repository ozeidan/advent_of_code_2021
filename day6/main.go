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
	scanner.Scan()
	timerLine := scanner.Text()
	amountOfFishWithDays := initialiseCounts(timerLine)

	idx := 0
	populationAfter80, idx := simulateFishLife(amountOfFishWithDays, 80, idx)
	populationAfter256, _ := simulateFishLife(amountOfFishWithDays, 256-80, idx)

	fmt.Printf("population count after:\n\t80 days: %d\n\t256 days: %d\n",
		populationAfter80,
		populationAfter256)
}

func initialiseCounts(initialTimers string) []int {
	amountOfFishWithDays := make([]int, 9)

	for _, timerStr := range strings.Split(initialTimers, ",") {
		timer, err := strconv.Atoi(timerStr)
		if err != nil {
			panic(err)
		}
		amountOfFishWithDays[timer]++
	}
	return amountOfFishWithDays
}

func simulateFishLife(amountOfFishWithDays []int, simulationLengthDays, idx int) (int, int) {
	for i := 0; i < simulationLengthDays; i++ {
		idx = step(amountOfFishWithDays, idx)
	}

	sum := 0

	for _, amount := range amountOfFishWithDays {
		sum += amount
	}

	return sum, idx
}

func step(amountOfFish []int, currentIndex int) int {
	reproducingFish := amountOfFish[calcIndex(currentIndex, 0)]

	currentIndex = calcIndex(currentIndex, 1)
	amountOfFish[calcIndex(currentIndex, 8)] = reproducingFish
	amountOfFish[calcIndex(currentIndex, 6)] += reproducingFish

	return currentIndex
}

func calcIndex(index, offset int) int {
	return (index + offset) % 9
}
