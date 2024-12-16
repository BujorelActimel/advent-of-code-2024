package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

type WareHouseMap struct {
	grid          [][]rune
	robotPosition Pos
}

var dirMap = map[rune]Pos{
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
	'^': {-1, 0},
}

var expandMap = map[rune]string{
	'#': "##",
	'O': "[]",
	'.': "..",
	'@': "@.",
}

func parseInputPart1(filename string) (WareHouseMap, []rune) {
	var whMap WareHouseMap

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		var row []rune
		for _, ch := range line {
			row = append(row, ch)
		}
		whMap.grid = append(whMap.grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var moves []rune
	for scanner.Scan() {
		for _, ch := range scanner.Text() {
			moves = append(moves, ch)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return whMap, moves
}

func (m *WareHouseMap) initRobot() {
	for i, row := range m.grid {
		for j, ch := range row {
			if ch == '@' {
				m.robotPosition.x = i
				m.robotPosition.y = j
				return
			}
		}
	}
}

func (m *WareHouseMap) moveRobotPart1(direction rune) {
	switch direction {
	case '<': // left
		switch m.grid[m.robotPosition.x][m.robotPosition.y-1] {
		case '.':
			m.grid[m.robotPosition.x][m.robotPosition.y-1] = '@'
			m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
			m.robotPosition.y--
		case 'O':
			for i := m.robotPosition.y - 1; i >= 0; i-- {
				if m.grid[m.robotPosition.x][i] == '.' {
					for j := i; j < m.robotPosition.y; j++ {
						m.grid[m.robotPosition.x][j] = m.grid[m.robotPosition.x][j+1]
					}
					m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
					m.robotPosition.y--
					return
				}
				if m.grid[m.robotPosition.x][i] == '#' {
					return
				}
			}
		case '#':
			return
		}
	case '>': // right
		switch m.grid[m.robotPosition.x][m.robotPosition.y+1] {
		case '.':
			m.grid[m.robotPosition.x][m.robotPosition.y+1] = '@'
			m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
			m.robotPosition.y++
		case 'O':
			for i := m.robotPosition.y + 1; i < len(m.grid[m.robotPosition.x]); i++ {
				if m.grid[m.robotPosition.x][i] == '.' {
					for j := i; j > m.robotPosition.y; j-- {
						m.grid[m.robotPosition.x][j] = m.grid[m.robotPosition.x][j-1]
					}
					m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
					m.robotPosition.y++
					return
				}
				if m.grid[m.robotPosition.x][i] == '#' {
					return
				}
			}
		case '#':
			return
		}

	case '^': // up
		switch m.grid[m.robotPosition.x-1][m.robotPosition.y] {
		case '.':
			m.grid[m.robotPosition.x-1][m.robotPosition.y] = '@'
			m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
			m.robotPosition.x--
		case 'O':
			for i := m.robotPosition.x - 1; i >= 0; i-- {
				if m.grid[i][m.robotPosition.y] == '.' {
					for j := i; j < m.robotPosition.x; j++ {
						m.grid[j][m.robotPosition.y] = m.grid[j+1][m.robotPosition.y]
					}
					m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
					m.robotPosition.x--
					return
				}
				if m.grid[i][m.robotPosition.y] == '#' {
					return
				}
			}
		case '#':
			return
		}

	case 'v': // down
		switch m.grid[m.robotPosition.x+1][m.robotPosition.y] {
		case '.':
			m.grid[m.robotPosition.x+1][m.robotPosition.y] = '@'
			m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
			m.robotPosition.x++
		case 'O':
			for i := m.robotPosition.x + 1; i < len(m.grid); i++ {
				if m.grid[i][m.robotPosition.y] == '.' {
					for j := i; j > m.robotPosition.x; j-- {
						m.grid[j][m.robotPosition.y] = m.grid[j-1][m.robotPosition.y]
					}
					m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
					m.robotPosition.x++
					return
				}
				if m.grid[i][m.robotPosition.y] == '#' {
					return
				}
			}
		case '#':
			return
		}
	}
}

func (m *WareHouseMap) sumPart1() int {
	sum := 0
	for i, row := range m.grid {
		for j, ch := range row {
			if ch == 'O' {
				sum += 100*i + j
			}
		}
	}
	return sum
}

// Part 2 functions
func parseInputPart2(filename string) (WareHouseMap, []rune) {
	var whMap WareHouseMap

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var gridLines []string
	scanner := bufio.NewScanner(file)

	// Read until empty line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		gridLines = append(gridLines, line)
	}

	// Create expanded grid
	for i, line := range gridLines {
		var row []rune
		for j, c := range line {
			expanded := expandMap[c]
			if c == '@' {
				whMap.robotPosition = Pos{i, 2 * j}
			}
			row = append(row, []rune(expanded)...)
		}
		whMap.grid = append(whMap.grid, row)
	}

	// Read moves
	var moves []rune
	for scanner.Scan() {
		line := scanner.Text()
		moves = append(moves, []rune(strings.TrimSpace(line))...)
	}

	return whMap, moves
}

func (m *WareHouseMap) moveRobotPart2(move rune) {
	dir := dirMap[move]
	newX, newY := m.robotPosition.x+dir.x, m.robotPosition.y+dir.y

	// Simple move to empty space
	if m.grid[newX][newY] == '.' {
		m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
		m.grid[newX][newY] = '@'
		m.robotPosition = Pos{newX, newY}
		return
	}

	// Hit a wall
	if m.grid[newX][newY] == '#' {
		return
	}

	// Horizontal movement
	if dir.x == 0 {
		tx, ty := newX, newY
		dist := 0
		// Count boxes to push
		for m.grid[tx][ty] == '[' || m.grid[tx][ty] == ']' {
			dist++
			tx, ty = tx+dir.x, ty+dir.y
		}

		if m.grid[tx][ty] == '#' {
			return
		}

		// Push boxes
		for i := 0; i < dist; i++ {
			m.grid[tx][ty] = m.grid[tx-dir.x][ty-dir.y]
			tx, ty = tx-dir.x, ty-dir.y
		}

		m.grid[newX][newY] = '@'
		m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
		m.robotPosition = Pos{newX, newY}
		return
	}

	// Vertical movement - need to handle multi-box pushes
	toPush := []map[Pos]bool{{Pos{m.robotPosition.x, m.robotPosition.y}: true}}
	noWall := true
	allEmpty := false

	for noWall && !allEmpty {
		nextPush := make(map[Pos]bool)
		allEmpty = true

		for pos := range toPush[len(toPush)-1] {
			if m.grid[pos.x][pos.y] == '.' {
				continue
			}

			tx, ty := pos.x+dir.x, pos.y+dir.y
			if m.grid[tx][ty] != '.' {
				allEmpty = false
			}

			nextPush[Pos{tx, ty}] = true

			if m.grid[tx][ty] == '#' {
				noWall = false
				break
			} else if m.grid[tx][ty] == '[' {
				nextPush[Pos{tx, ty + 1}] = true
			} else if m.grid[tx][ty] == ']' {
				nextPush[Pos{tx, ty - 1}] = true
			}
		}

		if !noWall {
			break
		}
		toPush = append(toPush, nextPush)
	}

	if !noWall {
		return
	}

	// Move all boxes
	for i := len(toPush) - 1; i > 0; i-- {
		for pos := range toPush[i] {
			prevX, prevY := pos.x-dir.x, pos.y-dir.y
			if toPush[i-1][Pos{prevX, prevY}] {
				m.grid[pos.x][pos.y] = m.grid[prevX][prevY]
			} else {
				m.grid[pos.x][pos.y] = '.'
			}
		}
	}

	m.grid[m.robotPosition.x][m.robotPosition.y] = '.'
	m.robotPosition = Pos{newX, newY}
	m.grid[newX][newY] = '@'
}

func (m *WareHouseMap) sumPart2() int {
	total := 0
	for i, row := range m.grid {
		for j, c := range row {
			if c == '[' {
				total += 100*i + j
			}
		}
	}
	return total
}

func (m *WareHouseMap) printMap() {
	for _, row := range m.grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func FirstPuzzle() int {
	whMap, moves := parseInputPart1("day15/input.txt")
	whMap.initRobot()

	for _, move := range moves {
		whMap.moveRobotPart1(move)
	}

	return whMap.sumPart1()
}

func SecondPuzzle() int {
	whMap, moves := parseInputPart2("day15/input.txt")

	for _, move := range moves {
		whMap.moveRobotPart2(move)
	}

	return whMap.sumPart2()
}
