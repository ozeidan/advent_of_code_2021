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
	populationAfter80 := simulateFishLife(amountOfFishWithDays, 80)
	populationAfter256 := simulateFishLife(amountOfFishWithDays, 256-80)

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

func simulateFishLife(amountOfFishWithDays []int, simulationLengthDays int) int {
	for i := 0; i < simulationLengthDays; i++ {
		step(amountOfFishWithDays)
	}

	sum := 0

	for _, amount := range amountOfFishWithDays {
		sum += amount
	}

	return sum
}

func step(amountOfFish []int) {
	reproducingFish := amountOfFish[0]

	for i := 0; i < 8; i++ {
		amountOfFish[i] = amountOfFish[i+1]
	}
	amountOfFish[8] = reproducingFish
	amountOfFish[6] += reproducingFish
}
