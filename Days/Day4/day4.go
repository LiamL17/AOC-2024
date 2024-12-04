package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	file, err := os.Open("days/day4/test.txt")
	check(err)
	defer file.Close()

	var grid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	total := partOne(grid)
	total2 := partTwo(grid)

	fmt.Println("Total part 1:", total)
	fmt.Println("Total part 2:", total2)
}

func partOne(grid [][]string) int {
	count := 0

	count += countHorizontal(grid)
	count += countVertical(grid)
	count += countDiagonalForward(grid)
	count += countDiagonalBackward(grid)

	return count
}

func countHorizontal(grid [][]string) int {
	count := 0

	for _, element := range grid {
		count += countInString(strings.Join(element, ""))
	}

	fmt.Printf("Horizontal Count: %d\n", count)
	return count
}

func countVertical(grid [][]string) int {
	count := 0

	for col := 0; col < len(grid[0]); col++ {
		str := ""
		for row := 0; row < len(grid); row++ {
			str += grid[row][col]
		}
		count += countInString(str)
	}

	fmt.Printf("Vertical Count: %d\n", count)
	return count
}

func countDiagonalForward(grid [][]string) int {
	count := 0

	for row := 0; row < len(grid)-3; row++ {
		for col := 0; col < len(grid[row])-3; col++ {
			str := ""
			for diag := 0; diag < 4; diag++ {
				str += grid[row+diag][col+diag]
			}
            // fmt.Println("forward: ", str)
			if str == "XMAS" || str == "SAMX" {
				count++
			}
		}
	}

	fmt.Printf("Diagonal forward Count: %d\n", count)
	return count
}

func countDiagonalBackward(grid [][]string) int {
	count := 0
	for row := len(grid) - 1; row >= 3; row-- {
		for col := len(grid[row]) - 1; col >= 3; col-- {
			str := ""
			for diag := 0; diag < 4; diag++ {
				str += grid[row-diag][col-diag]
			}
            // fmt.Println("backward: ", str)
			if str == "XMAS" || str == "SAMX" {
				count++
			}
		}
	}

	fmt.Printf("Diagonal backward Count: %d\n", count)
	return count
}

func partTwo(grid [][]string) int {
	count := 0

	return count
}

func countInString(str string) int {
	count := 0

	for i := 0; i < len(str)-3; i++ {
		substring := str[i : i+4]
		if substring == "XMAS" || substring == "SAMX" {
			count++
		}
	}

	return count
}
