package day12

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

	file, err := os.Open("days/day12/test.txt")
	check(err)
	defer file.Close()

	var grid [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	neighborsGrid := calculateNeighbors(grid)

	fmt.Println("Neighbors grid:", neighborsGrid)
}

func calculateNeighbors(grid [][]string) [][]int {
	var ans [][]int

	for i := 0; i < len(grid); i++ {
		var arr []int
		for j := 0; j < len(grid[i]); j++ {
			count := 0
			// up
			if i-1 < 0 || grid[i-1][j] != grid[i][j] {
				count++
			}
			// down
			if i+1 >= len(grid) || grid[i+1][j] != grid[i][j] {
				count++
			}
			// left
			if j-1 < 0 || grid[i][j-1] != grid[i][j] {
				count++
			}
			// right
			if j+1 >= len(grid[i]) || grid[i][j+1] != grid[i][j] {
				count++
			}
            arr = append(arr, count)
		}
		ans = append(ans, arr)
	}

	return ans
}
