package day2

import (
	"bufio"
	"fmt"
	"math"
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
	lineCount := 0

	safeSequence := 0
	safeSequenceDampened := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		array := toIntArray(parts)

		if IsSafeSequence(array) {
			safeSequence++
		}

		if IsSafeSequenceDampened(array) {
			safeSequenceDampened++
		}

		lineCount++
	}

	fmt.Println("Total:", safeSequence)
	fmt.Println("Total Part Two:", safeSequenceDampened)
}

func RemoveIdx(s []int, i int) []int {
	r := make([]int, 0)
	r = append(r, s[:i]...)
	return append(r, s[i+1:]...)
}

func IsSafeSequence(nums []int) bool {
	cmp := func(a, b int) bool { return a < b }
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] == 0 {
			continue
		}
		if nums[i]-nums[i+1] > 0 {
			cmp = func(a, b int) bool { return a > b }
		}
		break
	}

	for i := 1; i < len(nums); i++ {
		diff := math.Abs(float64(nums[i] - nums[i-1]))

		if diff < 1 || diff > 3 || !cmp(nums[i-1], nums[i]) {
			return false
		}
	}

	return true
}

func IsSafeSequenceDampened(nums []int) bool {
	if IsSafeSequence(nums) {
		return true
	}
	for i, _ := range nums {
		if IsSafeSequence(RemoveIdx(nums, i)) {
			return true
		}
	}
	return false
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
