package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {

	file, err := os.Open("days/day7/input.txt")
	check(err)
	defer file.Close()

	var values []int
	var terms [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		value, _ := strconv.Atoi(line[0:strings.Index(line, ":")])
		values = append(values, value)
		items := strings.Split(line[strings.Index(line, ":")+2:], " ")

		var lineTemrs []int
		for _, item := range items {
			term, _ := strconv.Atoi(item)
			lineTemrs = append(lineTemrs, term)
		}

		terms = append(terms, lineTemrs)
	}

	total := calculateTotalPartOne(values, terms)

	fmt.Println("Part 1 total:", total)
}

func calculateTotalPartOne(values []int, terms [][]int) int {
	total := 0

	for i, value := range values {
		terms := terms[i]

		if isValidOperation(value, terms) {
			total += value
		}
	}

	return total
}

func isValidOperation(value int, terms []int) bool {
	results := calculateLeftToRight(terms)
	// fmt.Println("Terms:", terms)
	// fmt.Println("Results:", results)
	for _, result := range results {
		if value == result {
			return true
		}
	}

	return false
}

func calculateLeftToRight(numbers []int) []int {
	var results []int
	if len(numbers) == 0 {
		return []int{}
	}

	var helper func(index, current int)
	helper = func(index, current int) {
		if index == len(numbers) {
			results = append(results, current)
			return
		}

		helper(index+1, current+numbers[index])
		helper(index+1, current*numbers[index])
        
	}

	helper(1, numbers[0])
	return results
}
