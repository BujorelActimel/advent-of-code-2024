package main

import (
	"aoc-2024/day1"
	"aoc-2024/day2"
	"aoc-2024/day3"
	"aoc-2024/day4"
	"aoc-2024/day5"
	"aoc-2024/day6"
	"aoc-2024/day7"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [day]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1":
		fmt.Println(day1.FirstPuzzle())
		fmt.Println(day1.SecondPuzzle())
	case "2":
		fmt.Println(day2.FirstPuzzle())
		fmt.Println(day2.SecondPuzzle())
	case "3":
		fmt.Println(day3.FirstPuzzle())
		fmt.Println(day3.SecondPuzzle())
	case "4":
		fmt.Println(day4.FirstPuzzle())
		fmt.Println(day4.SecondPuzzle())
	case "5":
		fmt.Println(day5.FirstPuzzle())
		fmt.Println(day5.SecondPuzzle())
	case "6":
		fmt.Println(day6.FirstPuzzle())
		fmt.Println(day6.SecondPuzzle())
	case "7":
		fmt.Println(day7.FirstPuzzle())
		fmt.Println(day7.SecondPuzzle())
	default:
		fmt.Println("Unknown argument:", os.Args[1])
	}
}
