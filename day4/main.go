package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bingoBoard struct {
	numbers [][]int
	marked  [][]bool
}

func newBingoBoard(numbers [][]int) bingoBoard {
	marked := make([][]bool, 5)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			marked[i] = append(marked[i], false)
		}
	}

	return bingoBoard{
		numbers: numbers,
		marked:  marked,
	}
}

func (b bingoBoard) print() {
	for i, row := range b.numbers {
		for j, number := range row {
			if !b.marked[i][j] {
				fmt.Printf("%-3d", number)
			} else {
				fmt.Printf("\033[1m%-3d\033[0m", number)
			}
		}
		fmt.Println()
	}
}

func (b bingoBoard) isDone() bool {
	for i := 0; i < 5; i++ {
		rowAllMarked := true
		colAllMarked := true

		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				rowAllMarked = false
			}
			if !b.marked[j][i] {
				colAllMarked = false
			}
		}

		if rowAllMarked || colAllMarked {
			return true
		}
	}
	return false
}

func (b bingoBoard) score(finallyMarkedNumber int) int {
	sum := 0
	for i, rows := range b.numbers {
		for j, number := range rows {
			if !b.marked[i][j] {
				sum += number
			}
		}
	}
	return sum * finallyMarkedNumber
}

func (b *bingoBoard) mark(number int) {
	for i, rows := range b.numbers {
		for j, existingNumber := range rows {
			if existingNumber == number {
				b.marked[i][j] = true
				return
			}
		}
	}
}

func lineToNumbers(line, sep string) []int {
	numberStrings := strings.Split(line, sep)
	var numbers []int
	for _, numberString := range numberStrings {
		if len(strings.TrimSpace(numberString)) == 0 {
			continue
		}
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbers := lineToNumbers(scanner.Text(), ",")
	scanner.Scan()

	var boards []bingoBoard

	for {
		var bingoRows [][]int
		for i := 1; i <= 5; i++ {
			scanner.Scan()
			newRow := lineToNumbers(scanner.Text(), " ")
			bingoRows = append(bingoRows, newRow)
		}

		boards = append(boards, newBingoBoard(bingoRows))
		if !scanner.Scan() {
			break
		}
	}

	winningBoard, winningNumber := boardWinsFirst(boards, numbers)
	fmt.Println("winning board:")
	winningBoard.print()
	fmt.Println("score:")
	fmt.Println(winningBoard.score(winningNumber))

	fmt.Println()

	losingBoard, losingNumber := boardWinsLast(boards, numbers)
	fmt.Println("last winning board:")
	winningBoard.print()
	fmt.Println("score:")
	fmt.Println(losingBoard.score(losingNumber))
}

func boardWinsFirst(boards []bingoBoard, numbers []int) (bingoBoard, int) {
	for _, number := range numbers {
		for _, board := range boards {
			board.mark(number)
			if board.isDone() {
				return board, number
			}
		}
	}

	panic("no board won")
}

func boardWinsLast(boards []bingoBoard, numbers []int) (bingoBoard, int) {
	var lastWinningBoard bingoBoard
	var lastWinningNumber int

	for _, number := range numbers {
		if len(boards) == 0 {
			break
		}

		var boardsOfNextRound []bingoBoard
		for _, board := range boards {
			board.mark(number)
			if board.isDone() {
				lastWinningBoard, lastWinningNumber = board, number
			} else {
				boardsOfNextRound = append(boardsOfNextRound, board)
			}
		}
		boards = boardsOfNextRound
	}

	return lastWinningBoard, lastWinningNumber
}
