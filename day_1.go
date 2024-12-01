package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("Day-1\\input1.txt")
	check(err)
	defer file.Close()

	var array1, array2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 4 {
			log.Fatalf("Not 2 items!")
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[3])
		check(err1)
		check(err2)

		array1 = append(array1, num1)
		array2 = append(array2, num2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner problem!")
	}

	slices.Sort(array1)
	slices.Sort(array2)

	total := 0
	for index, element := range array1 {
		diff := array2[index] - element
		if (diff < 0) {
			diff = diff * -1
		}
		total += diff
	}

	fmt.Println("Total:", total)
}
