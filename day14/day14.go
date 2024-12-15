package day14

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, y int
}

type Robot struct {
	start    Coordinates
	velocity Coordinates
	current  Coordinates
}

func parseInput(filename string) []Robot {
	var robots []Robot

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		startLine, velocityLine := parts[0], parts[1]
		startPos := parseCoord(startLine)
		velocity := parseCoord(velocityLine)
		robots = append(robots, Robot{
			startPos,
			velocity,
			startPos,
		})
	}

	return robots
}

func parseCoord(line string) Coordinates {
	parts := strings.Split(line, "=")
	coords := parts[1]
	nums := strings.Split(coords, ",")

	x, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatalln(err)
	}

	y, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatalln(err)
	}

	return Coordinates{x, y}
}

func FirstPuzzle() int {
	robots := parseInput("day14/input.txt")

	const (
		length = 103
		width  = 101
	)

	// drawGrid(robots, length, width)

	for i := range robots {
		robots[i].current = moveRobot(robots[i], length, width, 100) // after 100 secinds
	}

	// drawGrid(robots, length, width)

	return safetyFactor(robots, length, width)
}

func drawGrid(robots []Robot, length int, width int) {
	fmt.Println(length, width)
	grid := make([][]int, length)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			grid[i][j] = 0
		}
	}

	for _, robot := range robots {
		grid[robot.current.y][robot.current.x]++
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 0 {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Println()
	}
}

func getGrid(robots []Robot, length int, width int) [][]int {
	grid := make([][]int, length)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			grid[i][j] = 0
		}
	}

	for _, robot := range robots {
		grid[robot.current.y][robot.current.x]++
	}

	return grid
}

func moveRobot(robot Robot, length int, width int, seconds int) Coordinates {
	newX := (robot.start.x + seconds*robot.velocity.x) % width
	newY := (robot.start.y + seconds*robot.velocity.y) % length

	if newX < 0 {
		newX = width + newX
	}

	if newY < 0 {
		newY = length + newY
	}

	return Coordinates{newX, newY}
}

func safetyFactor(robots []Robot, length int, width int) int {
	first, second, third, fourth := 0, 0, 0, 0
	grid := getGrid(robots, length, width)

	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			// first
			if i < length/2 && j < width/2 {
				first += grid[i][j]
			}
			// second
			if i > length/2 && j < width/2 {
				second += grid[i][j]
			}
			// third
			if i > length/2 && j > width/2 {
				third += grid[i][j]
			}
			// fourth
			if i < length/2 && j > width/2 {
				fourth += grid[i][j]
			}
		}
	}

	return first * second * third * fourth
}

func SecondPuzzle() int {

	length := 103
	width := 101

	for i := 0; i <= 10000; i++ {
		robots := parseInput("day14/input.txt")

		for j := range robots {
			robots[j].current = moveRobot(robots[j], length, width, i)
		}

		grid := getGrid(robots, length, width)
		if hasLine(grid) {
			err := createImage(grid, "day14/easterEgg.png")
			if err != nil {
				panic(err)
			}
			return i
		}
	}
	return 0
}

func createImage(grid [][]int, filename string) error {
	length := len(grid)
	width := len(grid[0])

	img := image.NewRGBA(image.Rect(0, 0, width, length))

	for y := 0; y < length; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 0 {
				img.Set(x, y, color.RGBA{0, 0, 0, 255}) // black if 0
			} else {
				img.Set(x, y, color.RGBA{0, 255, 0, 255}) // green otherwise
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func hasLine(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		streak := 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != 0 {
				streak++
			} else {
				if streak > 30 { // arbitrary number
					return true
				}
				streak = 0
			}
		}
	}
	return false
}
