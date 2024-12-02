package day2

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
	file, err := os.Open("days/day2/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		array := toIntArray(parts)

		safe := false
		if array[0] < array[1] {
			safe = checkIncreasing(array)
		} else if array[0] > array[1] {
			safe = checkDecreasing(array)
		}

		if safe {
			total++
		}
	}

	fmt.Println("Total:", total)
}

func toIntArray(parts []string) []int {
	var intArray []int
	for _, element := range parts {
		num, err := strconv.Atoi(element)
		check(err)
		intArray = append(intArray, num)
	}

	return intArray
}

func checkIncreasing(array []int) bool {
	previous := array[0]
	for i := 1; i < len(array); i++ {
		diff := array[i] - previous
		if diff <= 0 || diff >= 4 {
			return false
		}

		previous = array[i]
	}
	return true
}

func checkDecreasing(array []int) bool {
	previous := array[0]
	for i := 1; i < len(array); i++ {
		diff := previous - array[i]
		if diff <= 0 || diff >= 4 {
			return false
		}

		previous = array[i]
	}
	return true
}
