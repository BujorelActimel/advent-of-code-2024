package day4

import (
	"bufio"
	"fmt"
	"os"
)

func FirstPuzzle() int {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return countXmas(input)
}

func countXmas(input []string) int {
	count := 0

	for i := range len(input) {
		for j := range len(input[i]) {
			count += (horizontal(input, i, j) + vertical(input, i, j) + diagonal(input, i, j))
		}
	}

	return count
}

func horizontal(input []string, x int, y int) int {
	count := 0
	word := ""

	if y+3 < len(input[x]) {
		word = input[x][y : y+4]
	}

	if word == "XMAS" {
		count++
	}

	if reverseString(word) == "XMAS" {
		count++
	}

	return count
}

func vertical(input []string, x int, y int) int {
	count := 0
	word := ""

	if x+3 < len(input) {
		for i := 0; i < 4; i++ {
			word += string(input[x+i][y])
		}
	}

	if word == "XMAS" {
		count++
	}

	if reverseString(word) == "XMAS" {
		count++
	}

	return count
}

func diagonal(input []string, x int, y int) int {
	count := 0

	// rigtht diagonal
	word1 := ""
	if x+3 < len(input) && y+3 < len(input[x]) {
		for i := 0; i < 4; i++ {
			word1 += string(input[x+i][y+i])
		}
	}

	if word1 == "XMAS" {
		count++
	}

	if reverseString(word1) == "XMAS" {
		count++
	}

	// left diagonal
	word2 := ""
	if x+3 < len(input) && y-3 >= 0 {
		for i := 0; i < 4; i++ {
			word2 += string(input[x+i][y-i])
		}
	}

	if word2 == "XMAS" {
		count++
	}

	if reverseString(word2) == "XMAS" {
		count++
	}

	return count
}

func SecondPuzzle() int {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return countX_MAS(input)
}

func countX_MAS(input []string) int {
	count := 0

	for i := range len(input) {
		for j := range len(input[i]) {
			if x_mas(input, i, j) {
				count++
			}
		}
	}

	return count
}

func x_mas(input []string, x int, y int) bool {
	if !(x+2 < len(input) && y+2 < len(input[x])) {
		return false
	}
	left := string(input[x][y]) + string(input[x+1][y+1]) + string(input[x+2][y+2])
	right := string(input[x][y+2]) + string(input[x+1][y+1]) + string(input[x+2][y])

	return (left == "MAS" || reverseString(left) == "MAS") && (right == "MAS" || reverseString(right) == "MAS")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
