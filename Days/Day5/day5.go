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
	file, err := os.Open("days/day5/input.txt")
	check(err)
	defer file.Close()

	rules := make(map[int][]int)
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 5 {
			num1, _ := strconv.Atoi(strings.Join(strings.Split(line, "")[0:2], ""))
			num2, _ := strconv.Atoi(strings.Join(strings.Split(line, "")[3:5], ""))
			rules[num1] = append(rules[num1], num2)
		} else if line != "" {
			var update []int
			split := strings.Split(line, ",")
			for _, element := range split {
				converted, _ := strconv.Atoi(element)
				update = append(update, converted)
			}
			updates = append(updates, update)
		}
	}

	fmt.Println(updates)
	fmt.Println(rules)

	total := 0
	total2 := 0
	for _, update := range updates {
		valid := validUpdate(update, rules)
		// fmt.Println("Update:", i + 1, "Valid:", valid, "\n")
		if valid {
			total += update[len(update)/2]
		} else {
			// newUpdate := invalidFixed(update, rules)
			// fmt.Println("New Update:", i+1, "Valid:", newUpdate, "\n")
			total2 += invalidFixed(update, rules)
		}
	}

	fmt.Println("Total part 1:", total)
	fmt.Println("Total part 2:", total2)
}

func invalidFixed(update []int, rules map[int][]int) int {
    // fmt.Println("Update fixing:", update)
	if validUpdate(update, rules) {
		return update[len(update)/2]
	}
	len := len(update)

	// fmt.Println("Update:", update)
	for i := 0; i < len-1; i++ {
		current := update[i]
		rule := rules[current]

		// fmt.Println("Current:", current, ". Rule:", rule)
		for j := i + 1; j < len; j++ {
			if !contains(rule, update[j]) {
				update[i], update[j] = update[j], update[i]
			}
		}

	}
	// return newUpdate
    return invalidFixed(update, rules)
}

func validUpdate(update []int, rules map[int][]int) bool {

	len := len(update)
	for i := 0; i < len-1; i++ {
		current := update[i]
		rule := rules[current]

		// fmt.Println("Current:", current, ". Rule:", rule)
		for j := i + 1; j < len; j++ {
			if !contains(rule, update[j]) {
				return false
			}
		}
	}

	return true
}

func contains(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
