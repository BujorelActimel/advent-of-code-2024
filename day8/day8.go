package day8

import (
	"bufio"
	"log"
	"os"
)

type Position struct {
	x int
	y int
}

func parseInput(filename string) (map[string][]Position, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	antenas := make(map[string][]Position)
	lineNum := 0
	colNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		colNum = len(line)

		for colNum, c := range line {
			if c != '.' {
				antenas[string(c)] = append(antenas[string(c)], Position{lineNum, colNum})
			}
		}
		lineNum++
	}

	return antenas, lineNum, colNum
}

func FirstPuzzle() int {
	antenas, numOfLines, numOfCols := parseInput("day8/input.txt")

	var antinodes []Position

	for key := range antenas {
		currentAntenas := antenas[key]
		length := len(currentAntenas)

		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if i == j {
					continue
				}

				firstAntena := currentAntenas[i]
				secondAntena := currentAntenas[j]

				difX := firstAntena.x - secondAntena.x
				difY := firstAntena.y - secondAntena.y

				antinode := Position{secondAntena.x - difX, secondAntena.y - difY}

				if !contains(antinodes, antinode) && validPosition(antinode, numOfLines, numOfCols) {
					antinodes = append(antinodes, antinode)
				}
			}
		}
	}

	return len(antinodes)
}

func validPosition(pos Position, lines int, cols int) bool {
	return pos.x >= 0 && pos.x < lines && pos.y >= 0 && pos.y < cols
}

func contains(positions []Position, position Position) bool {
	for _, pos := range positions {
		if equals(pos, position) {
			return true
		}
	}
	return false
}

func equals(pos1 Position, pos2 Position) bool {
	return pos1.x == pos2.x && pos1.y == pos2.y
}

func SecondPuzzle() int {
	antenas, numOfLines, numOfCols := parseInput("day8/input.txt")

	var antinodes []Position

	for key := range antenas {
		currentAntenas := antenas[key]
		length := len(currentAntenas)

		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if i == j {
					continue
				}

				firstAntena := currentAntenas[i]
				secondAntena := currentAntenas[j]

				difX := firstAntena.x - secondAntena.x
				difY := firstAntena.y - secondAntena.y

				antinode := Position{secondAntena.x - difX, secondAntena.y - difY}

				for validPosition(antinode, numOfLines, numOfCols) {
					if !contains(antinodes, antinode) && validPosition(antinode, numOfLines, numOfCols) {
						antinodes = append(antinodes, antinode)
					}

					antinode = Position{antinode.x - difX, antinode.y - difY}
				}

				if !contains(antinodes, firstAntena) {
					antinodes = append(antinodes, firstAntena)
				}
				if !contains(antinodes, secondAntena) {
					antinodes = append(antinodes, secondAntena)
				}
			}
		}
	}

	return len(antinodes)
}
