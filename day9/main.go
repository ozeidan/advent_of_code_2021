package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct {
	x, y int
}

func main() {
	file, _ := os.Open("./input")
	scanner := bufio.NewScanner(file)

	matrix := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		nums := make([]int, 0)

		for _, rune := range line {
			num := int(rune - '0')
			nums = append(nums, num)
		}

		matrix = append(matrix, nums)
	}

	lowpoints := make([]point, 0)
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			num := matrix[i][j]

			if i+1 < len(matrix) && num >= matrix[i+1][j] {
				continue
			}
			if i-1 >= 0 && num >= matrix[i-1][j] {
				continue
			}
			if j+1 < len(matrix[i]) && num >= matrix[i][j+1] {
				continue
			}
			if j-1 >= 0 && num >= matrix[i][j-1] {
				continue
			}
			lowpoints = append(lowpoints, point{i, j})
			res += num + 1
		}
	}

	scores := make([]int, 0)

	for _, lowpoint := range lowpoints {
		score := measureBasin(matrix, lowpoint, make(map[point]struct{}))
		scores = append(scores, score)
	}

	fmt.Println(res)
	sort.Ints(scores)
	fmt.Println(scores[len(scores)-3:])
	fmt.Println(scores[len(scores)-1] * scores[len(scores)-2] * scores[len(scores)-3])
}

func measureBasin(matrix [][]int, currentPoint point, visited map[point]struct{}) int {
	i := currentPoint.x
	j := currentPoint.y
	num := matrix[i][j]

	if _, ok := visited[currentPoint]; ok || num == 9 {
		return 0
	}
	score := 1

	visited[currentPoint] = struct{}{}
	if i+1 < len(matrix) && num <= matrix[i+1][j] {
		score += measureBasin(matrix, point{i + 1, j}, visited)
	}
	if i-1 >= 0 && num <= matrix[i-1][j] {
		score += measureBasin(matrix, point{i - 1, j}, visited)
	}
	if j+1 < len(matrix[i]) && num <= matrix[i][j+1] {
		score += measureBasin(matrix, point{i, j + 1}, visited)
	}
	if j-1 >= 0 && num <= matrix[i][j-1] {
		score += measureBasin(matrix, point{i, j - 1}, visited)
	}
	return score
}

//nolint
func printMatrix(matrix [][]int, visited map[point]struct{}) {
	for i, row := range matrix {
		for j, val := range row {
			if _, ok := visited[point{i, j}]; ok {
				fmt.Printf("X")
			} else {
				fmt.Printf("%d", val)
			}
		}
		fmt.Println()
	}
}
