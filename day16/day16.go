package day16

import (
	"bufio"
	"math"
	"os"
)

type Index struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Node struct {
	idx Index
	dir Direction
}

type State struct {
	reindeer Node
	path     []Index
	score    int
}

var directions = []Direction{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

func parseInput(filename string) ([][]byte, Index, Index) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]byte
	var start, end Index

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		copy(row, line)

		for i, ch := range line {
			if ch == 'S' {
				start = Index{len(grid), i}
				row[i] = '.'
			} else if ch == 'E' {
				end = Index{len(grid), i}
				row[i] = '.'
			}
		}
		grid = append(grid, row)
	}

	return grid, start, end
}

func getNeighbors(reindeer Node) []Node {
	neighbors := make([]Node, 0, 3)
	oppositeDir := Direction{-reindeer.dir.dx, -reindeer.dir.dy}

	for _, dir := range directions {
		if dir == oppositeDir {
			continue
		}
		newIdx := Index{
			x: reindeer.idx.x + dir.dx,
			y: reindeer.idx.y + dir.dy,
		}
		neighbors = append(neighbors, Node{newIdx, dir})
	}
	return neighbors
}

func solve(grid [][]byte, start, end Index) (int, int) {
	minScore := math.MaxInt
	sizeToIndices := make(map[int][]Index)
	visited := make(map[Node]int)

	// Start facing right
	startNode := Node{start, Direction{0, 1}}
	queue := []State{{
		reindeer: startNode,
		path:     []Index{start},
		score:    0,
	}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.score > minScore {
			continue
		}

		if curr.reindeer.idx == end {
			if curr.score <= minScore {
				minScore = curr.score
				sizeToIndices[minScore] = append(sizeToIndices[minScore], curr.path...)
			}
			continue
		}

		for _, next := range getNeighbors(curr.reindeer) {
			if next.idx.x < 0 || next.idx.x >= len(grid) ||
				next.idx.y < 0 || next.idx.y >= len(grid[0]) ||
				grid[next.idx.x][next.idx.y] == '#' {
				continue
			}

			score := curr.score + 1
			if curr.reindeer.dir != next.dir {
				score += 1000
			}

			if prev, exists := visited[next]; exists && prev < score {
				continue
			}
			visited[next] = score

			newPath := make([]Index, len(curr.path))
			copy(newPath, curr.path)

			queue = append(queue, State{
				reindeer: next,
				path:     append(newPath, next.idx),
				score:    score,
			})
		}
	}

	uniqueTiles := make(map[Index]bool)
	for _, idx := range sizeToIndices[minScore] {
		uniqueTiles[idx] = true
	}

	return minScore, len(uniqueTiles)
}

func FirstPuzzle() int {
	grid, start, end := parseInput("day16/input.txt")
	score, _ := solve(grid, start, end)
	return score
}

func SecondPuzzle() int {
	grid, start, end := parseInput("day16/input.txt")
	_, tiles := solve(grid, start, end)
	return tiles
}
