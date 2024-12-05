package day5

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
	file, err := os.Open("days/day5/test.txt")
	check(err)
	defer file.Close()

	// rules := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 5 {
			fmt.Println(line)
			num1, _ := strconv.Atoi(strings.Split(line, "")[0:1])
			num2, _ := strconv.Atoi(strings.Split(line, "")[3:4])
			fmt.Println(num1, ":", num2)
		} else {

		}
	}
}
