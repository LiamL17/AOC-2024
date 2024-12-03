package day3

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	buf, err := ioutil.ReadFile("days/day3/input.txt")
	check(err)

	// Must match mul(X,Y) where X & Y is any 3 digit number
	regex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := regex.FindAllString(string(buf), -1)
	total := 0

	for _, element := range matches {
		regex, _ := regexp.Compile(`\d{1,3}`)
		nums := regex.FindAllString(element, -1)

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		multiplied := num1 * num2

		// fmt.Printf("%s = %d * %d = %d\n", element, num1, num2, multiplied)
		total += multiplied
	}

	fmt.Println("Total:", total)

}
