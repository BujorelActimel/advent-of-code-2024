package main

import (
	"aoc-2024/day1"
	"aoc-2024/day10"
	"aoc-2024/day11"
	"aoc-2024/day12"
	"aoc-2024/day13"
	"aoc-2024/day14"
	"aoc-2024/day15"
	"aoc-2024/day16"
	"aoc-2024/day2"
	"aoc-2024/day3"
	"aoc-2024/day4"
	"aoc-2024/day5"
	"aoc-2024/day6"
	"aoc-2024/day7"
	"aoc-2024/day8"
	"aoc-2024/day9"
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
	case "8":
		fmt.Println(day8.FirstPuzzle())
		fmt.Println(day8.SecondPuzzle())
	case "9":
		fmt.Println(day9.FirstPuzzle())
		fmt.Println(day9.SecondPuzzle())
	case "10":
		fmt.Println(day10.FirstPuzzle())
		fmt.Println(day10.SecondPuzzle())
	case "11":
		fmt.Println(day11.FirstPuzzle())
		fmt.Println(day11.SecondPuzzle())
	case "12":
		fmt.Println(day12.FirstPuzzle())
		fmt.Println(day12.SecondPuzzle())
	case "13":
		fmt.Println(day13.FirstPuzzle())
		fmt.Println(day13.SecondPuzzle())
	case "14":
		fmt.Println(day14.FirstPuzzle())
		fmt.Println(day14.SecondPuzzle())
	case "15":
		fmt.Println(day15.FirstPuzzle())
		fmt.Println(day15.SecondPuzzle())
	case "16":
		fmt.Println(day16.FirstPuzzle())
		fmt.Println(day16.SecondPuzzle())
	default:
		fmt.Println("Unknown argument:", os.Args[1])
	}
}
