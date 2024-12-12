package day12

import (
	"aoc-2024/utils"
	"bufio"
	"log"
	"os"
)

type Plot struct {
	x, y      int
	plantType rune
	marked    bool
}

type Corner struct {
	x, y float32
}

func parseInput(filename string) ([][]Plot, [][]Plot) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]Plot
	var zones [][]Plot
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		var row []Plot
		for colNum, plantType := range line {
			row = append(row, Plot{lineNum, colNum, plantType, false})
		}
		input = append(input, row)
		lineNum++
	}

	markedPlots := 0

	for markedPlots < len(input)*len(input[0]) {
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[i]); j++ {
				if !input[i][j].marked {
					newZone, marked := getZone(input, input[i][j])
					zones = append(zones, newZone)
					markedPlots += marked
				}
			}
		}
	}

	return input, zones
}

func getZone(input [][]Plot, start Plot) ([]Plot, int) {
	var zone []Plot
	var q utils.Queue[Plot]
	marked := 0

	input[start.x][start.y].marked = true
	marked++
	zone = append(zone, start)
	q.Push(start)

	for !q.IsEmpty() {
		currPlot, err := q.Pop()
		if err != nil {
			log.Fatalln(err)
		}

		//up
		if currPlot.x-1 >= 0 && !input[currPlot.x-1][currPlot.y].marked && input[currPlot.x][currPlot.y].plantType == input[currPlot.x-1][currPlot.y].plantType {
			input[currPlot.x-1][currPlot.y].marked = true
			marked++
			zone = append(zone, input[currPlot.x-1][currPlot.y])
			q.Push(input[currPlot.x-1][currPlot.y])
		}

		//left
		if currPlot.y-1 >= 0 && !input[currPlot.x][currPlot.y-1].marked && input[currPlot.x][currPlot.y].plantType == input[currPlot.x][currPlot.y-1].plantType {
			input[currPlot.x][currPlot.y-1].marked = true
			marked++
			zone = append(zone, input[currPlot.x][currPlot.y-1])
			q.Push(input[currPlot.x][currPlot.y-1])
		}

		//down
		if currPlot.x+1 < len(input) && !input[currPlot.x+1][currPlot.y].marked && input[currPlot.x][currPlot.y].plantType == input[currPlot.x+1][currPlot.y].plantType {
			input[currPlot.x+1][currPlot.y].marked = true
			marked++
			zone = append(zone, input[currPlot.x+1][currPlot.y])
			q.Push(input[currPlot.x+1][currPlot.y])
		}

		//right
		if currPlot.y+1 < len(input[0]) && !input[currPlot.x][currPlot.y+1].marked && input[currPlot.x][currPlot.y].plantType == input[currPlot.x][currPlot.y+1].plantType {
			input[currPlot.x][currPlot.y+1].marked = true
			marked++
			zone = append(zone, input[currPlot.x][currPlot.y+1])
			q.Push(input[currPlot.x][currPlot.y+1])
		}
	}
	return zone, marked
}

func FirstPuzzle() int {
	input, zones := parseInput("day12/input.txt")
	totalPrice := 0

	for _, zone := range zones {
		zonePrice := calculatePrice(input, zone)
		totalPrice += zonePrice
		// fmt.Printf("Zona %d costa %d:\n", zoneNum, zonePrice)
		// for _, plt := range zone {
		// 	fmt.Printf("\t%s:{%d, %d}\n", string(plt.plantType), plt.x, plt.y)
		// }
		// fmt.Println()
	}

	return totalPrice
}

func calculatePrice(input [][]Plot, zone []Plot) int {
	perimeter := 0
	for _, plt := range zone {
		neighbors := 0
		//up
		if plt.x-1 >= 0 && input[plt.x][plt.y].plantType == input[plt.x-1][plt.y].plantType {
			neighbors++
		}

		//left
		if plt.y-1 >= 0 && input[plt.x][plt.y].plantType == input[plt.x][plt.y-1].plantType {
			neighbors++
		}

		//down
		if plt.x+1 < len(input) && input[plt.x][plt.y].plantType == input[plt.x+1][plt.y].plantType {
			neighbors++
		}

		//right
		if plt.y+1 < len(input[0]) && input[plt.x][plt.y].plantType == input[plt.x][plt.y+1].plantType {
			neighbors++
		}
		perimeter += 4 - neighbors
	}
	return len(zone) * perimeter
}

func SecondPuzzle() int {
	_, zones := parseInput("day12/input.txt")
	totalPrice := 0

	for _, zone := range zones {
		totalPrice += corners(zone) * len(zone)
	}

	return totalPrice
}

func corners(zone []Plot) int {
	s := utils.NewSet[Corner]()
	corners := 0

	for _, plt := range zone {
		// coltul stanga sus
		s.Add(Corner{float32(plt.x) - 0.5, float32(plt.y) - 0.5})

		// coltul dreapta sus
		s.Add(Corner{float32(plt.x) - 0.5, float32(plt.y) + 0.5})

		// coltul dreapta jos
		s.Add(Corner{float32(plt.x) + 0.5, float32(plt.y) + 0.5})

		// coltul stanga jos
		s.Add(Corner{float32(plt.x) + 0.5, float32(plt.y) - 0.5})
	}

	for _, corner := range s.Values() {
		upLeft := Plot{int(corner.x - 0.5), int(corner.y - 0.5), ' ', false}
		upRight := Plot{int(corner.x - 0.5), int(corner.y + 0.5), ' ', false}
		downLeft := Plot{int(corner.x + 0.5), int(corner.y - 0.5), ' ', false}
		downRight := Plot{int(corner.x + 0.5), int(corner.y + 0.5), ' ', false}

		neighbors := 0

		for _, plt := range zone {
			if plt.x == upLeft.x && plt.y == upLeft.y {
				upLeft.marked = true
				neighbors++
			}
			if plt.x == upRight.x && plt.y == upRight.y {
				upRight.marked = true
				neighbors++
			}
			if plt.x == downLeft.x && plt.y == downLeft.y {
				downLeft.marked = true
				neighbors++
			}
			if plt.x == downRight.x && plt.y == downRight.y {
				downRight.marked = true
				neighbors++
			}
		}

		if neighbors == 0 || neighbors == 4 {
			corners += 0
		} else if neighbors == 3 || neighbors == 1 {
			corners += 1
		} else {
			if downLeft.marked && upRight.marked {
				corners += 2
			}
			if downRight.marked && upLeft.marked {
				corners += 2
			}
		}
	}

	return corners
}
