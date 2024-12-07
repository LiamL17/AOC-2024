package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {

	file, err := os.Open("days/day6/input.txt")
	check(err)
	defer file.Close()

	var grid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	// fmt.Println("Grid:", grid)

	startRow, startCol := getStart(grid)
	// fmt.Println("Start: (", startRow, ",", startCol, ")")

	count := countGuardSteps(grid, startRow, startCol)

	fmt.Println("Total part 1:", count)
	// fmt.Println("Total part 2:", total2)

}

func countGuardSteps(grid [][]string, row int, col int) int {
	direction := Up
	count := 0
	for row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) {
		// for _, line := range grid {
		// 	fmt.Println(line)
		// }
		// fmt.Println("\n")

		if grid[row][col] == "#" {
			switch direction {
			case Up:
				col++
				row++
				direction = Right
			case Down:
				col--
				row--
				direction = Left
			case Left:
				row--
				col++
				direction = Up
			case Right:
				row++
				col--
				direction = Down
			}
		} else {
			if grid[row][col] != "X" {
				count++
				grid[row][col] = "X"
			}

			switch direction {
			case Up:
				row--
			case Down:
				row++
			case Left:
				col--
			case Right:
				col++
			}
		}
	}

	return count
}

func getStart(grid [][]string) (int, int) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "^" {
				return i, j
			}
		}
	}

	return -1, -1
}
