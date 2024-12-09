package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Block struct {
	symbol byte
	id     int
}

type Interval struct {
	start int
	end   int
}

func parseInput(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bytes := scanner.Bytes()

	return bytes
}

func FirstPuzzle() int {
	bytes := parseInput("day9/input.txt")

	var blocks []Block

	for i, c := range bytes {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatalln(err)
		}

		newBlock := Block{}

		if i%2 == 0 {
			newBlock.id = i / 2
			newBlock.symbol = '#'
		} else {
			newBlock.id = -1
			newBlock.symbol = '.'
		}

		for j := 0; j < val; j++ {
			blocks = append(blocks, newBlock)
		}
	}

	length := len(blocks)
	currBl := firstSpace(blocks, -1)
	lastBl := lastBlock(blocks, length)

	for currBl < lastBl && currBl < length {
		// swap
		blocks[currBl], blocks[lastBl] = blocks[lastBl], blocks[currBl]

		// move to the last non-space byte
		currBl = firstSpace(blocks, currBl)

		// move to the first available space byte
		lastBl = lastBlock(blocks, lastBl)
	}

	return checkSum(blocks)
}

func firstSpace(blocks []Block, prevSpace int) int {
	for i := prevSpace + 1; i < len(blocks); i++ {
		if blocks[i].id == -1 {
			return i
		}
	}
	return -1
}

func lastBlock(blocks []Block, prevBlock int) int {
	for i := prevBlock - 1; i >= 0; i-- {
		if blocks[i].id != -1 {
			return i
		}
	}
	return -1
}

func SecondPuzzle() int {
	bytes := parseInput("day9/input.txt")

	var blocks []Block

	for i, c := range bytes {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatalln(err)
		}

		newBlock := Block{}

		if i%2 == 0 {
			newBlock.id = i / 2
			newBlock.symbol = '#'
		} else {
			newBlock.id = -1
			newBlock.symbol = '.'
		}

		for j := 0; j < val; j++ {
			blocks = append(blocks, newBlock)
		}
	}

	file := getFile(blocks, len(blocks))

	for file.start > 0 {
		fileSize := file.end - file.start + 1
		spaces := getSpaces(blocks, fileSize)

		if len(spaces) > 0 && spaces[0].end < file.start {
			moveFile(blocks, file, spaces[0])
		}

		file = getFile(blocks, file.start) // get the next file
	}

	fileSize := file.end - file.start + 1
	spaces := getSpaces(blocks, fileSize)

	if len(spaces) > 0 && spaces[0].end < file.start {
		moveFile(blocks, file, spaces[0])
	}

	return checkSum(blocks)
}

func getSpaces(blocks []Block, size int) []Interval {
	var intervals []Interval

	var start, end int

	for i := 0; i < len(blocks); i++ {
		if blocks[i].id == -1 {
			start = i
			end = start + 1
			for end < len(blocks) && blocks[end].id == -1 {
				end++
			}
			if end < len(blocks) {
				if end-start >= size {
					intervals = append(intervals, Interval{start, end - 1})
				}
				i = end
			} else {
				break
			}
		}
	}

	return intervals
}

func getFile(blocks []Block, lastFileStart int) Interval {
	var file Interval

	for i := lastFileStart - 1; i >= 0; i-- {
		if blocks[i].id != -1 {
			file.end = i
			file.start = file.end - 1
			for file.start >= 0 && blocks[file.start].id == blocks[file.end].id {
				file.start--
			}
			file.start++
			break
		}
	}

	return file
}

func moveFile(blocks []Block, file Interval, space Interval) {
	fileSize := file.end - file.start
	for i := space.start; i <= space.start+fileSize; i++ {
		blocks[i], blocks[file.start+i-space.start] = blocks[file.start+i-space.start], blocks[i]
	}
}

func checkSum(blocks []Block) int {
	sum := 0

	for i := 0; i < len(blocks); i++ {
		if blocks[i].id != -1 {
			sum += i * blocks[i].id
		}
	}

	return sum
}

func status(blocks []Block) {
	fmt.Print("\n\nBytes: ")
	for i := 0; i < len(blocks); i++ {
		fmt.Print(string(blocks[i].symbol))
	}
	fmt.Print("\nID-s:  ")
	for i := 0; i < len(blocks); i++ {
		if blocks[i].id != -1 {
			fmt.Print(blocks[i].id)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
