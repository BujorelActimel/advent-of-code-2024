package day10

import (
	"aoc-2024/utils"
	"bufio"
	"log"
	"os"
	"strconv"
)

type Pos struct {
	x, y int
}

func parseInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	var mountainMap [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int

		for _, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalln(err)
			}
			row = append(row, val)
		}

		mountainMap = append(mountainMap, row)
	}

	return mountainMap
}

func FirstPuzzle() int {
	mountainMap := parseInput("day10/input.txt")

	result := 0

	for i, row := range mountainMap {
		for j, height := range row {
			if height == 0 {
				scr := score(mountainMap, Pos{i, j})
				// fmt.Printf("The trail head at [%d, %d] has a score of %d\n", i, j, scr)
				result += scr
			}
		}
	}

	return result
}

func score(mountainMap [][]int, start Pos) int {
	var q utils.Queue[Pos]
	destinations := utils.NewSet[Pos]()

	var (
		UP    Pos = Pos{-1, 0}
		DOWN  Pos = Pos{1, 0}
		LEFT  Pos = Pos{0, -1}
		RIGHT Pos = Pos{0, 1}
	)

	directions := []Pos{UP, DOWN, LEFT, RIGHT}

	q.Push(start)

	for !q.IsEmpty() {
		currPos, err := q.Pop()
		if err != nil {
			log.Fatalln(err)
		}

		for _, direction := range directions {
			newPos := Pos{currPos.x + direction.x, currPos.y + direction.y}
			if validPosition(mountainMap, newPos) && mountainMap[newPos.x][newPos.y] == mountainMap[currPos.x][currPos.y]+1 {
				if mountainMap[newPos.x][newPos.y] == 9 {
					destinations.Add(newPos)
				} else {
					q.Push(newPos)
				}
			}
		}
	}

	return destinations.Size()
}

func SecondPuzzle() int {
	mountainMap := parseInput("day10/input.txt")

	result := 0

	for i, row := range mountainMap {
		for j, height := range row {
			if height == 0 {
				rtng := rating(mountainMap, Pos{i, j})
				// fmt.Printf("The trail head at [%d, %d] has a rating of %d\n", i, j, rtng)
				result += rtng
			}
		}
	}

	return result
}

func rating(mountainMap [][]int, start Pos) int {
	var q utils.Queue[Pos]
	var destinations []Pos

	var (
		UP    Pos = Pos{-1, 0}
		DOWN  Pos = Pos{1, 0}
		LEFT  Pos = Pos{0, -1}
		RIGHT Pos = Pos{0, 1}
	)

	directions := []Pos{UP, DOWN, LEFT, RIGHT}

	q.Push(start)

	for !q.IsEmpty() {
		currPos, err := q.Pop()
		if err != nil {
			log.Fatalln(err)
		}

		for _, direction := range directions {
			newPos := Pos{currPos.x + direction.x, currPos.y + direction.y}
			if validPosition(mountainMap, newPos) && mountainMap[newPos.x][newPos.y] == mountainMap[currPos.x][currPos.y]+1 {
				if mountainMap[newPos.x][newPos.y] == 9 {
					destinations = append(destinations, newPos)
				} else {
					q.Push(newPos)
				}
			}
		}
	}

	return len(destinations)
}

func validPosition(mountainMap [][]int, pos Pos) bool {
	return pos.x >= 0 && pos.x < len(mountainMap) && pos.y >= 0 && pos.y < len(mountainMap[pos.x])
}
