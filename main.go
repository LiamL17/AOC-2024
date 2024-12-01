package main

import (
	"aoc-2024/days/day1"
	"aoc-2024/days/day10"
	"aoc-2024/days/day11"
	"aoc-2024/days/day12"
	"aoc-2024/days/day13"
	"aoc-2024/days/day14"
	"aoc-2024/days/day15"
	"aoc-2024/days/day16"
	"aoc-2024/days/day17"
	"aoc-2024/days/day18"
	"aoc-2024/days/day19"
	"aoc-2024/days/day2"
	"aoc-2024/days/day20"
	"aoc-2024/days/day21"
	"aoc-2024/days/day22"
	"aoc-2024/days/day23"
	"aoc-2024/days/day24"
	"aoc-2024/days/day25"
	"aoc-2024/days/day3"
	"aoc-2024/days/day4"
	"aoc-2024/days/day5"
	"aoc-2024/days/day6"
	"aoc-2024/days/day7"
	"aoc-2024/days/day8"
	"aoc-2024/days/day9"
	"flag"
	"fmt"
)

func main() {
	var day int
	flag.IntVar(&day, "Day", 1, "Specify the day to execute")
	flag.Parse()

	// Map of days to corresponding functions
	dayFunctions := map[int]func(){
		1:  day1.Run,
		2:  day2.Run,
		3:  day3.Run,
		4:  day4.Run,
		5:  day5.Run,
		6:  day6.Run,
		7:  day7.Run,
		8:  day8.Run,
		9:  day9.Run,
		10: day10.Run,
		11: day11.Run,
		12: day12.Run,
		13: day13.Run,
		14: day14.Run,
		15: day15.Run,
		16: day16.Run,
		17: day17.Run,
		18: day18.Run,
		19: day19.Run,
		20: day20.Run,
		21: day21.Run,
		22: day22.Run,
		23: day23.Run,
		24: day24.Run,
		25: day25.Run,
	}

	// Execute the function for the specified day
	if fn, exists := dayFunctions[day]; exists {
		fmt.Println("Executing day:", day)
		fmt.Println("----------------------------")
		fmt.Println("")
		fn()
	} else {
		fmt.Printf("day %d is not implemented yet.\n", day)
	}
}
