package Day1

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

func Run() {
	file, err := os.Open("days/day1/input1.txt")
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

	occurs := getOccurences(array2)

	total := 0
	total2 := 0
	for index, element := range array1 {
		diff := array2[index] - element
		if diff < 0 {
			diff = diff * -1
		}
		total += diff

		total2 += element * occurs[element]
	}

	fmt.Println("Total:", total)
	fmt.Println("Total Part 2:", total2)
}

func getOccurences(array []int) map[int]int {
	occurences := make(map[int]int)
	for _, element := range array {
		occurences[element]++
	}

	return occurences
}
