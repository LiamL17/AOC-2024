package day9

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
	buf, err := ioutil.ReadFile("days/day9/input.txt")
	check(err)

	bufferString := string(buf)

	total := partOne(bufferString)

	fmt.Println("Total part 1:", total)
}

func partOne(input string) int {
	total := 0

	arranged := arrangeDiskMap(input)
	// fmt.Println("Arranged disk map:", arranged)
	moved := moveBlocks(arranged)
	// fmt.Println("Moved disk map:", moved)

	id := 0
	for i := 0; i < len(moved); i++ {
		if moved[i] == 'X' {
			continue
		}
		value, _ := strconv.Atoi(string(moved[i]))
		total += id * value
		id++
	}

	return total
}

func moveBlocks(arranged string) string {
	for i := len(arranged) - 1; i > 0 && strings.Contains(arranged, "."); i-- {
		indexFirstOpening := strings.Index(arranged, ".")
		arranged = replaceCharAtIndex(arranged, indexFirstOpening, rune(arranged[i]))

		arranged = replaceCharAtIndex(arranged, i, 'X')
	}

	return arranged
}

func replaceCharAtIndex(str string, index int, newChar rune) string {
	runes := []rune(str)

	if index < 0 || index >= len(runes) {
		fmt.Println("Index out of bounds")
		return str
	}

	runes[index] = newChar

	return string(runes)
}

func arrangeDiskMap(diskMap string) string {
	id := 0
	arranged := ""
	free := false
	for _, diskMap := range diskMap {
		intValue, _ := strconv.Atoi(string(diskMap))
		if free {
			for i := 0; i < intValue; i++ {
				arranged += "."
			}
		} else {
			for i := 0; i < intValue; i++ {
				arranged += strconv.Itoa(id)
			}
			id++
		}

		free = !free
	}
	return arranged
}
