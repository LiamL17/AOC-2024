package day11

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run() {
	buf, err := ioutil.ReadFile("days/day11/input.txt")
	check(err)

	bufferString := string(buf)
	stones := strings.Split(bufferString, " ")

	// fmt.Println("\nStones after", blinks, "blinks:", splited)

	fmt.Println("Number of stones Part 1 (25 blinks):", len(blink(0, 25, stones)))
	// fmt.Println("Number of stones Part 2 (75 blinks):", len(blink(0, 75, stones)))

}

func blink(blinks int, maxBlinks int, array []string) []string {
	// fmt.Println("Bink no.", blinks, "Array", array)
	if blinks == maxBlinks {
		return array
	}

	var newArray []string

	for _, stone := range array {

		if len(stone)%2 == 0 {
			// split into 2 stones
			left, right := split(stone)
			leftInt, _ := strconv.Atoi(left)
			rightInt, _ := strconv.Atoi(right)

			// array = append(array[:index], append([]string{left, right}, array[index+1:]...)...)
			newArray = append(newArray, strconv.Itoa(leftInt))
			newArray = append(newArray, strconv.Itoa(rightInt))

		} else if stone == "0" {
			// replace stone with '1'
			newArray = append(newArray, "1")
			// array[index] = "1"
		} else {
			// stone * 1024
			stoneInt, _ := strconv.Atoi(stone)
			// array[index] = strconv.Itoa(stoneInt * 2024)
			newArray = append(newArray, strconv.Itoa(stoneInt*2024))
		}

	}

	return blink(blinks+1, maxBlinks, newArray)
}

func split(stone string) (string, string) {
	len := len(stone)
	left := stone[:len/2]
	right := stone[len/2:]
	return left, right
}
