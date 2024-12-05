package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {

	file, err := os.Open("days/day4/input.txt")
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
	count += countDiagonal(grid)

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

func partTwo(grid [][]string) int {
	count := 0
	regex := regexp.MustCompile("M.S.A.M.S|M.M.A.S.S|S.M.A.S.M|S.S.A.M.M")

	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[row])-2; col++ {
			str := ""

			for i := 0; i <= 2; i++ {
				for j := 0; j <= 2; j++ {
					str += grid[row+i][col+j]
				}
			}

			if regex.FindStringIndex(str) != nil {
				count++
			}
		}
	}

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

func countDiagonal(grid [][]string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Diagonal top-left to bottom-right
	for start := 0; start < rows+cols-1; start++ {
		diagonal := ""
		for i := 0; i < rows; i++ {
			j := start - i
			if j >= 0 && j < cols {
				diagonal += string(grid[i][j])
			}
		}
		count += countInString(diagonal)
	}

	// Diagonal top-right to bottom-left
	for start := 0; start < rows+cols-1; start++ {
		diagonal := ""
		for i := 0; i < rows; i++ {
			j := start - (rows - 1 - i)
			if j >= 0 && j < cols {
				diagonal += string(grid[i][j])
			}
		}
		count += countInString(diagonal)
	}

	fmt.Printf("Diagonal Count: %d\n", count)
	return count
}
