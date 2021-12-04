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

	numbers, bitCount := readLines(file)

	fmt.Println("output of part 1:")
	part1(numbers, bitCount)

	fmt.Println("output of part 2:")
	part2(numbers, bitCount)
}

func part1(numbers []int, bitCount int) {
	numberCount := len(numbers)
	oneCounts := countOnes(numbers, bitCount)

	var gamma uint = 0

	for i, count := range oneCounts {
		if count > numberCount/2 {
			gamma |= (1 << i)
		}
	}

	var epsilon uint = ^gamma & ((1 << bitCount) - 1)
	fmt.Printf("gamma: %012b\n", gamma)
	fmt.Printf("epsilon: %012b\n", epsilon)
	fmt.Printf("result: %d\n", gamma*epsilon)
}

func part2(numbers []int, bitCount int) {
	keepMostCommon := func(numberCount, oneCount int) bool {
		return oneCount >= (numberCount+numberCount%2)/2
	}
	keepLeastCommon := negate(keepMostCommon)

	oxyRating := findNumber(numbers, bitCount, keepMostCommon)
	co2Rating := findNumber(numbers, bitCount, keepLeastCommon)

	fmt.Printf("oxygen rating: %d\n", oxyRating)
	fmt.Printf("co2 rating: %d\n", co2Rating)
	fmt.Printf("life support rating: %d\n", oxyRating*co2Rating)
}

func negate(f func(int, int) bool) func(int, int) bool {
	return func(n, m int) bool {
		return !f(n, m)
	}
}

func findNumber(numbers []int, bitCount int, shouldKeepOne func(numberCount, ones int) bool) int {
	for i := bitCount - 1; len(numbers) > 1 && i >= 0; i-- {
		newNumbers := make([]int, 0, len(numbers))
		ones := countOnesAtIndex(numbers, i)

		for _, number := range numbers {
			if (number & (1 << i)) > 0 == shouldKeepOne(len(numbers), ones) {
				newNumbers = append(newNumbers, number)
			}
		}
		numbers = newNumbers
	}
	return numbers[0]
}

//nolint
func printNumbersBinary(numbers []int) {
	fmt.Printf("[")
	for i := 0; i < len(numbers)-1; i++ {
		fmt.Printf("%05b, ", numbers[i])
	}
	fmt.Printf("%05b]\n", numbers[len(numbers)-1])
}

func countOnes(numbers []int, bitCount int) []int {
	ones := make([]int, 0, bitCount)

	for i := 0; i < bitCount; i++ {
		ones = append(ones, countOnesAtIndex(numbers, i))
	}

	return ones
}

func countOnesAtIndex(numbers []int, bitIndex int) int {
	ones := 0

	for _, number := range numbers {
		if number&(1<<bitIndex) > 0 {
			ones++
		}
	}

	return ones
}

func readLines(file *os.File) ([]int, int) {
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	bitCount := 0

	for scanner.Scan() {
		if bitCount == 0 {
			bitCount = len(scanner.Text())
		}

		parsed, err := strconv.ParseInt(scanner.Text(), 2, 0)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, int(parsed))
	}
	return numbers, bitCount
}
