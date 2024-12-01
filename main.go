package main

import (
	"aoc-2024/Days/Day1"
	"flag"
	"fmt"
)

func main() {
	var day int
	flag.IntVar(&day, "Day", 1, "Specify the day to execute")
	flag.Parse()

	dayFunctions := map[int]func(){
		1: Day1.Run,
	}

	if fn, exists := dayFunctions[day]; exists {
		fmt.Println("Executing day:", day)
		fmt.Println("----------------------------")
		fmt.Println("")
		fn()
	} else {
		fmt.Printf("Day %d is not implemented yet.\n", day)
	}
}
