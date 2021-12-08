package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type inputLine struct {
	signalPatterns []string
	digitPatterns  []string
}

type segments map[rune]struct{}

var allSegments segments = map[rune]struct{}{
	'a': {},
	'b': {},
	'c': {},
	'd': {},
	'e': {},
	'f': {},
	'g': {},
}

type mappingPossibilities map[rune]segments

var segmentsOfDigits = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

var digitOfSegments = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

var digitsOfLengths = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	7: {8},
	6: {0, 6, 9},
	5: {2, 3, 5},
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var inputLines []inputLine
	for scanner.Scan() {
		inputLines = append(inputLines, parseInputLine(scanner.Text()))
	}
	part1(inputLines)
	part2(inputLines)
}

func part1(inputLines []inputLine) {
	result := 0
	for _, inputLine := range inputLines {
		for _, digit := range inputLine.digitPatterns {
			if _, ok := digitsOfLengths[len(digit)]; ok {
				result++
			}
		}
	}
	fmt.Println(result)
}

func part2(inputLines []inputLine) {
	result := 0
	for _, line := range inputLines {
		mappings := resolveMappingsFromSignalPatterns(line.signalPatterns)
		reverseMappings := make(map[rune]rune)
		for k, v := range mappings {
			reverseMappings[v] = k
		}

		digitStrings := ""

		for _, digitPattern := range line.digitPatterns {
			actualSegments := []rune{}
			digitSegments := signalPatternToSegments(digitPattern)
			for digitSegment := range digitSegments {
				actualSegments = append(actualSegments, reverseMappings[digitSegment])
			}
			sort.Slice(actualSegments, func(i, j int) bool {
				return actualSegments[i] < actualSegments[j]
			})

			if _, ok := digitOfSegments[string(actualSegments)]; !ok {
				panic("found invalid pattern " + string(actualSegments))
			}

			actualDigit := fmt.Sprint(digitOfSegments[string(actualSegments)])
			digitStrings += actualDigit
		}

		number, err := strconv.Atoi(digitStrings)
		if err != nil {
			panic(err)
		}
		result += number
	}

	fmt.Println(result)
}

func resolveMappingsFromSignalPatterns(signalPatterns []string) map[rune]rune {
	mappingPossibilities := initialisePossibilityMap()
	mappings := make(map[rune]rune)

	for len(mappings) < 7 {
		for _, signalPattern := range signalPatterns {
			if digit, ok := tryDeductDigit(mappingPossibilities, signalPatternToSegments(signalPattern)); ok {
				actualSegments := signalPatternToSegments(segmentsOfDigits[digit])
				segmentsInPattern := signalPatternToSegments(signalPattern)

				for segment := range mappingPossibilities {
					if _, ok := actualSegments[segment]; ok {
						mappingPossibilities[segment] = intersection(
							mappingPossibilities[segment],
							segmentsInPattern,
						)
					} else {
						mappingPossibilities[segment] = diff(
							mappingPossibilities[segment],
							segmentsInPattern,
						)
					}
					if len(mappingPossibilities[segment]) == 1 {
						mappings[segment] = segmentsToSlice(mappingPossibilities[segment])[0]
					}
				}
			}

		}
	}

	return mappings
}

func tryDeductDigit(possibilities mappingPossibilities, inputSegments segments) (int, bool) {
	if len(digitsOfLengths[len(inputSegments)]) == 1 {
		return digitsOfLengths[len(inputSegments)][0], true
	}

	newPoss := make(mappingPossibilities)
	for segment, possibleSegments := range possibilities {
		newPoss[segment] = intersection(possibleSegments, inputSegments)
	}

	var possibleDigits []int
	for _, digit := range digitsOfLengths[len(inputSegments)] {
		if digitIsPossible(newPoss, []rune(segmentsOfDigits[digit]), make(segments)) {
			possibleDigits = append(possibleDigits, digit)
		}
	}

	if len(possibleDigits) == 1 {
		return possibleDigits[0], true
	}

	return 0, false
}

func digitIsPossible(possibilities mappingPossibilities, digitSegments []rune, takenSegments segments) bool {
	if len(digitSegments) == 0 {
		return true
	}

	digitSegment := digitSegments[0]
	poss := diff(possibilities[digitSegment], takenSegments)

	for p := range poss {
		newTakenSegments := make(segments)
		for segment := range takenSegments {
			newTakenSegments[segment] = struct{}{}
		}
		newTakenSegments[p] = struct{}{}
		if digitIsPossible(possibilities, digitSegments[1:], newTakenSegments) {
			return true
		}
	}

	return false
}

func initialisePossibilityMap() mappingPossibilities {
	possibilites := make(mappingPossibilities)
	for segment := range allSegments {
		possibilites[segment] = allSegments
	}
	return possibilites
}

//nolint
func printPossibilityMap(possibilities mappingPossibilities) {
	for k, v := range possibilities {
		fmt.Printf("%s: ", string(k))
		for kv := range v {
			fmt.Printf("%s, ", string(kv))
		}
		fmt.Println()
	}
}

// golang generics are coming soon!!!!!!!!!!!!!!!!!!
func signalPatternToSegments(signalPattern string) segments {
	segments := make(segments)

	for _, segment := range signalPattern {
		segments[segment] = struct{}{}
	}
	return segments
}

func segmentsToSlice(segments segments) []rune {
	var res []rune
	for segment := range segments {
		res = append(res, segment)
	}
	return res
}

func intersection(set1, set2 segments) segments {
	res := make(segments)
	for k := range set1 {
		if _, ok := set2[k]; ok {
			res[k] = struct{}{}
		}
	}
	return res
}

func diff(set1, set2 segments) segments {
	res := make(segments)
	for k := range set1 {
		if _, ok := set2[k]; ok {
			continue
		}
		res[k] = struct{}{}
	}
	return res
}

func parseInputLine(line string) inputLine {
	halves := strings.Split(line, " | ")
	return inputLine{strings.Split(halves[0], " "), strings.Split(halves[1], " ")}
}
