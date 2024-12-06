package day6

import (
	"bufio"
	"log"
	"os"
)

type Pos struct {
	x int
	y int
}

// type Pair struct {
// 	marked    bool
// 	direction Pos
// }

func FirstPuzzle() int {
	file, err := os.Open("day6/test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	startingPos, direction := FindStarting(input)

	return move(startingPos, direction, input)
}

func FindStarting(input []string) (Pos, Pos) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '^' {
				return Pos{i, j}, Pos{-1, 0}
			}
			if input[i][j] == 'V' {
				return Pos{i, j}, Pos{1, 0}
			}
			if input[i][j] == '<' {
				return Pos{i, j}, Pos{0, -1}
			}
			if input[i][j] == '>' {
				return Pos{i, j}, Pos{0, 1}
			}
		}
	}
	return Pos{-1, -1}, Pos{-1, -1}
}

func move(pos Pos, direction Pos, input []string) int {
	steps := 0

	markedSteps := make([][]int, len(input))
	for i := range markedSteps {
		markedSteps[i] = make([]int, len(input[i]))
	}

	for insideMap(pos, input) {
		// fmt.Println("Sunt la pozitia", pos, "si merg in directia", direction)
		// fmt.Println("Am mers in", steps, "locuri unice")

		dirX := direction.x
		dirY := direction.y

		if !insideMap(Pos{pos.x + dirX, pos.y + dirY}, input) {
			return steps + 1
		}

		// daca in fata am obstacol, turn
		if input[pos.x+dirX][pos.y+dirY] == '#' {
			if markedSteps[pos.x][pos.y] == 0 {
				markedSteps[pos.x][pos.y] = 1
				steps++
			}
			direction = turnRight(direction)
			continue
		}

		// daca pot sa merg in fata, ma duc in fata
		if input[pos.x+dirX][pos.y+dirY] != '#' {
			if markedSteps[pos.x][pos.y] == 0 { // daca Nu e marcat
				markedSteps[pos.x][pos.y] = 1
				steps++
			}
			pos.x += dirX
			pos.y += dirY
		}
	}

	return steps
}

func insideMap(pos Pos, input []string) bool {
	return pos.x >= 0 && pos.x < len(input) && pos.y >= 0 && pos.y < len(input[pos.x])
}

func turnRight(direction Pos) Pos {
	if direction.x == -1 { // up
		return Pos{0, 1} // right
	}
	if direction.x == 1 { // down
		return Pos{0, -1} // left
	}
	if direction.y == 1 { // right
		return Pos{1, 0} // down
	}
	if direction.y == -1 { // left
		return Pos{-1, 0} // up
	}

	return Pos{}
}

func SecondPuzzle() int {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	startingPos, direction := FindStarting(input)

	updatedInput := make([]string, len(input))

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '.' {
				copy(updatedInput, input)

				runes := []rune(updatedInput[i])
				runes[j] = '#'
				updatedInput[i] = string(runes)

				if checkLoop(startingPos, direction, updatedInput) {
					result++
				}
			}
		}
	}
	return result
}

func checkLoop(pos Pos, direction Pos, input []string) bool {
	// if i go out of the matrix - false
	// if i find a marked positon and i have the
	// same direction as i had then, it's a loop

	// markedSteps := make([][]Pair, len(input))
	// for i := range markedSteps {
	// 	markedSteps[i] = make([]Pair, len(input[i]))
	// 	for j := range markedSteps[i] {
	// 		markedSteps[i][j] = Pair{false, Pos{-1, -1}}
	// 	}
	// }

	steps := 0

	for insideMap(pos, input) && steps < 10000 {
		dirX := direction.x
		dirY := direction.y

		if !insideMap(Pos{pos.x + dirX, pos.y + dirY}, input) {
			return false
		}

		// daca am mai fost pe aici si aveam aceeasi directie -> true
		// if markedSteps[pos.x][pos.y].marked && markedSteps[pos.x][pos.y].direction == direction {
		// 	return true
		// }

		// daca am obstacol in fata, turn si continue
		if input[pos.x+dirX][pos.y+dirY] == '#' {
			// markedSteps[pos.x][pos.y] = Pair{true, direction}
			direction = turnRight(direction)
			continue
		} else {
			// markedSteps[pos.x][pos.y] = Pair{true, direction}
			steps++
			pos.x += dirX
			pos.y += dirY
		}

	}

	return true
}
